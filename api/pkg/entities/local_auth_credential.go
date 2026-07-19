package entities

import "time"

// LocalAuthCredential stores password credentials for the Nasak-hosted auth fallback.
type LocalAuthCredential struct {
	Email        string    `json:"email" gorm:"primaryKey;type:string;"`
	UserID       UserID    `json:"user_id" gorm:"index;type:string;NOT NULL"`
	PasswordHash string    `json:"-" gorm:"NOT NULL"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
