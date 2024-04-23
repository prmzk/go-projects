package main

import "github.com/google/uuid"

type CreateAccountRequest struct {
	Fullname string `json:"fullname"`
}

type Account struct {
	ID        string `json:"id"`
	Balance   int    `json:"balance"`
	Fullname  string `json:"fullname"`
	AccNumber string `json:"acc_number"`
}

func NewAccount(fullname string) *Account {
	return &Account{
		ID:        uuid.New().String(),
		Fullname:  fullname,
		AccNumber: uuid.New().String(),
	}
}
