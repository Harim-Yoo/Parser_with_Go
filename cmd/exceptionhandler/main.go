package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	files, err := os.ReadDir("./folder")
	if err != nil {
		log.Printf("Err1:%v\n", err)
		return
	}
	count := 0
	for i := 0; i < len(files); i++ {
		data, err := os.ReadFile(filepath.Join("./folder", files[i].Name()))
		if err != nil {
			log.Printf("Err2:%v\n", err)
			return
		}
		n := bytes.Index([]byte(files[i].Name()), []byte("_"))
		fileName := files[i].Name()[n+1:]
		fileName = strings.Join([]string{"**", strings.TrimSuffix(fileName, ".md"), "**"}, "")

		m := bytes.Index(data, []byte(fileName))
		if m != 0 {
			count++
			fmt.Printf("Here is the file:%v\n", files[i].Name())
		}
	}
	fmt.Printf("Here are the counts:%v\n", count)
}
