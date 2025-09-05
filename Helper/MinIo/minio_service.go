package service

import (
	"context"
	"fmt"
	"mime"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var MinioClient *minio.Client
var bucketName string

// Initialize MinIO client (call this once in main.go)
func InitMinioClient() error {
	endpoint := os.Getenv("MINIO_ENDPOINT")
	portStr := os.Getenv("MINIO_PORT")
	useSSL := os.Getenv("MINIO_USE_SSL") == "true"
	accessKey := os.Getenv("MINIO_ACCESS_KEY")
	secretKey := os.Getenv("MINIO_SECRET_KEY")
	bucketName = os.Getenv("MINIO_BUCKET")

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return fmt.Errorf("invalid MINIO_PORT: %v", err)
	}
	fullEndpoint := fmt.Sprintf("%s:%d", endpoint, port)

	client, err := minio.New(fullEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return fmt.Errorf("failed to initialize MinIO client: %v", err)
	}

	MinioClient = client
	fmt.Println("✅ MinIO client initialized")

	// Ensure bucket exists
	ctx := context.Background()
	exists, err := client.BucketExists(ctx, bucketName)
	if err != nil {
		return fmt.Errorf("error checking bucket: %v", err)
	}
	if !exists {
		err = client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: "us-east-1"})
		if err != nil {
			return fmt.Errorf("failed to create bucket: %v", err)
		}
		fmt.Printf("✅ Bucket %s created\n", bucketName)
	}
	return nil
}

// Generate a presigned PUT (upload) URL
func CreateUploadURL(fileName string, expireMins int) (uploadUrl string, fileUrl string, err error) {
	if MinioClient == nil {
		return "", "", fmt.Errorf("Minio client is not initialized, call InitMinioClient() first")
	}

	ctx := context.Background()
	expiry := time.Duration(expireMins) * time.Minute

	uploadUrlObj, err := MinioClient.PresignedPutObject(ctx, bucketName, fileName, expiry)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate upload URL: %w", err)
	}
	uploadUrl = uploadUrlObj.String()

	fileUrl, err = GetFileURL(fileName, expireMins)
	if err != nil {
		return "", "", err
	}

	return uploadUrl, fileUrl, nil
}

func GetFileURL(fileName string, expireMins int) (string, error) {
	if MinioClient == nil {
		return "", fmt.Errorf("Minio client is not initialized, call InitMinioClient() first")
	}

	ctx := context.Background()
	expiry := time.Duration(expireMins) * time.Minute

	url, err := MinioClient.PresignedGetObject(ctx, bucketName, fileName, expiry, nil)
	if err != nil {
		return "", fmt.Errorf("failed to generate file URL: %w", err)
	}
	return url.String(), nil
}

func DeleteFile(fileName string) error {
	if MinioClient == nil {
		return fmt.Errorf("Minio client is not initialized, call InitMinioClient() first")
	}

	ctx := context.Background()
	err := MinioClient.RemoveObject(ctx, bucketName, fileName, minio.RemoveObjectOptions{})
	if err != nil {
		return fmt.Errorf("failed to delete file %s: %w", fileName, err)
	}

	fmt.Printf("✅ File %s deleted successfully\n", fileName)
	return nil
}

// Generate a presigned GET (download) URL
// func CreateDownloadURL(fileName string, expireMins int) (string, error) {
// 	if MinioClient == nil {
// 		return "", fmt.Errorf("Minio client is not initialized, call InitMinioClient() first")
// 	}

// 	ctx := context.Background()
// 	expiry := time.Duration(expireMins) * time.Minute

// 	// Use url.Values to set headers
// 	reqParams := make(url.Values)
// 	reqParams.Set("response-content-disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileName))

// 	downloadURL, err := MinioClient.PresignedGetObject(ctx, bucketName, fileName, expiry, reqParams)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to generate download URL: %w", err)
// 	}

// 	return downloadURL.String(), nil
// }

func CreateDownloadURL(objectName, downloadName string, expireMins int) (string, error) {
	if MinioClient == nil {
		return "", fmt.Errorf("Minio client is not initialized, call InitMinioClient() first")
	}

	ctx := context.Background()
	expiry := time.Duration(expireMins) * time.Minute

	// Use url.Values to set custom response headers
	reqParams := make(url.Values)
	// Force browser to download with given name
	reqParams.Set("response-content-disposition", fmt.Sprintf("attachment; filename=\"%s\"", downloadName))

	downloadURL, err := MinioClient.PresignedGetObject(ctx, bucketName, objectName, expiry, reqParams)
	if err != nil {
		return "", fmt.Errorf("failed to generate download URL: %w", err)
	}

	return downloadURL.String(), nil
}

func GetFileInfo(objectName string) (string, error) {
	if MinioClient == nil {
		return "", fmt.Errorf("Minio client is not initialized, call InitMinioClient() first")
	}

	ctx := context.Background()

	// Get metadata
	stat, err := MinioClient.StatObject(ctx, bucketName, objectName, minio.StatObjectOptions{})
	if err != nil {
		return "", fmt.Errorf("failed to stat object: %w", err)
	}

	// ---- File size formatting ----
	size := stat.Size
	var sizeStr string
	if size > 1024*1024 {
		sizeStr = fmt.Sprintf("%.1f MB", float64(size)/(1024*1024))
	} else if size > 1024 {
		sizeStr = fmt.Sprintf("%.1f KB", float64(size)/1024)
	} else {
		sizeStr = fmt.Sprintf("%d B", size)
	}

	// ---- File type (extension or content-type fallback) ----
	ext := strings.ToUpper(strings.TrimPrefix(filepath.Ext(objectName), "."))
	if ext == "" {
		// fallback to content-type
		exts, _ := mime.ExtensionsByType(stat.ContentType)
		if len(exts) > 0 {
			ext = strings.ToUpper(strings.TrimPrefix(exts[0], "."))
		} else {
			ext = "FILE"
		}
	}

	return fmt.Sprintf("%s • %s", ext, sizeStr), nil
}
