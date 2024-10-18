package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type PostgresRepository struct {
	databaseUrl string
	db          *sql.DB
}

func NewPostgresRepository(databaseUrl string) *PostgresRepository {
	return &PostgresRepository{databaseUrl: databaseUrl}
}

func (p *PostgresRepository) connect() error {
	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Open a db connection
	db, error := sql.Open("postgres", p.databaseUrl)
	if error != nil {
		log.Fatal(error)
	}

	p.db = db

	// Ping the database to verify the connection
	if err := db.PingContext(ctx); err != nil {
		fmt.Printf("Unable to connect to database: %v\n", err)
		return err
	}

	fmt.Println("Connected to the database successfully!")
	return nil
}

func (p *PostgresRepository) close() error {
	if err := p.db.Close(); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
