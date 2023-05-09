package interfaces

import "fmt"

type RedirectService interface {
	ParseMethod(command string) RedirectService
	ParseParameters(command string) RedirectService
	Redirect(command string) error
}

type RedirectServiceImpl struct {
	ParamsString string
	Method func()
	Params map[string]interface{}

}

func (r *RedirectServiceImpl) ParseMethod(command string) RedirectService {
	fmt.Println("Original parsemedthod Method Called")
	r.ParamsString = "Original parsemedthod Method Stored"

	return r

}

func (r *RedirectServiceImpl) ParseParameters(command string) RedirectService {

	return r
}



func (r *RedirectServiceImpl) Redirect(command string) error {
	panic("implement me")
}



