package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	Firts_Name    *string            `json:"firts_name" validate:"required, min=2, max=100"`
	Last_Name     *string            `json:"last_name" validate:"required , min=2, max=100"`
	Password      *string            `json:"password" validate:"required, min=6, max=100"`
	Email         *string            `json:"email" validate:"required,email"`
	Avatar        *string            `json:"avatar"`
	Phone         *string            `json:"phone" validate:"required"`
	Token         *string            `json:"token"`
	Refresh_Token *string            `json:"refresh_token"`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
	User_id       string             `json:"user_id"`
}
