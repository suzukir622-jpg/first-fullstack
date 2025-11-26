package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// ===== User 関連 =====

type User struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

// ユーザー一覧（グローバル変数）
var users = []User{
	{Name: "Rin", Role: "engineer"},
	{Name: "Kan", Role: "designer"},
}

// GET /users
func usersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "encode error", http.StatusInternalServerError)
		return
	}
}

// POST /add
func addUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newUser User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if newUser.Name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	users = append(users, newUser)

	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "encode error", http.StatusInternalServerError)
		return
	}
}

// POST /delete
func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Index int `json:"index"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Index < 0 || req.Index >= len(users) {
		http.Error(w, "Index out of range", http.StatusBadRequest)
		return
	}

	users = append(users[:req.Index], users[req.Index+1:]...)

	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "encode error", http.StatusInternalServerError)
		return
	}
}

// ===== Task 関連 =====

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// タスクリスト（グローバル変数）
var tasks = []Task{
	{ID: 1, Name: "Study Go"},
	{ID: 2, Name: "Study AWS"},
}

// GET /tasks
func TasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		http.Error(w, "encode error", http.StatusInternalServerError)
		return
	}
}

// POST /tasks/add
func AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	// 新しいID（最後のID + 1）
	newID := 1
	if len(tasks) > 0 {
		newID = tasks[len(tasks)-1].ID + 1
	}

	newTask := Task{
		ID:   newID,
		Name: req.Name,
	}

	tasks = append(tasks, newTask)

	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		http.Error(w, "encode error", http.StatusInternalServerError)
		return
	}
}

// ===== CORS ミドルウェア =====

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 必要なら Origin を絞ってもOK
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			// プリフライトリクエストはここで終了
			return
		}

		next.ServeHTTP(w, r)
	})
}

// ===== main =====

func main() {
	mux := http.NewServeMux()

	// User系
	mux.HandleFunc("/users", usersHandler)
	mux.HandleFunc("/add", addUser)
	mux.HandleFunc("/delete", deleteUser)

	// Task系
	mux.HandleFunc("/tasks", TasksHandler)
	mux.HandleFunc("/tasks/add", AddTaskHandler)

	log.Println("server starting on :8080")
	if err := http.ListenAndServe(":8080", enableCORS(mux)); err != nil {
		log.Fatal(err)
	}
}
