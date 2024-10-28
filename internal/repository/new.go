package repository

import (
	"friend-management-go/internal/model"
)

type IFriendshipRepository interface {
<<<<<<< HEAD
<<<<<<< HEAD
	MakeFriend(friendship model.Friendship) error
	GetFriends(email string) ([]string, error)
	GetCommonFriends(email1, email2 string) ([]string, error)
	GetReceivableUpdates(email string) ([]string, error)
	UpdateFriendshipStatus(friendship model.Friendship) error
<<<<<<< HEAD
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
=======
	MakeFriend(friendship model.Friendship) error
	GetFriends(email string) ([]string, error)
	GetCommonFriends(email1, email2 string) ([]string, error)
	GetReceivableUpdates(email string) ([]string, error)
	UpdateFriendshipStatus(friendship model.Friendship) error
<<<<<<< HEAD
	// CreateUser(user model.User) error
	// GetUserByEmail(email string) (model.User, error)
>>>>>>> eb24ee2 (FM-3)
=======
	FriendshipExists(email1, eamil2 string) (bool, error)
>>>>>>> c630ea8 (Customize response, complete FM-8)
=======
>>>>>>> 856e22d (Refactor structure)
}
