package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"parser/internal/database"
	"time"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Printf("Err:%v\n", err)
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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = dbQueries.UpdateClass(ctx, database.UpdateClassParams{
		Class: sql.NullString{
			String: "사역반(섬김반)",
			Valid:  true,
		},
		Column2: []string{
			"2-5 하나님 나라 확장을 위해 열심히 사십니까?(영적 전투론)",
			"전도를 열심히 하십니까?",
			"자기 은사를 아십니까?",
			"2-6 그리스도인의 직업관: 하나님 나라 확장(그리스도의 은혜의 복음 전파)을 우선순위로 살아가고 있습니까?",
			"다른 영혼을 돌보고 있습니까?",
			"신구약 성경을 몇 번 통독했습니까?",
			"영 분별론에 대해 얼마나 아십니까?",
			"귀신의 공격을 잘 물리치고 있습니까?",
			"영과 혼과 몸이 건강하십니까?",
			"종말론에 대해 얼마나 아십니까?",
			"교회사에 대해 얼마나 아십니까? 중생론의 역사",
			"칭의론의 역사",
			"교회론에 대해 아십니까?",
			"감사교회의 역사적 사명",
			"복음주의 대각성 운동 연합회 취지를 얼마나 이해하고, 위하여 기도하십니까?",
		},
	})
	if err != nil {
		log.Printf("Err:%v\n", err)
		return
	}

	log.Printf("Everything in the class")

}
