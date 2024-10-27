package utils

import (
	"context"
	"fmt"
	"friend-management-go/internal/repository/models"
	"regexp"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func IsValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

// ValidateEmailExists checks if an email exists in the users table.
func ValidateEmailExists(db boil.ContextExecutor, email string) error {
	if db == nil {
		return fmt.Errorf("database connection is nil")
	}
	ctx := context.Background()
	exists, err := models.Users(qm.Where("email = ?", email)).Exists(ctx, db)
	if err != nil {
		return fmt.Errorf("failed to check if email exists: %w", err)
	}
	if !exists {
		return fmt.Errorf("email %s does not exist", email)
	}
	return nil
}

// ValidateFriendshipExists checks if a friendship exists between two emails with the status "friends".
func ValidateFriendshipExists(db boil.ContextExecutor, userEmail, friendEmail string) error {
	ctx := context.Background()
	exists, err := models.Friendships(
		qm.Where("(user_email = ? AND friend_email = ? AND status = ?) OR (user_email = ? AND friend_email = ? AND status = ?)",
			userEmail, friendEmail, "friends", friendEmail, userEmail, "friends"),
	).Exists(ctx, db)
	if err != nil {
		return fmt.Errorf("failed to check if friendship exists: %w", err)
	}
	if exists {
		return fmt.Errorf("friendship between %s and %s with status 'friends' already exist", userEmail, friendEmail)
	}
	return nil
}

// ValidateFriendshipStatusExists checks if a friendship exists between two emails with a specific status.
func ValidateFriendshipStatusExists(db boil.ContextExecutor, userEmail, friendEmail, status string) error {
	ctx := context.Background()
	exists, err := models.Friendships(
		qm.Where("(user_email = ? AND friend_email = ? AND status = ?) OR (user_email = ? AND friend_email = ? AND status = ?)",
			userEmail, friendEmail, status, friendEmail, userEmail, status),
	).Exists(ctx, db)
	if err != nil {
		return fmt.Errorf("failed to check if friendship exists with status %s: %w", status, err)
	}
	if exists {
		return fmt.Errorf("friendship between %s and %s with status '%s' already exist", userEmail, friendEmail, status)
	}
	return nil
}
