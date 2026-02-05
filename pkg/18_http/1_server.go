package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// ref: https://go-cookbook.com/snippets/http/routing-http-requests

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	mux := http.NewServeMux()

	// RESTful API routes with method-specific handlers.
	mux.HandleFunc("GET /api/users", listUsers)
	mux.HandleFunc("POST /api/users", createUser)
	mux.HandleFunc("GET /api/users/{id}", getUser)
	mux.HandleFunc("PUT /api/users/{id}", updateUser)
	mux.HandleFunc("DELETE /api/users/{id}", deleteUser)

	http.ListenAndServe(":8080", mux)
}

func listUsers(w http.ResponseWriter, r *http.Request) {
	users := []User{{1, "Alice"}, {2, "Bob"}}
	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	user.ID = 3 // Simulate ID assignment
	json.NewEncoder(w).Encode(user)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	user := User{ID: id, Name: "Example User"}
	json.NewEncoder(w).Encode(user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	fmt.Fprintf(w, "Updating user %s\n", idStr)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	fmt.Fprintf(w, "Deleting user %s\n", idStr)
}
