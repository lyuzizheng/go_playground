package bandlabtest

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"playground/cloudflare"
)

func TryUploadImage() {
	err := cloudflare.Initialize(cloudflare.Config{
		AccessKey:  "7ff1f4515c58056eca10e76c387fbbb9",
		SecretKey:  "fd1226d104f2586ff4293f9992cc0803c35edd8fbdcb787dacdeda3d56874d22",
		AccountID:  "bf69da5249b63731ad79545d0095e8db",
		BucketName: "bandlab-assignment",
	})
	if err != nil {
		log.Fatalf("Failed to initialize R2: %v", err)
	}
	// 读取 bandlab_test 文件夹中的图像文件
	imagePath := "bandlab_test/lalal.jpg" // 请根据实际文件名修改
	file, err := os.Open(imagePath)
	if err != nil {
		log.Fatalf("Failed to open image file: %v", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatalf("Failed to get file info: %v", err)
	}

	fileSize := fileInfo.Size()
	fileBuffer := make([]byte, fileSize)
	_, err = file.Read(fileBuffer)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// 计算文件的 MD5 哈希值
	hash := md5.New()
	if _, err := io.Copy(hash, bytes.NewReader(fileBuffer)); err != nil {
		log.Fatalf("Failed to calculate MD5: %v", err)
	}
	md5Hash := hash.Sum(nil)
	base64MD5 := base64.StdEncoding.EncodeToString(md5Hash)

	// 打印图像的元数据
	mimeType := mime.TypeByExtension(filepath.Ext(fileInfo.Name()))
	fmt.Printf("File Name: %s\n", fileInfo.Name())
	fmt.Printf("File Size: %d bytes\n", fileSize)
	fmt.Printf("MIME Type: %s\n", mimeType)
	fmt.Printf("MD5: %x\n", md5Hash)

	url, err := cloudflare.GeneratePresignedURL(fileInfo.Name(), fileSize, mimeType, base64MD5)
	if err != nil {
		log.Fatalf("Failed to generate presigned URL: %v", err)
	}
	log.Printf("Presigned URL: %s", url)

	// 使用预签名 URL 上传图像
	req, err := http.NewRequest("PUT", url, bytes.NewReader(fileBuffer))
	if err != nil {
		log.Fatalf("Failed to create HTTP request: %v", err)
	}
	req.Header.Set("Content-Type", mimeType)
	req.Header.Set("Content-Length", fmt.Sprintf("%d", fileSize))
	req.Header.Set("Content-MD5", base64MD5)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to upload image: %v", err)
	}
	defer resp.Body.Close()

	// 打印上传结果
	fmt.Printf("Upload Status: %s\n", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}
	fmt.Printf("Response Body: %s\n", string(body))
}
