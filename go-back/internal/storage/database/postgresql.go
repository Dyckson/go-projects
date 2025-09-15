package postgres

import (
	"context"
	config "go-back/internal/cmd/server"
	"log"

	"github.com/vingarcia/ksql"
	"github.com/vingarcia/ksql/adapters/kpgx"
)

func GetDB() *ksql.DB {
	dbConnect, err := kpgx.New(context.Background(), config.DbURI, ksql.Config{})
	if err != nil {
		log.Panic(err)
	}
	dbConnect.Exec(context.Background(), "set enable_seqscan = off;")

	return &dbConnect
}
