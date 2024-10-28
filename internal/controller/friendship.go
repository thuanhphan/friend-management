package controller

import (
	"friend-management-go/internal/model"
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 856e22d (Refactor structure)
	"friend-management-go/internal/pkg/utils"
	"friend-management-go/internal/repository"

	"github.com/volatiletech/sqlboiler/v4/boil"
<<<<<<< HEAD
=======
	"friend-management-go/internal/repository"
>>>>>>> 07f2fdf (apply clean architecture)
=======
>>>>>>> 856e22d (Refactor structure)
)

type FriendshipController struct {
	friendshipRepository repository.IFriendshipRepository
<<<<<<< HEAD
<<<<<<< HEAD
	DB                   boil.ContextExecutor
}

func NewFriendshipController(mFriendshipRepoitory repository.IFriendshipRepository, db boil.ContextExecutor) *FriendshipController {
	return &FriendshipController{friendshipRepository: mFriendshipRepoitory, DB: db}
}

func (controller *FriendshipController) MakeFriend(friendship model.Friendship) error {
	if err := utils.ValidateEmailExists(controller.DB, friendship.UserEmail); err != nil {
		return err
	}

	if err := utils.ValidateEmailExists(controller.DB, friendship.FriendEmail); err != nil {
		return err
	}

	if err := utils.ValidateFriendshipExists(controller.DB, friendship.UserEmail, friendship.FriendEmail); err != nil {
		return err
	}

	return controller.friendshipRepository.MakeFriend(friendship)
}

func (controller *FriendshipController) GetFriends(email string) ([]string, error) {
	if err := utils.ValidateEmailExists(controller.DB, email); err != nil {
		return nil, err
	}

	return controller.friendshipRepository.GetFriends(email)
}

func (controller *FriendshipController) GetCommonFriends(email1, email2 string) ([]string, error) {
	if err := utils.ValidateEmailExists(controller.DB, email1); err != nil {
		return nil, err
	}

	if err := utils.ValidateEmailExists(controller.DB, email2); err != nil {
		return nil, err
	}

	return controller.friendshipRepository.GetCommonFriends(email1, email2)
}

func (controller *FriendshipController) UpdateFriendshipStatus(friendship model.Friendship) error {
	if err := utils.ValidateEmailExists(controller.DB, friendship.UserEmail); err != nil {
		return err
	}

	if err := utils.ValidateEmailExists(controller.DB, friendship.FriendEmail); err != nil {
		return err
	}

	if err := utils.ValidateFriendshipStatusExists(controller.DB, friendship.UserEmail, friendship.FriendEmail, friendship.Status); err != nil {
		return err
	}

	return controller.friendshipRepository.UpdateFriendshipStatus(friendship)
}

func (controller *FriendshipController) GetReceivableUpdates(email string) ([]string, error) {
	if err := utils.ValidateEmailExists(controller.DB, email); err != nil {
		return nil, err
	}
	return controller.friendshipRepository.GetReceivableUpdates(email)
=======
=======
	DB                   boil.ContextExecutor
>>>>>>> 856e22d (Refactor structure)
}

func NewFriendshipController(mFriendshipRepoitory repository.IFriendshipRepository, db boil.ContextExecutor) *FriendshipController {
	return &FriendshipController{friendshipRepository: mFriendshipRepoitory, DB: db}
}

func (controller *FriendshipController) MakeFriend(friendship model.Friendship) error {
	if err := utils.ValidateEmailExists(controller.DB, friendship.UserEmail); err != nil {
		return err
	}

	if err := utils.ValidateEmailExists(controller.DB, friendship.FriendEmail); err != nil {
		return err
	}

	if err := utils.ValidateFriendshipExists(controller.DB, friendship.UserEmail, friendship.FriendEmail); err != nil {
		return err
	}

	return controller.friendshipRepository.MakeFriend(friendship)
}

func (controller *FriendshipController) GetFriends(email string) ([]string, error) {
	if err := utils.ValidateEmailExists(controller.DB, email); err != nil {
		return nil, err
	}

	return controller.friendshipRepository.GetFriends(email)
}

func (controller *FriendshipController) GetCommonFriends(email1, email2 string) ([]string, error) {
	if err := utils.ValidateEmailExists(controller.DB, email1); err != nil {
		return nil, err
	}

	if err := utils.ValidateEmailExists(controller.DB, email2); err != nil {
		return nil, err
	}

	return controller.friendshipRepository.GetCommonFriends(email1, email2)
}

func (controller *FriendshipController) UpdateFriendshipStatus(friendship model.Friendship) error {
	if err := utils.ValidateEmailExists(controller.DB, friendship.UserEmail); err != nil {
		return err
	}

	if err := utils.ValidateEmailExists(controller.DB, friendship.FriendEmail); err != nil {
		return err
	}

	if err := utils.ValidateFriendshipStatusExists(controller.DB, friendship.UserEmail, friendship.FriendEmail, friendship.Status); err != nil {
		return err
	}

	return controller.friendshipRepository.UpdateFriendshipStatus(friendship)
}

<<<<<<< HEAD
func (service *FriendshipController) FriendshipExists(email1, email2 string) (bool, error) {
	return service.friendshipRepository.FriendshipExists(email1, email2)
}

func (service *FriendshipController) GetReceivableUpdates(email string) ([]string, error) {
	return service.friendshipRepository.GetReceivableUpdates(email)
>>>>>>> 07f2fdf (apply clean architecture)
=======
func (controller *FriendshipController) GetReceivableUpdates(email string) ([]string, error) {
	if err := utils.ValidateEmailExists(controller.DB, email); err != nil {
		return nil, err
	}
	return controller.friendshipRepository.GetReceivableUpdates(email)
>>>>>>> 856e22d (Refactor structure)
}
