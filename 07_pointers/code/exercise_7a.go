package main

import "fmt"

type User struct {
	ID int
	FirstName, LastName, Email string
}

func updateEmail(u *User, newEmail string) {
	u.Email = newEmail
}

func main() {
	fmt.Println("Pointers!")

	// update an email
	user := User{
		ID: 0,
		FirstName: "Test",
		LastName: "User",
		Email: "before@test.com",
	}
	fmt.Println(user.Email)
	desiredEmail := "after@test.com"
	updateEmail(&user, desiredEmail)
	fmt.Println(user.Email)
}
