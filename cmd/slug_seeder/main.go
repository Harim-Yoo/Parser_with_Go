package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"parser/internal/database"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("Err1:%v\n", err)
		return
	}

	dbURL := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Printf("Err2:%v\n", err)
		return
	}
	defer db.Close()

	dbQueries := database.New(db)

	ctx := context.Background()

	err = dbQueries.UpdateSlug(ctx, database.UpdateSlugParams{
		Slug: sql.NullString{
			String: "복음주의 대각성 운동 연합회 취지를 얼마나 이해하고, 위하여 기도하십니까?",
			Valid:  true,
		},
		Column2: sql.NullString{
			String: "49반",
			Valid:  true,
		},
	})
	if err != nil {
		log.Printf("Err3:%v\n", err)
		return
	}
	log.Printf("Updated successfully.")

}
