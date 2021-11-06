package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// BaseURL is the base endpoint for the star wars API
const BaseURL = "https://swapi.dev/api/"


// func getPerson(id int) {
// 	url := BaseURL + "people/" + string(id)
// 	fmt.Println(url)
// }

// StarWarsPeople represents all people in the Star Wars galaxy
type StarWarsPeople struct {
	Results []StarWarsPerson `json:"results"`
	Count int `json:"count"`
}

// StarWarsPerson represents a singular person in the Star Wars galaxy
type StarWarsPerson struct {
	Name string `json:"name"`
	BirthYear string `json:"birth_year"`
	HomeworldURL string `json:"homeworld"`
	Homeworld Planet
}

// Planet represents a planet in the Star Wars galaxy
type Planet struct {
	Name string `json:"name"`
	Population string `json:"population"`
	Terrain string `json:"terrain"`
}

// Method to find homeworld
func (person *StarWarsPerson) getHomeworld() {
	var r *http.Response
	var err error
	var bytes []byte
	// Make GET request
	if r, err = http.Get(person.HomeworldURL); err != nil {
		fmt.Print("Failed to fetch homeworld", err)
	}
	// Parse GET response
	if bytes, err = ioutil.ReadAll(r.Body); err != nil {
		fmt.Print("Failed to parse response body", err)
	}
	// Parse bytes into JSON, assign to person.Homeworld field
	if err := json.Unmarshal(bytes, &person.Homeworld); err != nil {
		fmt.Print("Error parsing JSON", err)
	}
}

func getPeople(w http.ResponseWriter) StarWarsPeople {
	url := BaseURL + "people"
	var r *http.Response
	var err error
	var bytes []byte
	var people StarWarsPeople
	// Make GET request
	if r, err = http.Get(url); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Print("Failed to request for Star Wars people", err)
	}
	// Parse GET response
	if bytes, err = ioutil.ReadAll(r.Body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Print("Failed to parse Star Wars people response", err)
	}
	// Parse bytes into JSON, assign to people
	if err := json.Unmarshal(bytes, &people); err != nil {
			fmt.Print("Error parsing JSON", err)
	}
	// Get planet for each person, call method
	for _, p := range people.Results {
		p.getHomeworld()
		// fmt.Println(p) // TODO: planet data is available here, but not in people
	}
	// Return results
	return people
}

func peopleHandler(w http.ResponseWriter, r *http.Request) {
	people := getPeople(w)
	// View log to see results
	fmt.Println(people)
	fmt.Fprintf(w, "<p>Response was successful</p>")
}


func main() {
	// Routes
	http.HandleFunc("/people", peopleHandler)

	// Init server
	fmt.Println("HTTP server listening on port 7777....")
	log.Fatal(http.ListenAndServe("localhost:7777", nil))
}
