package interface_test

import "fmt"

type PushService interface {
	HandlePush(push string)
}

type PushServiceImpl struct {
	ParamsString string
	Method int
	Params map[string]interface{}

}

func (r *PushServiceImpl) HandlePush(push string)  {
	fmt.Println("Original parsemedthod Method Called")
	r.ParamsString = "Original parsemedthod Method Stored"
	r.Method = 0

}



