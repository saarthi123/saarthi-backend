package utils

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var S3Client *s3.Client
var BucketName string

func InitAWS() {
	BucketName = os.Getenv("SAARTHI_DATA_BUCKET")
	if BucketName == "" {
		log.Fatal("❌ SAARTHI_DATA_BUCKET not set in environment")
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("❌ Failed to load AWS config: %v", err)
	}

	S3Client = s3.NewFromConfig(cfg)
	log.Println("✅ AWS S3 Initialized")
}
