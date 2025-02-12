package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"

	"AnalyticsAndReporting/util"
)

var testStore Store

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot connect to config file: ", err)
	}

	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err := connPool.Ping(context.Background()); err != nil {
		log.Fatal("cannot ping db:", err)
	}

	testStore = NewStore(connPool)

	os.Exit(m.Run())
}
