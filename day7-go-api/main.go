package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

// CORS許可ミドルウェア
func enableCORS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

        if r.Method == http.MethodOptions {
            return
        }

        next.ServeHTTP(w, r)
    })
}


// 初期データ
var users = []User{
	{Name: "Rin", Role: "engineer"},
	{Name: "Kan", Role: "designer"},
}

// GET /users
func usersHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

// POST /add
func addUser(w http.ResponseWriter, r *http.Request) {
	var newUser User

	json.NewDecoder(r.Body).Decode(&newUser)

	users = append(users, newUser)

	json.NewEncoder(w).Encode(users)
}

// POST /delete
func deleteUser(w http.ResponseWriter, r *http.Request) {
    var req struct {
        Index int `json:"index"`
    }

    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    if req.Index < 0 || req.Index >= len(users) {
        http.Error(w, "Index out of range", http.StatusBadRequest)
        return
    }

    users = append(users[:req.Index], users[req.Index+1:]...)

    json.NewEncoder(w).Encode(users)
}




func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("/users", usersHandler)
    mux.HandleFunc("/add", addUser)
    mux.HandleFunc("/delete", deleteUser)

    http.ListenAndServe(":8080", enableCORS(mux))
}
