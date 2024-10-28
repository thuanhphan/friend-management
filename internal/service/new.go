package service

<<<<<<< HEAD
<<<<<<< HEAD
import (
	"friend-management-go/internal/model"
)

type IFriendshipService interface {
	MakeFriend(friendship model.Friendship) error
	GetFriends(email string) ([]string, error)
	GetCommonFriends(email1, email2 string) ([]string, error)
	GetReceivableUpdates(email string) ([]string, error)
	UpdateFriendshipStatus(friendship model.Friendship) error
	FriendshipExists(email1, eamil2 string) (bool, error)
=======
import "friend-management-go/internal/model"
=======
import (
	"friend-management-go/internal/model"
)
>>>>>>> eb24ee2 (FM-3)

type IFriendshipService interface {
	MakeFriend(friendship model.Friendship) error
	GetFriends(email string) ([]model.User, error)
<<<<<<< HEAD
	GetCommonFriends(email1 string, email2 string) ([]model.User, error)
	GetReceivableUpdates(email string) ([]model.User, error)
	Subcribe(email1 string, email2 string) error
	Block(email1 string, email2 string) error
>>>>>>> d88719a (Arrange the layered architecture)
=======
	// GetCommonFriends(email1, email2 string) ([]model.User, error)
	// // GetReceivableUpdates(email string) (]model.User, error)
	// UpdateFriendshiptstaus(friendship model.Friendship[) error
>>>>>>> eb24ee2 (FM-3)
}
