package services

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
)

func SavePDF(file *multipart.FileHeader, uploadPath string) (string, error) {

	// Create uploads directory if it doesn't exist
	err := os.MkdirAll(uploadPath, os.ModePerm)
	if err != nil {
		return "", err
	}

	// Destination path
	dst := filepath.Join(uploadPath, file.Filename)

	// Open uploaded file
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Create destination file
	out, err := os.Create(dst)
	if err != nil {
		return "", err
	}
	defer out.Close()

	// Copy uploaded data
	_, err = out.ReadFrom(src)
	if err != nil {
		return "", err
	}

	fmt.Println("PDF Saved:", dst)

	return file.Filename, nil
}
