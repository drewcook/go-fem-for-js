package main

import "fmt"

// User represents a user of our application
type User2 struct {
	ID int
	FirstName, LastName, Email string
}

// User Group represents a set of Users
type UserGroup struct {
	role string
	users []User2
	newestUser User2
	spaceAvailable bool
}

// This takes in a User and returns a formatted description of it.
func (u User2) describe() string {
	desc := fmt.Sprintf("User Name: %s %s. Email: %s", u.FirstName, u.LastName, u.Email)
	return desc
}

// This takes in a UserGroup and returns formatted description of it.
func (g UserGroup) describe() string {
	// exercise 6a code starts here
	if len(g.users) >= 2 {
		g.spaceAvailable = false // actually needs a pointer to update original param value
	}
	// ends
	desc := fmt.Sprintf("The %s user group has %d user(s). Newest User: %s %s. Accepting New Users: %t", g.role, len(g.users), g.newestUser.FirstName, g.newestUser.LastName, g.spaceAvailable)
	return desc
}

func main() {
	var man User2
	man.ID = 0
	man.FirstName = "John"
	man.LastName = "Kennedy"
	man.Email = "potus@us.gov"
	woman := User2{ID: 1, FirstName: "Marilyn", LastName: "Monroe", Email: "marilyn.monroe@email.com"}
	manDesc := man.describe()
	womanDesc := woman.describe()
	readers := UserGroup{role: "Readers", spaceAvailable: true}
	dog := User2{ID: 2, FirstName: "Rosky", LastName: "Puppers", Email: "ruff@bark.com"}
	readers.users = []User2{man, woman, dog}
	readers.newestUser = dog
	readersDesc := readers.describe()
	admins := UserGroup{role: "Admins", spaceAvailable: false, users: []User2{man}, newestUser: man}
	adminsDesc := admins.describe()
	fmt.Println(manDesc, womanDesc, readersDesc, adminsDesc)
}
