package http

import (
	"code.byted.org/middleware/hertz/byted"
	"code.byted.org/middleware/hertz/pkg/protocol"
	"code.byted.org/middleware/hertz/pkg/protocol/consts"
	"context"
	"fmt"
)

func TestBot(){

	c, err := byted.NewClient()
	if err != nil {
		return
	}
	req := &protocol.Request{}
	res := &protocol.Response{}
	req.SetMethod(consts.MethodGet)                            //设置请求方法
	req.Header.SetContentTypeBytes([]byte("application/json")) //设置请求header
	req.SetRequestURI("http://example.com")                    //设置请求url
	err = c.Do(context.Background(), req, res)
	if err != nil {
		return
	}
	fmt.Printf("%v", string(res.Body())) //读取响应body


}