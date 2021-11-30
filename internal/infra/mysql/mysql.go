package mysql

import (
	"context"
	"database/sql"
	log "github.com/sirupsen/logrus"
	"os"
)

var (
	db *sql.DB
)

func Connect(ctx context.Context) {
	var err error
	db, err = sql.Open("mysql", os.Getenv("MYSQL_HOST"))
	if err != nil {
		log.Panicln("connecting to mysql: %+v", err)
	}
}


func Check(ctx context.Context) error {
	return db.PingContext(ctx)
}
