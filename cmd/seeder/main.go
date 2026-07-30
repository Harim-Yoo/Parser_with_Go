package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"parser/internal/database"
	"path/filepath"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

type DataStruct database.InsertDataRow

func main() {

	files, err := os.ReadDir("./folder")
	if err != nil {
		log.Printf("Err1:%v\n", err)
		return
	}
	var wg sync.WaitGroup

	db, err := sql.Open("postgres", "postgres://postgres:reachwise@localhost:5432/fivehundred")

	if err != nil {
		log.Printf("Err2:%v\n", err)
		return
	}

	defer db.Close()
	dbQueries := database.New(db)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var Results []database.InsertDataRow
	for _, file := range files {
		wg.Add(1)
		go func(fileName string) {
			defer wg.Done()

			data, err := os.ReadFile(filepath.Join("./folder", fileName))
			if err != nil {
				log.Printf("Err:%v\n", err)
			}
			insData, err := dbQueries.InsertData(ctx, database.InsertDataParams{
				Title:    fileName,
				Contents: string(data),
			})
			if err != nil {
				return
			}
			Results = append(Results, insData...)
		}(file.Name())

	}
	wg.Wait()
	log.Println("All work is done.")
	for _, result := range Results {
		fmt.Printf("Title:%v\n", result.Title)
	}
	fmt.Printf("Total:%v\n", len(Results))

}
