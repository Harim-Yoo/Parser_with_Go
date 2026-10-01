package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"parser/internal/database"
	"time"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type Data struct {
	Class      string `yaml:"class"`
	Title      string `yaml:"title"`
	Slug       string `yaml:"slug"`
	Contents   string `yaml:"contents"`
	ClassOrder int64  `yaml:"class_order"`
}

func ReadData(data Data) Data {
	file, err := os.Open("./data/9-10.yaml")
	if err != nil {
		log.Printf("Err3:%v\n", err)
		return Data{}
	}

	err = yaml.NewDecoder(file).Decode(&data)
	if err != nil {
		log.Printf("Err4:%v\n", err)
		return Data{}
	}
	return data
}

func main() {

	// Load env file.
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("Err:%v\n", err)
		return
	}

	// Get database url.
	dbURL := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Printf("Err2:%v\n", err)
		return
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	queries := database.New(db)

	var data Data
	newData := ReadData(data)

	insErr := queries.InsertBookData(ctx, database.InsertBookDataParams{
		Class:      sql.NullString{String: newData.Class, Valid: true},
		Title:      newData.Title,
		Contents:   newData.Contents,
		Slug:       sql.NullString{String: newData.Slug, Valid: true},
		ClassOrder: sql.NullInt64{Int64: newData.ClassOrder, Valid: true},
	})
	if insErr != nil {
		log.Printf("Err4:%v\n", err)
		return
	}
	log.Printf("Successful seeding!")

}
