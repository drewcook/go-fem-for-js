package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Home</h1><p>This is the home page.</p>")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About</h1><p>This is the about page.</p>")
}


// To-Do's
type Todo struct {
	Title string
	Content string
}
type TodosPageData struct {
	PageTitle string
	PageTodos []Todo
}
// setup some mock data to display
var todos []Todo

func todosHandler(w http.ResponseWriter, r *http.Request) {
	pageData := TodosPageData{
		PageTitle: "Get Todos",
		PageTodos: todos,
	}

	t, err := template.ParseFiles("todos.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Template parsing error:", err)
	}

	err = t.Execute(w, pageData)
}

func addTodoHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Request parsing error:", err)
	}
	todo := Todo{
		Title: r.FormValue("title"),
		Content: r.FormValue("content"),
	}
	todos = append(todos, todo)

	log.Print(todos)
	http.Redirect(w, r, "/todos", http.StatusSeeOther)
}

func main() {
	// Route handlers
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/todos", todosHandler)
	http.HandleFunc("/todos/add-todo", addTodoHandler)

	// Listen
	fmt.Println("Go HTTP server running on port 7777...")
	log.Fatal(http.ListenAndServe(":7777", nil))
}
