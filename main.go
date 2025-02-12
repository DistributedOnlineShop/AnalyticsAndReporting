package main

import (
	"context"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"AnalyticsAndReporting/util"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	config, err := util.LoadConfig("./")
	if err != nil {
		log.Error().Err(err).Msg("app.env is not found")
		os.Exit(1)
	}

	conn, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Error().Err(err).Msg("Failed to connect to the database")
		os.Exit(1)
	}

	defer conn.Close()

	// InitMigration(config.MigrationURL, config.DBSource)
	// store := db.NewStore(conn)
}

func InitMigration(migrationURL, DBSource string) {
	migration, err := migrate.New(migrationURL, DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize migration")
		os.Exit(1)
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Err(err).Msg("Failed to run migration")
		os.Exit(1)
	}
	log.Info().Msg("Migration initialized successfully")
}
