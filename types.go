package main

import "math/rand"

type Account struct {
	ID        int
	FirstName string
	LastName  string
	Number    int64
	Balance   int64
}

func NewAccount(firstName string, lastName string) *Account {
	return &Account{
		ID:        rand.Intn(1000000),
		FirstName: firstName,
		LastName:  lastName,
		Number:    rand.Int63n(1000000),
	}
}
