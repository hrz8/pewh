package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/aws/aws-lambda-go/lambda"
)

const (
	DatabaseURL = "postgres://postgres:toor@postgres-dev:5432/devdb?sslmode=disable"
)

type Body struct {
	Message string `json:"message"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func main() {
	lambda.Start(handler)
}

func handler(_ context.Context, payload Body) (Response, error) {
	log.Println("lambda invoked", payload)

	conn, err := sql.Open("pgx", DatabaseURL)
	if err != nil {
		log.Println("cannot connect to database")
		return Response{
			Status:  500,
			Message: "Not Ok!",
		}, nil
	}

	rows, err := conn.Query("SELECT * FROM test.names")
	if err != nil {
		log.Println("cannot perform sql query")
		return Response{
			Status:  500,
			Message: "Not Ok!",
		}, nil
	}

	defer rows.Close()

	for rows.Next() {
		var id int64
		var name string

		err := rows.Scan(&id, &name)
		if err != nil {
			log.Println("Error scanning row:", err)
			continue
		}

		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}

	return Response{
		Status:  200,
		Message: "Ok!",
	}, nil
}
