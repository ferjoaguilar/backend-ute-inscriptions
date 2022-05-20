package models

import (
	"time"
)

type Signup struct {
	Name      string    `json:"name"`
	Lastname  string    `json:"lastname"`
	Age       int       `json:"age"`
	Dni       string    `json:"dni"`
	Nit       string    `json:"nit"`
	Country   string    `json:"country"`
	City      string    `json:"city"`
	Address   string    `json:"address"`
	Cellphone string    `json:"cellphone"`
	Graduated string    `json:"graduated"`
	User      string    `json:"user"`
	Document  string    `json:"document"`
	Comments  []string  `json:"comments"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}
