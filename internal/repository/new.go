package repository

import (
	"friend-management-go/internal/model"
)

type IFriendshipRepository interface {
	MakeFriend(friendship model.Friendship) error
	GetFriends(email string) ([]model.User, error)
	// GetCommonFriends(email1, email2 string) ([]model.User, error)
	// GetReceivableUpdates(email string) (]model.User, error)
	// UpdateFriendshiptstaus(friendship model.Friendship[) error
	// CreateUser(user model.User) error
	// GetUserByEmail(email string) (model.User, error)
}
