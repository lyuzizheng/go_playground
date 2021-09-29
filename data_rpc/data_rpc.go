
package main

import (
	ServiceClient "code.byted.org/dp/data_service_rpc/clients/dp/data_service/api_rpc"

	"time"

	"code.byted.org/gopkg/context"
	"code.byted.org/gopkg/logs"
	"code.byted.org/inf/infsecc"
	"code.byted.org/kite/kitc"
	"code.byted.org/kite/kitutil"

	"code.byted.org/dp/data_service_rpc/thrift_gen/base"
	"code.byted.org/dp/data_service_rpc/thrift_gen/dp/data_service/api_rpc"
)

var (
	client *ServiceClient.Client
)

func InitClient() {
	client = ServiceClient.MustNewClient(
		"dp.data_service.api_rpc", // 服务PSM
		kitc.WithTimeout(100 * time.Second),
		kitc.WithConnTimeout(5 * time.Second),
		// kitc.WithHostPort("localhost:8888"),
		// kitc.WithIDC("maliva"),
	)
}

func RunClient() {
	var err error

	// put your api data
	user := `lyuzizheng`  // 调用者
	apiPath := `/im_cloud/ods_metrics_test/ods_metrics_test` // api path
	apiData := `{"metric":"im.proxy.gateway.calledby.success.status.throughput","date":"20210928","hour":"03"}` // 请求体

	// use dps-sdk to fetch dps token (for TCE/FaaS)
	// for devbox or mac, please use doas (https://bytedance.feishu.cn/docs/85d9fb0GEEHalMJqmNvJNg)
	dpsToken, err := infsecc.GetToken(false)


	// general rpc client code for all api
	base := base.NewBase()
	base.Addr, base.Caller, base.Client, base.LogID = "127.0.0.1", "example_client", "binary", "00001"

	ctx := context.Background()
	ctx = kitutil.NewCtxWithServiceName(ctx, "RunClient")

	req := api_rpc.NewGeneralApiRequest()
	req.Base = base
	req.DpsToken = &dpsToken

	req.User = &user
	req.Path = apiPath
	req.Data = apiData

	Method1Resp, err := client.CallApi(ctx, req)
	if err != nil || Method1Resp == nil {
		logs.Errorf("Call CallApi error: %s", err)
	} else {
		if Method1Resp.BaseResp.StatusCode == 0 {
			logs.Infof("Call CallApi success: %+v", Method1Resp)
		} else {
			logs.Error("Call CallApi failed: %+v", Method1Resp.BaseResp.StatusMessage)
		}
	}

}

func main() {
	kitc.SetCallLog(logs.DefaultLogger())
	logs.SetLevel(logs.LevelTrace)
	InitClient()
	RunClient()
	logs.Flush()
}
