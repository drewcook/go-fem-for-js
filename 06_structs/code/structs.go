package main

import "fmt"

// User represents a user of our application
type User struct {
	ID int
	FirstName, LastName, Email string
}

// User Group represents a set of Users
type UserGroup struct {
	role string
	users []User
	newestUser User
	spaceAvailable bool
}

// This takes in a User and returns a formatted description of it.
func descUser(u User) string {
	desc := fmt.Sprintf("User Name: %s %s. Email: %s", u.FirstName, u.LastName, u.Email)
	return desc
}

// This takes in a UserGroup and returns formatted description of it.
func descGroup(g UserGroup) string {
	// exercise 6a code starts here
	if len(g.users) >= 2 {
		g.spaceAvailable = false // actually needs a pointer to update original param value
	}
	// ends
	desc := fmt.Sprintf("The %s user group has %d user(s). Newest User: %s %s. Accepting New Users: %t", g.role, len(g.users), g.newestUser.FirstName, g.newestUser.LastName, g.spaceAvailable)
	return desc
}

func main() {
	var man User
	man.ID = 0
	man.FirstName = "John"
	man.LastName = "Kennedy"
	man.Email = "potus@us.gov"

	woman := User{ID: 1, FirstName: "Marilyn", LastName: "Monroe", Email: "marilyn.monroe@email.com"}

	manDesc := descUser(man)
	womanDesc := descUser(woman)
	fmt.Println(manDesc)
	fmt.Println(womanDesc)

	readers := UserGroup{role: "Readers", spaceAvailable: true}
	dog := User{ID: 2, FirstName: "Rosky", LastName: "Puppers", Email: "ruff@bark.com"}
	readers.users = []User{man, woman, dog}
	readers.newestUser = dog
	readersDesc := descGroup(readers)


	admins := UserGroup{role: "Admins", spaceAvailable: false, users: []User{man}, newestUser: man}
	adminsDesc := descGroup(admins)

	fmt.Println(readersDesc)
	fmt.Println(adminsDesc)
}
