package main

import (
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Number    int64     `json:"number"`
	Password  string    `json:"password"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginResponseDto struct {
	Number int64  `json:"number"`
	Token  string `json:"token"`
}

type LoginRequestDto struct {
	Number   int64  `json:"number"`
	Password string `json:"password"`
}

type CreateAccountRequestDto struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

type TransferAccountRequestDto struct {
	ToAccount int `json:"to_account"`
	Amount    int `json:"amount"`
}

func NewAccount(FirstName, LastName, Password string) (*Account, error) {
	bcrypt, err := bcrypt.GenerateFromPassword([]byte(Password), 10)
	if err != nil {
		return nil, err
	}
	return &Account{
		FirstName: FirstName,
		LastName:  LastName,
		Password:  string(bcrypt),
		Number:    int64(rand.Intn(100000)),
		CreatedAt: time.Now(),
	}, nil
}
