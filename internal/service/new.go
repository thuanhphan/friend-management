package service

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

type IFriendshipService interface {
	MakeFriend(Friendship model.Friendship) error
	GetFriends(email string) ([]model.User, error)
	GetCommonFriends(email1 string, email2 string) ([]model.User, error)
	GetReceivableUpdates(email string) ([]model.User, error)
	Subcribe(email1 string, email2 string) error
	Block(email1 string, email2 string) error
>>>>>>> d88719a (Arrange the layered architecture)
}
