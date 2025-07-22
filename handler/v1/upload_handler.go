package v1

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/labstack/echo/v4"
)

type UploadHandler struct {
	S3Client *s3.Client
	Bucket   string
}

func NewUploadHandler(e *echo.Echo, bucket string) (*UploadHandler, error) {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return nil, err
	}
	s3Client := s3.NewFromConfig(cfg)
	h := &UploadHandler{
		S3Client: s3Client,
		Bucket:   bucket,
	}
	e.POST("/api/v1/upload", h.Upload)
	return h, nil
}

// Upload godoc
// @Summary 上传文件到 S3
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "文件"
// @Success 200 {object} map[string]string
// @Router /api/v1/upload [post]
func (h *UploadHandler) Upload(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "missing file"})
	}
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "open file failed"})
	}
	defer src.Close()

	key := file.Filename
	_, err = h.S3Client.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket: aws.String(h.Bucket),
		Key:    aws.String(key),
		Body:   src,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": fmt.Sprintf("upload failed: %v", err)})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "upload success", "key": key})
}
