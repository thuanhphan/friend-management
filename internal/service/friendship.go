package service

import (
	"friend-management-go/internal/model"
	"friend-management-go/internal/repository"
)

type FriendshipService struct {
	friendshipRepository repository.IFriendshipRepository
}

func NewFriendshipService(mFriendshipRepoitory repository.IFriendshipRepository) *FriendshipService {
	return &FriendshipService{friendshipRepository: mFriendshipRepoitory}
}

func (service *FriendshipService) MakeFriend(friendship model.Friendship) error {
	return service.friendshipRepository.MakeFriend(friendship)
}

func (service *FriendshipService) GetFriends(email string) ([]model.User, error) {
	return service.friendshipRepository.GetFriends(email)
}
