package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)
func main() {
	password := []byte("123")
	hashPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}
	fmt.Println(string(hashPassword))

	err = bcrypt.CompareHashAndPassword(hashPassword, password)
	fmt.Println(err)
}