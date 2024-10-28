<<<<<<< HEAD
<<<<<<< HEAD
package controller
=======
package service
>>>>>>> 07f2fdf (apply clean architecture)
=======
package controller
>>>>>>> 856e22d (Refactor structure)

import (
	"friend-management-go/internal/model"
)

type IFriendshipController interface {
	MakeFriend(friendship model.Friendship) error
	GetFriends(email string) ([]string, error)
	GetCommonFriends(email1, email2 string) ([]string, error)
	GetReceivableUpdates(email string) ([]string, error)
	UpdateFriendshipStatus(friendship model.Friendship) error
}
