package handler

import (
	"friend-management-go/internal/service"
	"net/http"
)

type friendshipHandler struct {
	service service.IFriendshipService
}

func NewAlbumHandler(ms service.IFriendshipService) *friendshipHandler {
	return &friendshipHandler{service: ms}
}

func (fh friendshipHandler) handleListFriends(w http.ResponseWriter, r *http.Request) {
	// friends, err := fh.service.GetFriends()
	// if err != nil {
	// 	http.Error(w, "Unable to get all movies", http.StatusInternalServerError)
	// 	return
	// }
	// render.RenderList(w, r, friends)
}
