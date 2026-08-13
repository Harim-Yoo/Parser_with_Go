package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"parser/internal/database"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type DataStruct database.InsertDataRow

func main() {

	godotenv.Load(".env")
	dbURL := os.Getenv("DATABASE_URL")

	files, err := os.ReadDir("./folder")

	if err != nil {
		log.Printf("Err1:%v\n", err)
		return
	}

	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Printf("Err2:%v\n", err)
		return
	}

	defer db.Close()
	dbQueries := database.New(db)

	log.Printf("All connected\n")

	var Results []database.InsertDataRow
	for _, file := range files {

		fileName := file.Name()
		data, err := os.ReadFile(filepath.Join("./folder", fileName))
		if err != nil {
			log.Printf("Err:%v\n", err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		insData, err := dbQueries.InsertData(ctx, database.InsertDataParams{
			Title:    fileName,
			Contents: string(data),
		})
		cancel()

		if err != nil {
			log.Printf("Err:%v\n", err)
			return
		}
		Results = append(Results, insData...)
	}

	log.Println("All work is done.")

	for _, result := range Results {
		fmt.Printf("Title:%v\n", result.Title)
	}
	fmt.Printf("Total:%v\n", len(Results))

}
