package models

import "time"

type User struct {
	Firstname string
	Lastname  string
	Nickname  string
	Email     string
	// AvatarURL string
	// AccessToken       string
	// AccessTokenSecret string
	// RefreshToken      string
	// ExpiresAt         time.Time
	PassHash     string
	UserMetaData UserMetaData
}

type UserMetaData struct {
	CreatedAt time.Time
	Role      string
	Gender    Sex
}
