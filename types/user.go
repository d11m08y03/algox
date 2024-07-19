package types

import "time"

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(user RegisterUserPayload) error
  GetUserByName(firstName string, lastName string) (*User, error)
}

type User struct {
	ID         int       `json:"id"`
	FirstName  string    `json:"firstName"`
	LastName   string    `json:"lastName"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	IsHospital bool      `json:"isHospital"`
	IsNGO      bool      `json:"isNGO"`
	BloodType  string    `json:"bloodType"`
	CreatedAt  time.Time `json:"createdAt"`
	Points     int       `json:"point"`
}

type RegisterUserPayload struct {
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	IsHospital bool   `json:"isHospital"`
	IsNGO      bool   `json:"isNGO"`
	BloodType  string `json:"bloodType"`
}

type LoginUserPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type FindUserPayload struct {
	FirstName  string    `json:"firstName"`
	LastName   string    `json:"lastName"`
}
