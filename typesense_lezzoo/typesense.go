package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/typesense/typesense-go/typesense"
	typesenseAPI "github.com/typesense/typesense-go/typesense/api"
)

func ImportSplitFiles(pattern string) error {
	// 获取所有拆分文件
	files, err := filepath.Glob(pattern)
	if err != nil {
		return fmt.Errorf("查找文件失败: %w", err)
	}

	// 按文件名排序确保顺序导入
	sort.Strings(files)

	var totalSuccess, totalError int

	fmt.Printf("找到 %d 个文件待导入\n", len(files))

	// 依次处理每个文件
	for i, file := range files {
		fmt.Printf("\n处理第 %d/%d 个文件: %s\n", i+1, len(files), file)

		err := ImportTypesenseNew(file)
		if err != nil {
			fmt.Printf("导入文件 %s 失败: %v\n", file, err)
			continue
		}
	}

	fmt.Printf("\n导入完成!\n总成功: %d\n总失败: %d\n", totalSuccess, totalError)
	return nil
}

func ImportTypesense(fileName string) error {
	client := typesense.NewClient(
		typesense.WithServer("http://192.168.100.199:8108"),
		typesense.WithAPIKey("****"),
	)

	truePtr := true
	productId := "product_id"
	schema := &typesenseAPI.CollectionSchema{
		Name: "app-search",
		Fields: []typesenseAPI.Field{
			{Name: "product_id", Type: "int64"},
			{Name: "merchant_id", Type: "string", Facet: &truePtr},
			{Name: "name", Type: "string", Index: &truePtr},
			{Name: "product_name", Type: "string", Index: &truePtr},
			{Name: "product_stock", Type: "string", Facet: &truePtr},
			{Name: "product_status", Type: "string", Facet: &truePtr},
			{Name: "merchant_name", Type: "string", Index: &truePtr},
			{Name: "city", Type: "string", Facet: &truePtr},
			{Name: "merchant_status", Type: "string", Facet: &truePtr},
			{Name: "category_id", Type: "string", Facet: &truePtr},
			{Name: "category_ids", Type: "string[]", Facet: &truePtr},
			{Name: "deaprtment_ids", Type: "string[]", Facet: &truePtr},
			{Name: "tags", Type: "string[]", Facet: &truePtr, Index: &truePtr},
			{Name: "vertical", Type: "string[]", Facet: &truePtr, Index: &truePtr},
			{Name: "is_pickup", Type: "string", Facet: &truePtr},
			{Name: "merchant", Type: "string"},
			{Name: "product", Type: "string"},
		},
		DefaultSortingField: &productId,
	}

	ctx := context.Background()

	fmt.Println("Starting import to Typesense...")

	// Create collection if not exists
	_, err := client.Collections().Create(ctx, schema)
	if err != nil && !strings.Contains(err.Error(), "already exists") {
		return fmt.Errorf("failed to create collection: %w", err)
	}

	// Open and read JSONL file
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	fmt.Println("Opened file for import...")

	// Get collection
	collection := client.Collection("app-search")
	if collection == nil {
		return fmt.Errorf("failed to get collection")
	}

	// Import parameters
	params := &typesenseAPI.ImportDocumentsParams{
		Action:    stringPtr("create"),
		BatchSize: intPtr(40),
	}

	fmt.Println("Starting import...")

	// Perform import
	result, err := collection.Documents().ImportJsonl(ctx, file, params)
	if err != nil {
		return fmt.Errorf("import failed: %w", err)
	}
	defer result.Close()

	fmt.Println("Import completed. Processing results...")

	// Process results
	scanner := bufio.NewScanner(result)
	var successCount, errorCount int
	for scanner.Scan() {
		var result = struct {
			Code    int    `json:"code"`
			Error   string `json:"error"`
			Success string `json:"success"`
		}{}
		_ = json.Unmarshal(scanner.Bytes(), &result)

		if result.Code == 0 {
			successCount++
		} else {
			errorCount++
			line, _ := json.Marshal(result)
			fmt.Println(string(line))
		}
	}

	fmt.Printf("Import completed. Successful: %d, Failed: %d\n", successCount, errorCount)
	return nil
}

func CreateSchema() error {
	client := typesense.NewClient(
		typesense.WithServer("https://1qnohjb47fpzycgup-1.a1.typesense.net:443"),
		typesense.WithAPIKey("4ybOZTuQuQOFsVPqsaz6KdGAHPt8B1Ka"),
	)
	productId := "product_id"
	truePtr := true
	falsePtr := false
	schema := &typesenseAPI.CollectionSchema{
		Name: "app-search",
		Fields: []typesenseAPI.Field{
			{Name: "id", Type: "string"},
			{Name: "merchant", Type: "string", Index: &falsePtr},
			{Name: "product", Type: "string", Index: &falsePtr},
			{Name: "name", Type: "string", Index: &truePtr},
			{Name: "product_name", Type: "string", Index: &truePtr},
			{Name: "merchant_name", Type: "string", Index: &truePtr},
			{Name: "city", Type: "string", Facet: &truePtr},
			{Name: "merchant_status", Type: "string", Facet: &truePtr},
			{Name: "product_status", Type: "int32", Facet: &truePtr},
			{Name: "vertical", Type: "string[]", Facet: &truePtr},
			{Name: "tags", Type: "string[]", Facet: &truePtr, Index: &truePtr},
			{Name: "merchant_id", Type: "int32", Facet: &truePtr},
			{Name: "is_pickup", Type: "int32", Facet: &truePtr},
			{Name: "category_id", Type: "int32", Facet: &truePtr},
			{Name: "deaprtment_ids", Type: "string[]", Facet: &truePtr},
			{Name: "category_ids", Type: "int32[]", Facet: &truePtr},
			{Name: "product_stock", Type: "int32", Facet: &truePtr},
			{Name: "product_id", Type: "int32", Index: &falsePtr},
		},
		DefaultSortingField: &productId,
	}

	ctx := context.Background()

	fmt.Println("Starting creating schema to Typesense...")

	// Create collection if not exists
	result, err := client.Collections().Create(ctx, schema)
	if err != nil && !strings.Contains(err.Error(), "already exists") {
		return fmt.Errorf("failed to create collection: %w", err)
	}
	fmt.Printf("Schema created: %+v\n", result)
	return nil

}

