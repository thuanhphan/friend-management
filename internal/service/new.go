package service

import (
	"friend-management-go/internal/model"
)

type IFriendshipService interface {
	MakeFriend(friendship model.Friendship) error
	GetFriends(email string) ([]model.User, error)
	// GetCommonFriends(email1, email2 string) ([]model.User, error)
	// // GetReceivableUpdates(email string) (]model.User, error)
	// UpdateFriendshiptstaus(friendship model.Friendship[) error
}
