package http

import (
	"encoding/csv"
	"fmt"
	"github.com/go-resty/resty/v2"
	"io"
	"os"
	"sync"
	"time"
)

var client  *resty.Client
var cnt Count
var waitGroup sync.WaitGroup

type Count struct {
	count int64
	lock sync.Mutex
}



func init() {
	client = resty.New()
	cnt = Count{count: 0}
}



func doDataCloudReq(ip string) {

	defer func() {
		waitGroup.Done()
	}()

	req := client.R()
	req.SetQueryParam("ip", ip)
	req.SetQueryParam("localityLanguage", "en")
	req.SetQueryParam("key", "96e9d7222d6d42acbf2266ae15e09a9a")
	rsp, err := req.Execute(resty.MethodGet, "https://api.bigdatacloud.net/data/ip-geolocation-full")
	req.SetHeader("Accept", "application/json")
	req.SetHeader("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36")
	if err != nil {
		fmt.Printf("ReqError: ", err)
	}
	cnt.lock.Lock()
	cnt.count++
	result := fmt.Sprintf("RspCode: %d, Count: %d, Time : %d, IP: %s", rsp.StatusCode(), cnt.count, time.Now().Second(), ip)
	fmt.Println(result)
	cnt.lock.Unlock()


}


func TestDataCloudQps(){

	f, err := os.Open("http/test.csv")

	if err != nil {
		fmt.Println(err)
		return
	}

	r := csv.NewReader(f)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		for value := range record {
			waitGroup.Add(1)
			go doDataCloudReq(record[value])
		}
	}

	waitGroup.Wait()



}