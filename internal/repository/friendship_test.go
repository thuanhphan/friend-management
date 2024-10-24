package repository

import (
	"friend-management-go/internal/model"
	"friend-management-go/internal/repository/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRepo_MakeFriend(t *testing.T) {
	mockRepo := new(mocks.IFriendshipRepository)

	friendship := model.Friendship{
		UserEmail:   "user@example.com",
		FriendEmail: "friend@example.com",
		Status:      "friendship",
	}

	// Set up the mock to expect the Insert call and return nil (no error)
	mockRepo.On("MakeFriend", mock.Anything).Return(nil)

	// Call the method
	err := mockRepo.MakeFriend(friendship)

	// Assert that the method was called with the expected arguments and no error was returned
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestRepo_GetFriends(t *testing.T) {
	// Create an instance of the mock repository
	mockRepo := new(mocks.IFriendshipRepository)

	// Define the expected output
	expectedFriends := []string{"friend1@example.com", "friend2@example.com"}

	// Setup the mock to return the expected output
	mockRepo.On("GetFriends", mock.Anything).Return(expectedFriends, nil)

	// Call the method
	friends, err := mockRepo.GetFriends("user@example.com")

	// Assert the results
	assert.NoError(t, err)
	assert.Equal(t, expectedFriends, friends)

	// Ensure that the expectations were met
	mockRepo.AssertExpectations(t)
}

func TestRepo_GetCommonFriends(t *testing.T) {
	mockRepo := new(mocks.IFriendshipRepository)

	expectedFriends := []string{"friend1@example.com", "friend2@example.com"}

	mockRepo.On("GetCommonFriends", mock.Anything, mock.Anything).Return(expectedFriends, nil)

	friends, err := mockRepo.GetCommonFriends("user@example.com", "user1@example.com")

	assert.NoError(t, err)
	assert.Equal(t, expectedFriends, friends)

	mockRepo.AssertExpectations(t)
}

func TestRepo_UpdateFriendshipStatus(t *testing.T) {
	mockRepo := new(mocks.IFriendshipRepository)

	friendship := model.Friendship{
		UserEmail:   "user@example.com",
		FriendEmail: "friend@example.com",
		Status:      "friendship",
	}

	mockRepo.On("UpdateFriendshipStatus", mock.Anything).Return(nil)

	err := mockRepo.UpdateFriendshipStatus(friendship)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestRepo_GetReceivableUpdates(t *testing.T) {
	mockRepo := new(mocks.IFriendshipRepository)

	expectedFriends := []string{"friend1@example.com", "friend2@example.com"}

	mockRepo.On("GetReceivableUpdates", mock.Anything).Return(expectedFriends, nil)

	friends, err := mockRepo.GetReceivableUpdates("user@example.com")

	assert.NoError(t, err)
	assert.Equal(t, expectedFriends, friends)

	mockRepo.AssertExpectations(t)
}
