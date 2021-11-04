// // Uncomment the entire file

package main

import "fmt"

func main() {

	var userEmails map[int]string = make(map[int]string)

	userEmails[1] = "user1@gmail.com"
	userEmails[2] = "user2@gmail.com"

	fmt.Println(userEmails)

	var wallets map[string]string = make(map[string]string)

	wallets["eth"] = "w90233209rfnafl29302n32k"
	wallets["btc"] = "a322l30000akmm3o2opakaanxx0002n"

	fmt.Println(wallets)

	// shorthand
	emails := map[int]string{
		1: "alpha@test.com",
		2: "beta@test.com",
	}

	fmt.Println(emails, len(emails))
	emails[5] = "epsilon@test.com"
	emails[1] = "Alpha@test.com"
	fmt.Println(emails, len(emails))

	// check if pair exists in map
	firstEmail := emails[1]
	fourthEmail, ok := emails[4] // use ok var to check if, fourthEmail will be empty string (default value)
	fmt.Println(firstEmail, fourthEmail, ok) // if (!ok) handle error

	// if (value, exists of collection; boolean condition) { execute}
	if email, ok := emails[3]; ok {
		fmt.Println("email exists:", email)
	} else {
		fmt.Println("email does not exist")
	}

	// deleting a key
	delete(emails, 2)

	fmt.Println(emails)

// 	// ****************************

	// var userEmails map[int]string = make(map[int]string)
	// // userEmails := make(map[int]string)

	// userEmails[1] = "user1@gmail.com"
	// userEmails[2] = "user2@gmail.com"

	// fmt.Println(userEmails)

// 	// ****************************

// 	userEmails := map[int]string{
// 		1: "user1@gmail.com",
// 		2: "user2@gmail.com",
// 	}

// 	fmt.Println(userEmails)

// 	// ****************************

// 	userEmails := map[int]string{
// 		1: "user1@gmail.com",
// 		2: "user2@gmail.com",
// 	}

// 	fmt.Println(userEmails)

// 	fmt.Println(userEmails[1])

// 	// ****************************

// 	userEmails := map[int]string{
// 		1: "user1@gmail.com",
// 		2: "user2@gmail.com",
// 	}

// 	fmt.Println(userEmails)

// 	userEmails[1] = "newUser1@gmail.com"

// 	fmt.Println(userEmails)

// 	fmt.Println(userEmails[3])

// 	// ****************************

// 	userEmails := map[int]string{
// 		1: "user1@gmail.com",
// 		2: "user2@gmail.com",
// 	}

// 	email1, ok := userEmails[1]
// 	fmt.Println("Email:", email1, "Present?", ok)

// 	// email3, ok := userEmails[3]
// 	// fmt.Println("Email", email3, "Present?", ok)

// 	// ****************************

// 	userEmails := map[int]string{
// 		1: "user1@gmail.com",
// 		2: "user2@gmail.com",
// 	}

// 	if email, ok := userEmails[1]; ok {
// 		fmt.Println(email)
// 	} else {
// 		fmt.Println("I don't know what you want from me")
// 	}

// 	// ****************************

// 	userEmails := map[int]string{
// 		1: "user1@gmail.com",
// 		2: "user2@gmail.com",
// 	}

// 	delete(userEmails, 2)

// 	fmt.Println(userEmails)
// 	// ****************************

// 	userEmails := map[int]string{
// 		1: "user1@gmail.com",
// 		2: "user2@gmail.com",
// 	}

// 	for k, v := range userEmails {
// 		fmt.Printf("%s has an ID of %d.\n", v, k)
// 	}
// 	// ****************************
}
