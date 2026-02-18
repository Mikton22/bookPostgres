package simple_sql

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

type BookModel struct {
	ID           int
	Title        string
	Author       string
	Review       string
	Year         string
	Read         bool
	ReadStarted  time.Time
	ReadFinished time.Time
}

func CreateTable(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
	CREATE TABLE IF NOT EXISTS books (
	id SERIAL PRIMARY KEY,
	title VARCHAR(200) NOT NULL,
	author VARCHAR(1000) NOT NULL,
	review VARCHAR(1000),
	year VARCHAR(1000) NOT NULL,
	read BOOLEAN NOT NULL,
	read_started TIMESTAMP NOT NULL,
	read_finished TIMESTAMP,

	UNIQUE(title)
	);
	`

	_, err := conn.Exec(ctx, sqlQuery)
	return err
}

func DeleteRow(ctx context.Context, conn *pgx.Conn, booksIDs []int) error {
	sqlQuery := `
	DELETE FROM books
	WHERE id = ANY($1);
	`
	_, err := conn.Exec(ctx, sqlQuery, booksIDs)

	return err
}

func InsertRow(
	ctx context.Context,
	conn *pgx.Conn,
	book BookModel,
) error {
	sqlQuery := `
	INSERT INTO books(title, author, review, year, read, read_started, read_finished)
	VALUES ($1, $2, $3, $4, $5, $6, $7);
	`

	_, err := conn.Exec(
		ctx,
		sqlQuery,
		book.Title,
		book.Author,
		book.Review,
		book.Year,
		book.Read,
		book.ReadStarted,
		book.ReadFinished,
	)

	return err
}

func SelectRows(ctx context.Context, conn *pgx.Conn) ([]BookModel, error) {
	sqlQuery := `
	SELECT id, title, author, review, year, read, read_started, read_finished
	FROM books
	ORDER BY id ASC
	`

	rows, err := conn.Query(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := make([]BookModel, 0)

	for rows.Next() {
		var book BookModel

		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.Review,
			&book.Year,
			&book.Read,
			&book.ReadStarted,
			&book.ReadFinished,
		)
		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}

func printBooks(book BookModel) {
	fmt.Println("--------------------------------------")
	fmt.Println("id", book.ID)
	fmt.Println("title", book.Title)
	fmt.Println("author", book.Author)
	fmt.Println("review", book.Review)
	fmt.Println("year", book.Year)
	fmt.Println("read", book.Read)
	fmt.Println("read_started", book.ReadStarted)
	fmt.Println("read_finished", book.ReadFinished)
}

func UpdateBook(
	ctx context.Context,
	conn *pgx.Conn,
	book BookModel,
) error {
	sqlQuery := `
	UPDATE books
	SET title=$1, author=$2, review=$3, year=$4, read=$5, read_finished=$6
	WHERE id=$7;
	`

	_, err := conn.Exec(
		ctx,
		sqlQuery,
		book.Title,
		book.Author,
		book.Review,
		book.Year,
		book.Read,
		// book.ReadStarted,
		book.ReadFinished,
		book.ID,
	)

	return err
}
