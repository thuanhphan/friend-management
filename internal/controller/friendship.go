package controller

import (
	"friend-management-go/internal/model"
	"friend-management-go/internal/repository"
)

type FriendshipController struct {
	friendshipRepository repository.IFriendshipRepository
}

func NewFriendshipController(mFriendshipRepoitory repository.IFriendshipRepository) *FriendshipController {
	return &FriendshipController{friendshipRepository: mFriendshipRepoitory}
}

func (service *FriendshipController) MakeFriend(friendship model.Friendship) error {
	return service.friendshipRepository.MakeFriend(friendship)
}

func (service *FriendshipController) GetFriends(email string) ([]string, error) {
	return service.friendshipRepository.GetFriends(email)
}

func (service *FriendshipController) GetCommonFriends(email1, email2 string) ([]string, error) {
	return service.friendshipRepository.GetCommonFriends(email1, email2)
}

func (service *FriendshipController) UpdateFriendshipStatus(friendship model.Friendship) error {
	return service.friendshipRepository.UpdateFriendshipStatus(friendship)
}

func (service *FriendshipController) FriendshipExists(email1, email2 string) (bool, error) {
	return service.friendshipRepository.FriendshipExists(email1, email2)
}

func (service *FriendshipController) GetReceivableUpdates(email string) ([]string, error) {
	return service.friendshipRepository.GetReceivableUpdates(email)
}
