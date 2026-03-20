package simple_connection

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

func CreateConnection(ctx context.Context) (*pgx.Conn, error) {
	CONN_STRING := os.Getenv("CONN_STRING")
	return pgx.Connect(ctx, CONN_STRING)
}