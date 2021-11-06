package main

import "fmt"

// User is a user type
type User struct {
	ID int
	FirstName, LastName, Email string
}

func describeUser(u User) string {
	desc := fmt.Sprintf("Name: %s %s, Email: %s", u.FirstName, u.LastName, u.Email)
	return desc
}

func (u *User) describe() string {
	desc := fmt.Sprintf("Name: %s %s, Email: %s", u.FirstName, u.LastName, u.Email)
	return desc
}

func main() {
	user := User{ID: 1, FirstName: "Marilyn", LastName: "Monroe", Email: "marilyn.monroe@gmail.com"}
	user2 := User{ID: 2, FirstName: "Johnny", LastName: "Kennedy", Email: "potus@gmail.com"}
	// as function
	desc := describeUser(user)
	fmt.Println(desc)
	// as method
	desc = user.describe()
	fmt.Println(desc)
	desc = user2.describe()
	fmt.Println(desc)
}
