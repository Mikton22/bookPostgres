package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
)

func getConnString() string {
	// Пример: читаем из переменной окружения или используем значение по умолчанию
	connString := os.Getenv("DATABASE_URL")
	if connString == "" {
		// Для локального теста без пароля и с отключенным SSL
		connString = "postgres://postgres:1010@localhost:5432/postgres"
	}
	return connString
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	connString := getConnString()

	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	err = conn.Ping(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("успех")
}
