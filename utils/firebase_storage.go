package utils

import (
	cloud "cloud.google.com/go/storage"
	"context"
	"fmt"
	uuid "github.com/google/uuid"
	"google.golang.org/api/option"
	"io"
	"os"
)

var (
	accessJson = "config/credentials.json"
	bucket     = "mbungeapp.appspot.com"
)

func UploadFile(filePath string) (string, error) {
	ctx := context.Background()

	clientOptions := option.WithCredentialsFile(accessJson)

	storage, err := cloud.NewClient(ctx, clientOptions)
	if err != nil {
		return "", err
	}
	f, err := os.Open(filePath)

	if err != nil {
		return "", err
	}

	defer f.Close()

	imagePath := f.Name()

	wc := storage.Bucket(bucket).Object(imagePath).NewWriter(ctx)

	uuid := uuid.New()
	wc.Metadata = map[string]string{
		"firebaseStorageDownloadTokens": uuid.String(),
	}
	_, err = io.Copy(wc, f)
	if err != nil {
		return "", err
	}

	if err := wc.Close(); err != nil {
		return "", err
	}

	if err != nil {
		return "", err
	}
	imageUrl := fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%s/o/%s?alt=media&token=%s", bucket, imagePath, uuid.String())
	return imageUrl, nil
}
