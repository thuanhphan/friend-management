package main

import (
	"database/sql"
	"fmt"
	"friend-management-go/internal/handler"
	"friend-management-go/internal/repository"
	"friend-management-go/internal/service"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbname   = "postgres"
)

func main() {
	fmt.Println("Server is running 8080")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open the database connection
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	// Verify the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database!")

	_, err = db.Exec(`INSERT INTO friendships (user_email, friend_email, status) VALUES($1, $2, $3)`, "alice@example.com", "example@example.com", "friend")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Insert successful")

	// Initialize repository, service, and handler
	friendshipRepo := repository.NewPostgresRepository(db)
	friendshipService := service.NewFriendshipService(friendshipRepo)
	friendshipHandler := handler.NewFrienshipHandler(friendshipService)

	// Set up the router
	r := chi.NewRouter()
	r.Post("/make-friends", friendshipHandler.CreateFriendship)
	// r.Get("/friends/{email}", friendHandler.GetFriends)
	// r.Post("/common-friends", friendHandler.GetCommonFriends)
	// r.Post("/subscribe", friendHandler.Subscribe)
	// r.Post("/block", friendHandler.Block)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", r))
}
