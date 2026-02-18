package main

import (
	"context"
	"fmt"
	"time"

	"sql/simple_connection"
	"sql/simple_sql"
	// "github.com/jackc/pgx/v5"
)

// "postgres://postgres:1010@localhost:5432/postgres"

func main() {
	ctx := context.Background()

	conn, err := simple_connection.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}

	if err := simple_sql.CreateTable(ctx, conn); err != nil {
		panic(err)
	}

	//insert book:
	book := simple_sql.BookModel{
		Title:        "тихий дон",
		Author:       "шелохов",
		Review:       "отлично",
		Year:         "1940",
		Read:         false,
		ReadStarted:  time.Now(),
		ReadFinished: time.Time{},
	}
	if err := simple_sql.InsertRow(ctx, conn, book); err != nil {
		panic(err)
	}

	// upd book:
	// if err := simple_sql.UpdateBook(ctx, conn, simple_sql.BookModel{
	// 	ID: 3,
	// 	Title:"1984",
	// 	Author: "оруэлл",
	// 	Year: "1949",
	// 	Review: "отлично",
	// 	Read: true,
	// 	ReadFinished: time.Now(),
	// }); err != nil {
	// 	panic(err)
	// }

	// del book:
	// if err := simple_sql.DeleteRow(ctx, conn, []int{5}); err != nil {
	// 	panic(err)
	// }

	// terminal check
	books, err := simple_sql.SelectRows(ctx, conn)
	if err != nil {
		panic(err)
	}
	fmt.Println("books:", books)

	fmt.Println("succeed!")
}
