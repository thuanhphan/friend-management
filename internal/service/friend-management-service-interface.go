package service

import "friend-management-go/internal/models"

type IFriendManagementService interface {
	MakeFriend(user1 models.User, user2 models.User) error
	GetFriends(user models.User) ([]models.User, error)
	GetCommonFriends(user1 models.User, user2 models.User) ([]models.User, error)
	GetSubcriberFriends(user models.User) ([]models.User, error)
	UpdateFriend(id int, user models.User) error
}
