package utils

import (
	"errors"

	"github.com/ferjoaguilar/backend-utec-inscriptions/models"
	"github.com/ferjoaguilar/backend-utec-inscriptions/server"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTokenInformation(s server.Server, tokenString string) (primitive.ObjectID, error) {

	token, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.Config().JWTSecret), nil
	})

	if err != nil {
		return primitive.NilObjectID, err
	}

	claims, ok := token.Claims.(*models.AppClaims)
	if !ok {
		return primitive.NilObjectID, errors.New("Somithing failed to getting information")
	}
	return claims.UserId, nil

}
