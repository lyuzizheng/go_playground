package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	ElasticSearch struct {
	  ID string `yaml:"id"`
	  Password string `yaml:"password"`
	} `yaml:"elastic_search"`
}

var GlobalConfig Config


func init() {

	yamlFile, err := os.ReadFile("config.yml")
    if err != nil {
        log.Fatalf("yamlFile.Get err   #%v ", err)
    }
    err = yaml.Unmarshal(yamlFile, &GlobalConfig)
    if err != nil {
        log.Fatalf("Unmarshal: %v", err)
    }
	
}