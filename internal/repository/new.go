package repository

import (
	"friend-management-go/internal/model"
)

type IFriendshipRepository interface {
<<<<<<< HEAD
	MakeFriend(friendship model.Friendship) error
	GetFriends(email string) ([]string, error)
	GetCommonFriends(email1, email2 string) ([]string, error)
	GetReceivableUpdates(email string) ([]string, error)
	UpdateFriendshipStatus(friendship model.Friendship) error
=======
	MakeFriend(friendship *model.Friendship) error
	GetFriends(email string) ([]model.User, error)
	GetCommonFriends(email1 string, email2 string) ([]*model.User, error)
	GetReceivableUpdates(email string) ([]*model.User, error)
	Subcribe(email1 string, email2 string) error
	Block(email1 string, email2 string) error
	CreateUser(user model.User) error
	GetUserByEmail(email string) (model.User, error)
>>>>>>> d88719a (Arrange the layered architecture)
}
