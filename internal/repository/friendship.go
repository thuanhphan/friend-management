package repository

import (
	"friend-management-go/internal/model"
	"log"
)

func (p *PostgresRepository) MakeFriend(friendship model.Friendship) error {
	// check connection
	error := p.connect()
	if error != nil {
		log.Fatal(error)
	}
	defer p.close()

	// model := models.Friendship{
	// 	UserEmail:   null.StringFrom(friendship.FriendEmail),
	// 	FriendEmail: null.StringFrom(friendship.FriendEmail),
	// 	Status:      null.StringFrom(friendship.Status),
	// }
	// // error = model.Insert(context.Background(), p.db, boil.Infer())
	// // if error != nil {
	// // 	return error
	// // }

	log.Println("New friend inserted successfully!")
	return nil
}

func (p *PostgresRepository) GetFriends(user model.User) error {
	// check connection
	error := p.connect()
	if error != nil {
		log.Fatal(error)
	}
	defer p.close()

	//Todos

	return nil
}
