package main

import (
	"database/sql"
	"fmt"
	"friend-management-go/internal/controller"
	"friend-management-go/internal/handler"
	"friend-management-go/internal/repository"
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

	// Initialize repository, service, and handler
	friendshipRepo := repository.NewPostgresRepository(db)
	friendshipController := controller.NewFriendshipController(friendshipRepo)
	friendshipHandler := handler.NewFrienshipHandler(friendshipController)

	// Set up the router
	r := chi.NewRouter()
	r.Post("/make-friends", friendshipHandler.CreateFriendship)
	r.Post("/friends", friendshipHandler.GetFriendsList)
	r.Post("/common-friends", friendshipHandler.GetCommonFriends)
	r.Post("/receivable-updates", friendshipHandler.GetReceivableUpdates)
	r.Post("/subscribe", friendshipHandler.UpdateFriendshipStatus)
	r.Post("/block", friendshipHandler.UpdateFriendshipStatus)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", r))
}
