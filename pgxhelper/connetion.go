package pgxhelper

import (
	"context"
	"fmt"
	"os"

	"github.com/bbcoder-gh/api/config"
	"github.com/jackc/pgx/v4"
)

func Connect() (Connection *pgx.Conn, err error) {
	if Connection != nil {
		return
	}
	Connection, err = pgx.Connect(context.Background(), config.DBString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Println(Connection.IsClosed())
	}

	return
}
