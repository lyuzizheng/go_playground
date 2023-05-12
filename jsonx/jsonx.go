package jsonx

import (
	"code.byted.org/gopkg/jsonx"
	"fmt"
)

func TryJson() {
	any := jsonx.GetFromString(attrResultDiversifyID, "diversify_id")
	any2 := jsonx.GetFromString(attrResultDiversifyID, "mt_diversity_tier3_p90")
	fmt.Println(any.GetString())
	fmt.Println(any2.Data())

	//ant3, _ := jsonx.GetFromString(any2.Data().(string)).Map()
	//fmt.Println(ant3)
	//any2Str, _ := any2.MapInterface()
	//fmt.Println(any2Str)
	//
	//any3 := jsonx.GetFromString(myStr, "mt_diversity_tier3_top1")
	//if !any3.Exists() {
	//	fmt.Println("hahahha")
	//}
	//p90 := make(map[string]string)
	//_ = jsonx.UnmarshalFromString(any2Str, p90)
	//fmt.Println(p90)
}
