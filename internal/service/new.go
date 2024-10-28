package service

import "friend-management-go/internal/model"

type IFriendshipService interface {
	MakeFriend(Friendship model.Friendship) error
	GetFriends(email string) ([]model.User, error)
	GetCommonFriends(email1 string, email2 string) ([]model.User, error)
	GetReceivableUpdates(email string) ([]model.User, error)
	Subcribe(email1 string, email2 string) error
	Block(email1 string, email2 string) error
}
