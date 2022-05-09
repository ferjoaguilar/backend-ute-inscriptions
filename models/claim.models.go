package models

import (
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AppClaims struct {
	UserId   primitive.ObjectID `json:"user_id"`
	Email    string             `json:"email"`
	Username string             `json:"username"`
	jwt.StandardClaims
}
