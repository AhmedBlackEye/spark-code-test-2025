package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
)

const Port = "8080"

type Todo struct {
	Id          uint32 `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

var (
	nextId uint32
	todos  []Todo
	mu     sync.Mutex
)

func main() {
	http.HandleFunc("/", toDoListHandler)
	http.ListenAndServe(":"+Port, nil)
}

func toDoListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	switch r.Method {

	case http.MethodGet:
		getTodos(w, r)

	case http.MethodPost:
		createTodo(w, r)

	case http.MethodDelete:
		deleteTodo(w, r)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func getTodos(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item Todo

	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if item.Title == "" || item.Description == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	mu.Lock()
	item.Id = nextId
	nextId++
	todos = append(todos, item)
	mu.Unlock()

	json.NewEncoder(w).Encode(item)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for i, t := range todos {
		if t.Id == uint32(id) {
			todos = append(todos[:i], todos[i+1:]...)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"message": "Todo deleted"})
			return
		}
	}
	http.Error(w, "Todo not found", http.StatusNotFound)
}
