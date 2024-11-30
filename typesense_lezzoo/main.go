package main

import "log"

func main() {
	// ImportTypesense("transformed_data.jsonl")
	//splitJSONL("transformed_data.jsonl", 1000)

	err := ImportSplitFiles("./split/split_data_*.jsonl")
	if err != nil {
		log.Fatal(err)
	}
}
