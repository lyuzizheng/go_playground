package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/typesense/typesense-go/typesense"
	typesenseAPI "github.com/typesense/typesense-go/typesense/api"
)

var client = typesense.NewClient(
	typesense.WithServer("https://1qnohjb47fpzycgup-1.a1.typesense.net:443"),
	typesense.WithAPIKey("4ybOZTuQuQOFsVPqsaz6KdGAHPt8B1Ka"),
)

func CreateSchemaPopularQueries() error {

	schema := &typesenseAPI.CollectionSchema{
		Name: "app_search_popular_queries",
		Fields: []typesenseAPI.Field{
			{Name: "q", Type: "string"},
			{Name: "count", Type: "int32"},
		},
	}
	ctx := context.Background()
	fmt.Println("Starting creating schema to Typesense...")

	// Create collection if not exists
	result, err := client.Collections().Create(ctx, schema)
	if err != nil && !strings.Contains(err.Error(), "already exists") {
		return fmt.Errorf("failed to create collection: %w", err)
	}
	fmt.Printf("PopularQueries Schema created: %+v\n", result)
	return nil

}

func CreateSchemaNoHit() error {

	schema := &typesenseAPI.CollectionSchema{
		Name: "app_search_no_hit",
		Fields: []typesenseAPI.Field{
			{Name: "q", Type: "string"},
			{Name: "count", Type: "int32"},
		},
	}
	ctx := context.Background()
	fmt.Println("Starting creating schema to Typesense...")

	// Create collection if not exists
	result, err := client.Collections().Create(ctx, schema)
	if err != nil && !strings.Contains(err.Error(), "already exists") {
		return fmt.Errorf("failed to create collection: %w", err)
	}
	fmt.Printf("NoHit Schema created: %+v\n", result)
	return nil

}

func CreateSchemaPopularity() error {

	schema := &typesenseAPI.CollectionSchema{
		Name: "app_search_item_popularity",
		Fields: []typesenseAPI.Field{
			{Name: "q", Type: "string"},
			{Name: "count", Type: "int32"},
		},
	}
	ctx := context.Background()
	fmt.Println("Starting creating schema to Typesense...")

	// Create collection if not exists
	result, err := client.Collections().Create(ctx, schema)
	if err != nil && !strings.Contains(err.Error(), "already exists") {
		return fmt.Errorf("failed to create collection: %w", err)
	}
	fmt.Printf("PopularQueries Schema created: %+v\n", result)
	return nil

}
