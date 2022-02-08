package json

import (
	"encoding/json"
	"fmt"
	"reflect"

)

type Card interface {
	ToCard(defaultOpenChatId string) interface{}
}

type CardTemplate1 struct {
	OpenId     string `json:"open_id,omitempty"`
	UserId     string `json:"user_id,omitempty"`
	OpenChatId string `json:"open_chat_id"`

	//Compulsory
	Title      string               `json:"title"`
	TitleColor string `json:"title_color"`
	//Compulsory
	Message string `json:"message"`
	//Compulsory
	FootNote string              `json:"foot_note"`
	Extra  []string `json:"extra"`
}

func (c CardTemplate1) ToCard(defaultOpenChatID string) interface{} {

	return nil

}

type CardTemplateDefault struct {
	OpenId     string `json:"open_id,omitempty"`
	UserId     string `json:"user_id,omitempty"`
	OpenChatId string `json:"open_chat_id"`

	//Compulsory
	Message string `json:"message"`
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
		fmt.Println(len(cardNew.Extra))
	default:
		cardNew := CardTemplateDefault{}
		err = json.Unmarshal([]byte(rawCard), &cardNew)
		card = cardNew
	}
	//err = json.Unmarshal([]byte(rawCard), &card)
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


	//var result interfaces{}
	//err := json.Unmarshal([]byte(test1), &result)
	//if err != nil {
	//	fmt.Printf("Error: %s", err.Error())
	//} else {
	//	fmt.Printf("Struct: %+v", result)
	//	fmt.Printf("type is %s", reflect.TypeOf(result))
	//}
	//var result2 interfaces{}
	//err2 := json.Unmarshal([]byte(test2), &result2)
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


