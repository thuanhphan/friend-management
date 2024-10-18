package model

type Friendship struct {
	ID          int    `boil:"id" json:"id"`
	UserEmail   string `boil:"user_email" json:"user_email"`
	FriendEmail string `boil:"friend_email" json:"friend_email"`
	Status      string `boil:"status" json:"status"`
}
