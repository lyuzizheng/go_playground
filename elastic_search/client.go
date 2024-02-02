package elasticsearch

import (
	"fmt"
	"playground/config"

	elasticsearch "github.com/elastic/go-elasticsearch/v8"
)

var client *elasticsearch.Client

func Init() {
	var err error
	client, err = elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{config.GlobalConfig.ElasticSearch.Host},
		APIKey:    config.GlobalConfig.ElasticSearch.ApiKey,
	})

	if err != nil {
		fmt.Printf("Error creating the client: %s\n", err)
		return
	}

	resp, err := client.Nodes.Info()
	if err != nil {
		fmt.Println("err querying node info")
		return
	}
	fmt.Println(resp.String())

}
