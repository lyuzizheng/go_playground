package main

import "fmt"

func main() {
	// ImportTypesense("transformed_data.jsonl")
	//splitJSONL("transformed_data.jsonl", 1000)

	// err := ImportSplitFiles("./transformed_data_burger.jsonl")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	err := CreateSchemaPopularQueries()
	if err != nil {
		fmt.Println(err)
	}
	err = CreateSchemaNoHit()
	if err != nil {
		fmt.Println(err)
	}
}
