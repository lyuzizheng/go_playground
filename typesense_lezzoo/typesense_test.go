package main

import (
	"os"
	"testing"
)
// TestImportTypesense is the test function
func TestImportTypesense(t *testing.T) {
	// Create a temporary JSONL file for testing
	testData := `{"product_id": 1, "name": "Test Product"}
{"product_id": 2, "name": "Another Product"}`

	tmpfile, err := os.Create("test1.jsonl")
	if err != nil {
		t.Fatal(err)
	}

	if _, err := tmpfile.WriteString(testData); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// Mock the typesense client here if needed
	// Run the import
	if err := ImportTypesense("test1.jsonl"); err != nil {
		t.Errorf("ImportTypesense failed: %v", err)
	}
}
