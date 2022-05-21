package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Signup struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `json:"name"`
	Lastname  string             `json:"lastname"`
	Age       int                `json:"age"`
	Dni       string             `json:"dni"`
	Nit       string             `json:"nit"`
	Country   string             `json:"country"`
	City      string             `json:"city"`
	Address   string             `json:"address"`
	Cellphone string             `json:"cellphone"`
	Graduated string             `json:"graduated"`
	User      primitive.ObjectID `json:"user"`
	Completed bool               `json:"completed"`
	CreatedAt time.Time          `json:"created_at"`
}

type SignupLookup struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `json:"name"`
	Lastname  string             `json:"lastname"`
	Age       int                `json:"age"`
	Dni       string             `json:"dni"`
	Nit       string             `json:"nit"`
	Country   string             `json:"country"`
	City      string             `json:"city"`
	Address   string             `json:"address"`
	Cellphone string             `json:"cellphone"`
	Graduated string             `json:"graduated"`
	User      User               `json:"user"`
	Completed bool               `json:"completed"`
	CreatedAt time.Time          `json:"created_at"`
}
