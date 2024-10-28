package handler

import (
<<<<<<< HEAD
	"encoding/json"
	"friend-management-go/internal/controller"
	"friend-management-go/internal/model"
	"friend-management-go/internal/pkg/utils"
	"net/http"
)

type FriendshipHandler struct {
	controller controller.IFriendshipController
}

type GetFriendsRequest struct {
	Email string `json:"email"`
}

type CommonFriendsRequest struct {
	Email1 string `json:"email1"`
	Email2 string `json:"email2"`
}

type FriendListResponse struct {
	Success    bool     `json:"success"`
	Friends    []string `json:"friends,omitempty"`
	Recipients []string `json:"recipients,omitempty"`
	Count      int      `json:"count,omitempty"`
}

func NewFrienshipHandler(controller controller.IFriendshipController) *FriendshipHandler {
	return &FriendshipHandler{controller: controller}
}

func (h *FriendshipHandler) CreateFriendship(w http.ResponseWriter, r *http.Request) {
	var friendship model.Friendship
	if err := json.NewDecoder(r.Body).Decode(&friendship); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate emails
	if !utils.IsValidEmail(friendship.UserEmail) || !utils.IsValidEmail(friendship.FriendEmail) {
		http.Error(w, "Invalid email format", http.StatusBadRequest)
		return
	}

	if err := h.controller.MakeFriend(friendship); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := FriendListResponse{
		Success: true,
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *FriendshipHandler) GetFriendsList(w http.ResponseWriter, r *http.Request) {
	var request GetFriendsRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	email := request.Email

	// Validate email
	if !utils.IsValidEmail(request.Email) {
		http.Error(w, "Invalid email format", http.StatusBadRequest)
		return
	}

	friends, err := h.controller.GetFriends(email)
	if err != nil {
		http.Error(w, "Error retrieving friends", http.StatusInternalServerError)
		return
	}

	response := FriendListResponse{
		Success: true,
		Friends: friends,
		Count:   len(friends),
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *FriendshipHandler) GetCommonFriends(w http.ResponseWriter, r *http.Request) {
	var request CommonFriendsRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate emails
	if !utils.IsValidEmail(request.Email1) || !utils.IsValidEmail(request.Email2) {
		http.Error(w, "Invalid email format", http.StatusBadRequest)
		return
	}

	friends, err := h.controller.GetCommonFriends(request.Email1, request.Email2)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := FriendListResponse{
		Success: true,
		Friends: friends,
		Count:   len(friends),
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *FriendshipHandler) UpdateFriendshipStatus(w http.ResponseWriter, r *http.Request) {
	var friendship model.Friendship
	if err := json.NewDecoder(r.Body).Decode(&friendship); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate emails
	if !utils.IsValidEmail(friendship.UserEmail) || !utils.IsValidEmail(friendship.FriendEmail) {
		http.Error(w, "Invalid email format", http.StatusBadRequest)
		return
	}

	if err := h.controller.UpdateFriendshipStatus(friendship); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := FriendListResponse{
		Success: true,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *FriendshipHandler) GetReceivableUpdates(w http.ResponseWriter, r *http.Request) {
	var request GetFriendsRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate email
	if !utils.IsValidEmail(request.Email) {
		http.Error(w, "Invalid email format", http.StatusBadRequest)
		return
	}

	friends, err := h.controller.GetReceivableUpdates(request.Email)
	if err != nil {
		http.Error(w, "Error retrieving friends", http.StatusInternalServerError)
		return
	}

	response := FriendListResponse{
		Success:    true,
		Recipients: friends,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
=======
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
>>>>>>> d88719a (Arrange the layered architecture)
}
