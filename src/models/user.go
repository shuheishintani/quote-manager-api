package models

import "time"

type User struct {
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	Username        string    `json:"username" validate:"required,max=100"`
	ProfileImageUrl string    `json:"profile_image_url" validate:"url"`
	Provider        string    `json:"provider" validate:"required,ma=100"`
	Quotes          []Quote   `json:"quotes"`
	FavoriteQuotes  []Quote   `gorm:"many2many:users_quotes;" json:"favorite_quotes"`
}
