package repository

import (
	"context"
	"database/sql"
	"friend-management-go/internal/model"
	"friend-management-go/internal/repository/models"
	"log"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

const SUBCRIBED_FRIENDSHIP_STATUS string = "subcribed"

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

func (p *PostgresRepository) GetFriends(email string) ([]string, error) {
	var friends []string

	// Custom SQL query to join users and friendships tables
	query := `Select email from friendships 
	    Join users on friendships.user_email = users.email
		Where friend_email = $1 and status = 'friend'
		Union
		Select friend_email
		From friendships
		Where user_email = $1 and status = 'friend'`

	// Execute the query and bind the results to the friends slice
	ctx := context.Background()
	rows, err := p.db.QueryContext(ctx, query, email)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var friendEmail string
		if err := rows.Scan(&friendEmail); err != nil {
			return nil, err
		}
		friends = append(friends, friendEmail)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return friends, nil
}

func (p *PostgresRepository) GetCommonFriends(email1, email2 string) ([]string, error) {
	var friends []string

	query := `Select email from 
	    (Select user_email from friendships where friend_email = $1
		Union 
		Select friend_email from friendships where user_email = $1) as list1
		Inner Join
		(Select user_email from friendships where friend_email = $2
		Union 
		Select friend_email from friendships where user_email = $2) as list2
		on list1.user_email = list2.user_email
		Inner Join users on users.email = list1.user_email`

	ctx := context.Background()
	rows, err := p.db.QueryContext(ctx, query, email1, email2)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var friendEmail string
		if err := rows.Scan(&friendEmail); err != nil {
			return nil, err
		}
		friends = append(friends, friendEmail)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return friends, nil
}

func (p *PostgresRepository) UpdateFriendshipStatus(friendship model.Friendship) error {
	friend := models.Friendship{
		UserEmail:   null.StringFrom(friendship.UserEmail),
		FriendEmail: null.StringFrom(friendship.FriendEmail),
		Status:      null.StringFrom(friendship.Status),
	}

	ctx := context.Background()
	// Check if the friendship already exists
	existingFriendship, err := models.Friendships(
		qm.Where("user_email = ? AND friend_email = ?", friend.UserEmail, friend.FriendEmail),
	).One(ctx, p.db)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if existingFriendship != nil {
		existingFriendship.Status = null.StringFrom(friend.Status.String)
		_, err = existingFriendship.Update(ctx, p.db, boil.Infer())
		return err
	}

	newFriendship := &models.Friendship{
		UserEmail:   null.StringFrom(friendship.UserEmail),
		FriendEmail: null.StringFrom(friendship.FriendEmail),
		Status:      null.StringFrom(friendship.Status),
	}
	return newFriendship.Insert(ctx, p.db, boil.Infer())
}
