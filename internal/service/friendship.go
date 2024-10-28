package service

import (
	"friend-management-go/internal/model"
	"friend-management-go/internal/repository"
)

type DefaultFriendshipService struct {
	friendshipRepository repository.IFriendshipRepository
}

func NewDefaultFriendshipService(mFriendshipRepoitory repository.IFriendshipRepository) *DefaultFriendshipService {
	return &DefaultFriendshipService{friendshipRepository: mFriendshipRepoitory}
}

func (service *DefaultFriendshipService) GetFriends(email string) ([]model.User, error) {
	return service.friendshipRepository.GetFriends(email)
}
