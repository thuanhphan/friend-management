package repository

import (
	"friend-management-go/internal/model"
)

type IFriendshipRepository interface {
	MakeFriend(friendship model.Friendship) error
	GetFriends(email string) ([]string, error)
	GetCommonFriends(email1, email2 string) ([]string, error)
	GetReceivableUpdates(email string) ([]string, error)
	UpdateFriendshipStatus(friendship model.Friendship) error
}
