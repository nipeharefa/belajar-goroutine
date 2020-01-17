package main

import (
	"encoding/json"
	"fmt"
)

type (
	Balance int
	User    struct {
		ID   int
		Name string
	}
	Response struct {
		User    User
		Balance Balance
	}
)

func getUser(userID int) User {
	return User{
		ID:   1,
		Name: "John Doe",
	}
}

func getBalance(userID int) Balance {
	return 10000
}

func main() {

	response := Response{}

	balance := make(chan Balance)
	user := make(chan User)

	var goGetBalance = func(userID int, ch chan Balance) {
		ch <- getBalance(userID)
	}

	var goGetUser = func(userID int, ch chan User) {
		ch <- getUser(userID)
	}

	go goGetBalance(1, balance)
	go goGetUser(1, user)

	select {
	case b := <-balance:
		response.Balance = b
	case u := <-user:
		response.User = u
	}

	jsonResponse, _ := json.Marshal(response)
	fmt.Println(string(jsonResponse))
}
