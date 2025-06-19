package handlers

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/gin-gonic/gin"
	"github.com/saarthi123/saarthi-backend/utils"
)

func UploadHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	f, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
		return
	}
	defer f.Close()

	buf := new(bytes.Buffer)
	io.Copy(buf, f)

	key := "uploads/" + time.Now().Format("20060102150405") + "-" + file.Filename
	contentType := file.Header.Get("Content-Type")

	_, err = utils.S3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      &utils.BucketName,
		Key:         &key,
		Body:        bytes.NewReader(buf.Bytes()),
		ContentType: &contentType,
		ACL:         types.ObjectCannedACLPublicRead, // Or private
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "S3 upload failed", "details": err.Error()})
		return
	}

	url := "https://" + utils.BucketName + ".s3.amazonaws.com/" + key
	c.JSON(http.StatusOK, gin.H{"url": url})
}
