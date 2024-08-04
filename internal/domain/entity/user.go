package entity

import (
	"time"
)

type User struct {
	Id                int    `json:"id"`
	Email             string `json:"email"`
	EncryptedPassword string `json:"encryptedPassword"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	AvatarUrl string `json:"avatarUrl"`
	Name      string `json:"name"`
	About     string `json:"about"`

	/*
	 * Get Subscribers and Subscriptions from virtual table based on user_subscription
	 */
	SubscribersCount   int `json:"subscribersCount"`
	SubscriptionsCount int `json:"subscriptionsCount"`

	Settings UserSettings `json:"userSettings"`
}
