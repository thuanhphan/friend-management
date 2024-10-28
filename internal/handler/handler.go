package handler

import (
	"encoding/json"
	"fmt"
	"friend-management-go/internal/model"
	"friend-management-go/internal/service"
	"net/http"
	"net/mail"
)

type FriendshipHandler struct {
	service service.IFriendshipService
}

type GetFriendsRequest struct {
	Email string `json:"email"`
}

type CommonFriendsRequest struct {
	Email1 string `json:"email1"`
	Email2 string `json:"email2"`
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
	// email := chi.URLParam(r, "email")
	fmt.Println("GetFriendsList called") // Debugging statement
	var request GetFriendsRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	email := request.Email

	if email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	// Validate the email format
	if _, err := mail.ParseAddress(email); err != nil {
		http.Error(w, "Invalid email format", http.StatusBadRequest)
		return
	}

	friends, err := h.service.GetFriends(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(friends)
}

func (h *FriendshipHandler) GetCommonFriends(w http.ResponseWriter, r *http.Request) {
	var request CommonFriendsRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	friends, err := h.service.GetCommonFriends(request.Email1, request.Email2)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(friends)
}

func (h *FriendshipHandler) UpdateFriendshipStatus(w http.ResponseWriter, r *http.Request) {
	var friendship model.Friendship
	if err := json.NewDecoder(r.Body).Decode(&friendship); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.UpdateFriendshipStatus(friendship); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
