package handler

import (
<<<<<<< HEAD
<<<<<<< HEAD
	"encoding/json"
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

type FriendListResponse struct {
	Success    bool     `json:"success"`
	Friends    []string `json:"friends,omitempty"`
	Recipients []string `json:"recipients,omitempty"`
	Count      int      `json:"count,omitempty"`
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

	exists, err := h.service.FriendshipExists(friendship.UserEmail, friendship.FriendEmail)
	if err != nil {
		http.Error(w, "Error checking friendship existence", http.StatusInternalServerError)
		return
	}
	if exists {
		response := FriendListResponse{Success: false}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	if err := h.service.MakeFriend(friendship); err != nil {
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

	friends, err := h.service.GetCommonFriends(request.Email1, request.Email2)
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

	if err := h.service.UpdateFriendshipStatus(friendship); err != nil {
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

	friends, err := h.service.GetReceivableUpdates(request.Email)
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
=======
	"encoding/json"
	"friend-management-go/internal/model"
>>>>>>> eb24ee2 (FM-3)
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

<<<<<<< HEAD
func (fh friendshipHandler) handleListFriends(w http.ResponseWriter, r *http.Request) {
	// friends, err := fh.service.GetFriends()
	// if err != nil {
	// 	http.Error(w, "Unable to get all movies", http.StatusInternalServerError)
	// 	return
	// }
	// render.RenderList(w, r, friends)
>>>>>>> d88719a (Arrange the layered architecture)
=======
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
>>>>>>> eb24ee2 (FM-3)
}
