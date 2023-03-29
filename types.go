package main

import (
	"math/rand"
	"time"
)

type Account struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Number    int64     `json:"number"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateAccountRequestDto struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type TransferAccountRequestDto struct {
	ToAccount int `json:"to_account"`
	Amount    int `json:"amount"`
}

func NewAccount(FirstName, LastName string) *Account {
	return &Account{
		FirstName: FirstName,
		LastName:  LastName,
		Number:    int64(rand.Intn(100000)),
		CreatedAt: time.Now(),
	}
}
