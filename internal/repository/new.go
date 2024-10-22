package repository

import (
	"friend-management-go/internal/model"
)

type IFriendshipRepository interface {
	MakeFriend(friendship model.Friendship) error
	GetFriends(email string) ([]string, error)
	GetCommonFriends(email1, email2 string) ([]string, error)
	// GetReceivableUpdates(email string) (]model.User, error)
	UpdateFriendshipStatus(friendship model.Friendship) error
	// CreateUser(user model.User) error
	// GetUserByEmail(email string) (model.User, error)
}
