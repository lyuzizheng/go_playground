package jsonPlay

import (
	"encoding/json"
	"fmt"
	"reflect"

)

type Card interface {
	ToCard(defaultOpenChatId string) interface{}
}

type CardTemplate1 struct {
	OpenId     string `jsonPlay:"open_id,omitempty"`
	UserId     string `jsonPlay:"user_id,omitempty"`
	OpenChatId string `jsonPlay:"open_chat_id"`

	//Compulsory
	Title      string               `jsonPlay:"title"`
	TitleColor string `jsonPlay:"title_color"`
	//Compulsory
	Message string `jsonPlay:"message"`
	//Compulsory
	FootNote string              `jsonPlay:"foot_note"`
}

func (c CardTemplate1) ToCard(defaultOpenChatID string) interface{} {

	return nil

}

type CardTemplateDefault struct {
	OpenId     string `jsonPlay:"open_id,omitempty"`
	UserId     string `jsonPlay:"user_id,omitempty"`
	OpenChatId string `jsonPlay:"open_chat_id"`

	//Compulsory
	Message string `jsonPlay:"message"`
	//Compulsory

}

func (c CardTemplateDefault) ToCard(defaultOpenChatID string) interface{}{
	return nil


}


func UnmarshallFromInterface(templateID string){
	rawCard := "{ \"message\" : \"Hello from demo tcc config template\"}"


	var card interface{}

	var err error
	switch templateID {
	case "1":
		cardNew := CardTemplate1{}
		err = json.Unmarshal([]byte(rawCard), &cardNew)
		card = cardNew
	default:
		cardNew := CardTemplateDefault{}
		err = json.Unmarshal([]byte(rawCard), &cardNew)
		card = cardNew
	}
	//err = jsonPlay.Unmarshal([]byte(rawCard), &card)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Printf("Struct: %+v", card)
	}

}

func UnmarshallToInterface()  {

	//test1 := "{ \"message\" : \"Hello from demo tcc config template\"}"
	//test2 := "[{ \"message\" : \"Hello from demo tcc config template\"},{ \"message\" : \"Hello from demo tcc config template\"}]"
	test3 := "{ \"integer\" : 234, \"booltest\" : true}"


	//var result interface{}
	//err := jsonPlay.Unmarshal([]byte(test1), &result)
	//if err != nil {
	//	fmt.Printf("Error: %s", err.Error())
	//} else {
	//	fmt.Printf("Struct: %+v", result)
	//	fmt.Printf("type is %s", reflect.TypeOf(result))
	//}
	//var result2 interface{}
	//err2 := jsonPlay.Unmarshal([]byte(test2), &result2)
	//if err2 != nil {
	//	fmt.Printf("Error: %s", err2.Error())
	//} else {
	//	fmt.Printf("Struct: %+v", result2)
	//
	//	fmt.Printf("type is %s", reflect.TypeOf(result2))
	//	fmt.Printf("element is %s", reflect.TypeOf(result2).Elem())
	//}

	var result3 map[string]string
	err3 := json.Unmarshal([]byte(test3), &result3)
	if err3 != nil {
		fmt.Printf("Error: %s", err3.Error())
	} else {
		fmt.Printf("Struct: %+v", result3)

		fmt.Printf("type is %s", reflect.TypeOf(result3))
		fmt.Printf("element is %s", reflect.TypeOf(result3).Elem())
	}

}


