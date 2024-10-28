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

func (p *PostgresRepository) MakeFriend(friendship model.Friendship) error {
	friend := models.Friendship{
		UserEmail:   null.StringFrom(friendship.FriendEmail),
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
	friends, err := models.Friendships(qm.Where("user_email = ? AND status = ?", email, "friends")).All(context.Background(), p.db)
	if err != nil {
		return nil, err
	}
	var users []model.User
	for _, friend := range friends {
		if friend.FriendEmail.Valid {
			user, err := models.Users(qm.Where("email = ?", friend.FriendEmail)).One(context.Background(), p.db)
			if err != nil {
				return nil, err
			}
			users = append(users, model.User{
				Email: user.Email,
				Name:  user.Name.String,
			})
		}
	}
	return users, nil
}
