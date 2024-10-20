package handler

import (
	"encoding/json"
	"friend-management-go/internal/model"
	"friend-management-go/internal/service"
	"net/http"

	"github.com/go-chi/chi"
)

type FriendshipHandler struct {
	service service.IFriendshipService
}

func NewFrienshipHandler(service service.IFriendshipService) *FriendshipHandler {
	return &FriendshipHandler{service: service}
}

func (h *FriendshipHandler) CreateFriendship(w http.ResponseWriter, r *http.Request) {
	var friendship model.Friendship
	if err := json.NewDecoder(r.Body).Decode(&friendship); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.MakeFriend(friendship); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *FriendshipHandler) GetFriendsList(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")
	friends, err := h.service.GetFriends(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(friends)
}
