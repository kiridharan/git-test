package main

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

// Todo represents a task in the todo list.
type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var todos []Todo

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", index).Methods("GET")
	r.HandleFunc("/todos", getTodos).Methods("GET")
	r.HandleFunc("/todos", createTodo).Methods("POST")

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("index").Parse(`
<!DOCTYPE html>
<html>
<head>
    <title>Todo List</title>
</head>
<body>
    <h1>Todo List</h1>
    <form id="todoForm" action="/todos" method="post">
        <input type="text" name="title" placeholder="Enter a new todo" />
        <button type="submit">Add</button>
    </form>
    
    <h2>New Todos:</h2>
    <ul id="newTodos">
        {{range .}}
        <li>{{.Title}}</li>
        {{end}}
    </ul>
    <script>
        document.getElementById("todoForm").addEventListener("submit", function(event) {
            event.preventDefault();
            const input = document.querySelector("input[name='title']");
            const title = input.value;
            
            fetch("/todos", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({ title: title })
            })
            .then(response => response.json())
            .then(todo => {
                input.value = ""; // Clear the input field
                const ul = document.getElementById("newTodos");
                const li = document.createElement("li");
                li.textContent = todo.title;
                ul.appendChild(li);
            });
        });
    </script>
</body>
</html>
`))
	tmpl.Execute(w, todos)
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	todo.ID = len(todos) + 1
	todos = append(todos, todo)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}
