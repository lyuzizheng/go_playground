package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//getPrize := getAliasFunc(a)
	getPrize := getAliasFunc(a)

	prizes := make(map[string]int)

	for i := 0; i < 1000000; i++ {
		prize := getPrize()
		if item, ok := prizes[prize]; ok {
			prizes[prize] = item + 1
		} else {
			prizes[prize] = 1
		}

	}
	allprizes, _ := json.Marshal(prizes)
	fmt.Println(string(allprizes))
	//collection := make([]int, 0, 100)
	//collection = append(collection, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	//for i, content := range collection {
	//	collection = append(collection, 444)
	//	fmt.Println(i, content)
	//	if i >= 10000 {
	//		break
	//	}
	//}
	//for i := 0; i < len(collection); i++ {
	//	collection = append(collection, 444)
	//	fmt.Println(i, collection[i])
	//	if i >= 10000 {
	//		break
	//	}
	//}

	//
	results := []map[string][]int64{}
	for i := 0; i < 100000; i++ {
		result := DeliverCards([]int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 54}, []string{"A", "B", "C", "D"})
		results = append(results, result)
	}

	testResult := testCards(results)
	fmt.Println("testResult", testResult)

}
