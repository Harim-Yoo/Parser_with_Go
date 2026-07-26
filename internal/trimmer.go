package internal

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Trimmer(fileName string) []byte {

	data, err := os.ReadFile(filepath.Join("./folder", fileName))

	if err != nil {
		log.Printf("Err2:%v\n", err)
		return []byte{}
	}

	idx := bytes.Index([]byte(fileName), []byte("_"))
	fileName = fileName[idx+1:]
	fileName = strings.TrimSuffix(fileName, ".md")
	fileName = strings.Join([]string{"**", fileName, "**"}, "")
	fmt.Printf("%v\n", fileName)
	n := bytes.Index(data, []byte(fileName))
	if n == -1 {
		return data
	}
	return data[n:]
}

func UpdateBucketsToFile(fileName string, fileContent []byte) error {

	filePath := filepath.Join("./folder", fileName)
	err := os.WriteFile(filePath, []byte(fileContent), 0644)
	if err != nil {
		log.Printf("Err3:%v\n", err)
		return err
	}
	return nil
}
