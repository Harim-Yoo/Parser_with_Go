package main

import (
	"log"
	"os"
	"parser/internal"
	"sync"
)

func main() {
	files, err := os.ReadDir("./folder")
	if err != nil {
		log.Printf("Err1:%v\n", err)
		return
	}
	var wg sync.WaitGroup

	for _, file := range files {
		wg.Add(1)
		go func(fileName string) {
			defer wg.Done()
			data := internal.Trimmer(fileName)
			err := internal.UpdateBucketsToFile(fileName, data)
			if err != nil {
				log.Printf("Err4:%v\n", err)
				return
			}
		}(file.Name())
	}
	wg.Wait()

	log.Printf("We are all done here...")
}
