package main

import (
	"log"
	"os"
	"parser/internal"
	"path/filepath"
	"sync"
)

func main() {

	files, err := os.ReadDir("./contents")

	if err != nil {
		log.Printf("Err0:%v\n", err)
		return
	}

	var wg sync.WaitGroup

	for _, entry := range files {
		wg.Add(1)
		go func(fileEntry os.DirEntry) {
			defer wg.Done()

			file, err := os.Open(filepath.Join("./contents", fileEntry.Name()))
			if err != nil {
				log.Printf("Err1:%v\n", err)
				return
			}
			defer file.Close()

			results, err := internal.HtmlToMarkdown(file)
			if err != nil {
				log.Printf("Err2:%v\n", err)
				return
			}
			err = internal.StringToFile(fileEntry.Name(), results)
			if err != nil {
				log.Printf("Err3:%v\n", err)
				return
			}
		}(entry)
	}
	wg.Wait()

}
