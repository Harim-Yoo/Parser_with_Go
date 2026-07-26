package internal

import (
	"bytes"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	htmltomarkdown "github.com/JohannesKaufmann/html-to-markdown/v2"
)

func Reader(dirName string) ([]byte, error) {
	data, err := os.ReadFile(dirName)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

func Parser(data []byte, separater string) ([][]byte, error) {
	buffer := make([][]byte, 0)
	var bucket []byte
	for {
		n := bytes.Index(data, []byte(separater))
		if n == -1 {
			buffer = append(buffer, data)
			break
		}
		bucket = bytes.TrimSpace(data[:n])
		buffer = append(buffer, bucket)
		data = data[n+1:]
	}
	return buffer, nil
}

func HtmlToMarkdown(filePath io.Reader) (string, error) {

	markdown, err := htmltomarkdown.ConvertReader(filePath)
	if err != nil {
		return "", err
	}
	return string(markdown), nil
}

func StringToFile(fileName, fileContent string) error {
	trimmedName := strings.TrimSuffix(fileName, ".html")
	newFileName := trimmedName + ".md"
	filePath := filepath.Join("./folder", newFileName)
	err := os.WriteFile(filePath, []byte(fileContent), 0644)
	if err != nil {
		log.Printf("Err:%v\n", err)
		return err
	}
	return nil
}
