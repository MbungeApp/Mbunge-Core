package utils

import (
	cloud "cloud.google.com/go/storage"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/api/option"
	"io"
	"os"
	"strings"
)

var (
	accessJson = "../config/credentials.json"
	bucket     = "mbungeapp.appspot.com"
)

func UploadFile(fileString string) (string, error) {
	///
	unique := uuid.New()
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	if !strings.Contains(path, "/uploads") {
		err = os.Chdir("./uploads")
		if err != nil {
			return "", err
		}
	}

	dec, err := base64.StdEncoding.DecodeString(fileString)
	if err != nil {
		return "", err
	}
	f, err := os.Create(fmt.Sprintf("%s.png", unique.String()))
	if err != nil {
		return "", err
	}
	defer f.Close()
	if _, err := f.Write(dec); err != nil {
		return "", err
	}
	if err := f.Sync(); err != nil {
		return "", err
	}

	src, err := os.Open(f.Name())
	if err != nil {
		return "", err
	}
	defer src.Close()

	imgUrl, err := firestoreStorageService(fmt.Sprintf("../uploads/%s", f.Name()), f.Name())
	if err != nil {
		return "", err
	}
	defer os.Remove(f.Name())

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
