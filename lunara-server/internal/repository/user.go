package repository

import (
	"context"
	"time"

	"github.com/nishaj0/lunara-app/lunara-server/internal/db"
	"github.com/nishaj0/lunara-app/lunara-server/internal/model"
)

func CreateUser(ctx context.Context, user *model.User) error {
	query := `
			INSERT INTO users (username, email, password_hash, full_name, bio, profile_image_url, is_verified, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
			RETURNING id, created_at, updated_at
	`
	now := time.Now()
	return db.GetDB().QueryRow(
		ctx, query,
		user.Username, user.Email, user.PasswordHash, user.FullName, user.Bio, user.ProfileImageURL, user.IsVerified, now, now,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
}
