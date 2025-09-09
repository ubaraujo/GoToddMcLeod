package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := "password123"
	byteSlice, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		panic(err)
	}
	encryptS := string(byteSlice)
	fmt.Println("Password:", s)
	fmt.Println("Encrypted Password:", encryptS)

	err = bcrypt.CompareHashAndPassword(byteSlice, []byte(s))
	if err != nil {
		fmt.Println("PASSWORD INCORRECT")
	} else {
		fmt.Println("PASSWORD CORRECT")
	}
}
