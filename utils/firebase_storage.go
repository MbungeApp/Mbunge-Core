package utils

import (
	cloud "cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/api/option"
	"io"
	"mime/multipart"
	"os"
)

var (
	accessJson = "../config/credentials.json"
	bucket     = "mbungeapp.appspot.com"
)

func UploadFile(file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	// Destination
	err = os.Chdir("./uploads")
	if err != nil {
		return "", err
	}
	dst, err := os.Create(file.Filename)
	if err != nil {
		return "", err
	}
	defer dst.Close()
	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}

	imgUrl, err := firestoreStorageService(fmt.Sprintf("../uploads/%s", file.Filename), file.Filename)
	if err != nil {
		return "", err
	}
	defer os.Remove(file.Filename)

	return imgUrl, nil
}

func firestoreStorageService(filePath string, imagePath string) (string, error) {
	ctx := context.Background()

	f, err := os.Open(filePath)

	if err != nil {
		return "", err
	}

	defer f.Close()

	clientOptions := option.WithCredentialsFile(accessJson)

	storage, err := cloud.NewClient(ctx, clientOptions)
	if err != nil {
		return "", err
	}

	wc := storage.Bucket(bucket).Object(imagePath).NewWriter(ctx)

	uniqueId := uuid.New()
	wc.Metadata = map[string]string{
		"firebaseStorageDownloadTokens": uniqueId.String(),
	}
	_, err = io.Copy(wc, f)
	if err != nil {
		return "", err
	}

	if err := wc.Close(); err != nil {
		return "", err
	}
	imageUrl := fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%s/o/%s?alt=media&token=%s", bucket, imagePath, uniqueId.String())
	return imageUrl, nil
}