func ImportTypesenseNew(fileName string) error {
	client := typesense.NewClient(
		typesense.WithServer("https://1qnohjb47fpzycgup-1.a1.typesense.net:443"),
		typesense.WithAPIKey("4ybOZTuQuQOFsVPqsaz6KdGAHPt8B1Ka"),
	)

	productId := "product_id"
	truePtr := true
	schema := &typesenseAPI.CollectionSchema{
		Name: "app-search-new",
		Fields: []typesenseAPI.Field{
			{Name: "id", Type: "string"},
			{Name: "merchant", Type: "string"},
			{Name: "product", Type: "string"},
			{Name: "name", Type: "string", Index: &truePtr},
			{Name: "product_name", Type: "string", Index: &truePtr},
			{Name: "merchant_name", Type: "string", Index: &truePtr},
			{Name: "city", Type: "string", Facet: &truePtr},
			{Name: "merchant_status", Type: "string", Facet: &truePtr},
			{Name: "product_status", Type: "string", Facet: &truePtr},
			{Name: "vertical", Type: "string[]", Facet: &truePtr},
			{Name: "tags", Type: "string[]", Facet: &truePtr, Index: &truePtr},
			{Name: "merchant_id", Type: "string", Facet: &truePtr},
			{Name: "is_pickup", Type: "int32", Facet: &truePtr},
			{Name: "category_id", Type: "int32", Facet: &truePtr},
			{Name: "deaprtment_ids", Type: "string[]", Facet: &truePtr},
			{Name: "category_ids", Type: "string[]", Facet: &truePtr},
			{Name: "product_stock", Type: "int32", Facet: &truePtr},
			{Name: "product_id", Type: "int32"},
		},
		DefaultSortingField: &productId,
	}

	ctx := context.Background()

	fmt.Println("Starting import to Typesense...")

	// Create collection if not exists
	_, err := client.Collections().Create(ctx, schema)
	if err != nil && !strings.Contains(err.Error(), "already exists") {
		return fmt.Errorf("failed to create collection: %w", err)
	}

	// Open and read JSONL file
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	fmt.Println("Opened file for import...")

	// Get collection
	collection := client.Collection("test")
	if collection == nil {
		return fmt.Errorf("failed to get collection")
	}

	// Import parameters
	params := &typesenseAPI.ImportDocumentsParams{
		Action:    stringPtr("create"),
		BatchSize: intPtr(40),
	}

	fmt.Println("Starting import...")

	// Perform import
	result, err := collection.Documents().ImportJsonl(ctx, file, params)
	if err != nil {
		return fmt.Errorf("import failed: %w", err)
	}
	defer result.Close()

	fmt.Println("Import completed. Processing results...")

	// Process results
	scanner := bufio.NewScanner(result)
	var successCount, errorCount int
	for scanner.Scan() {
		var result = struct {
			Code    int    `json:"code"`
			Error   string `json:"error"`
			Success string `json:"success"`
		}{}
		_ = json.Unmarshal(scanner.Bytes(), &result)

		if result.Code == 0 {
			successCount++
		} else {
			errorCount++
			line, _ := json.Marshal(result)
			fmt.Println(string(line))
		}
	}

	fmt.Printf("Import completed. Successful: %d, Failed: %d\n", successCount, errorCount)
	return nil
}

func stringPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}

func splitJSONL(inputFile string, batchSize int) error {
	file, err := os.Open(inputFile)
	if err != nil {
		return fmt.Errorf("打开文件失败: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fileCount := 1
	lineCount := 0
	totalCount := 0

	// 准备第一个输出文件
	currentFile, writer, err := createOutputFile(fileCount)
	if err != nil {
		return err
	}
	defer currentFile.Close()

	fmt.Println("开始分割文件...")

	for scanner.Scan() {
		// 写入当前行
		if _, err := writer.WriteString(scanner.Text() + "\n"); err != nil {
			return fmt.Errorf("写入数据失败: %w", err)
		}

		lineCount++
		totalCount++

		// 达到批次大小，创建新文件
		if lineCount >= batchSize {
			currentFile.Close()
			fileCount++
			lineCount = 0

			fmt.Printf("已处理 %d 条记录，创建第 %d 个文件\n", totalCount, fileCount)

			currentFile, writer, err = createOutputFile(fileCount)
			if err != nil {
				return err
			}
		}
	}

	fmt.Printf("\n分割完成!\n总记录数: %d\n生成文件数: %d\n", totalCount, fileCount)
	return scanner.Err()
}

func createOutputFile(fileNum int) (*os.File, *bufio.Writer, error) {
	filename := fmt.Sprintf("./split/split_data_%04d.jsonl", fileNum)

	file, err := os.Create(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("创建文件失败: %w", err)
	}

	writer := bufio.NewWriter(file)

	return file, writer, nil
}
