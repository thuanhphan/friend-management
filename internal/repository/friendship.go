package repository

import (
	"context"
	"database/sql"
	"friend-management-go/internal/model"
	"friend-management-go/internal/repository/models"
	"log"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (p *PostgresRepository) MakeFriend(friendship model.Friendship) error {
	friend := models.Friendship{
		UserEmail:   null.StringFrom(friendship.UserEmail),
		FriendEmail: null.StringFrom(friendship.FriendEmail),
		Status:      null.StringFrom(friendship.Status),
	}
	error := friend.Insert(context.Background(), p.db, boil.Infer())
	if error != nil {
		return error
	}

	log.Println("New friend inserted successfully!")
	return nil
}

func (p *PostgresRepository) GetFriends(email string) ([]model.User, error) {
	var friends []model.User

	// Custom SQL query to join users and friendships tables
	query := `Select email from friendships join users on friendships.user_email = users.email
			Where friend_email = $1 and status = 'Friend'
			Union
			Select friend_email
			From friendships
			Where user_email = $1 and status = 'Friend'`

	// Execute the query and bind the results to the friends slice
	ctx := context.Background()
	_, err := p.db.ExecContext(ctx, query, email)

	if err != nil {
		return nil, err
	}
	return friends, nil
}
