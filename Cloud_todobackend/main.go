package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"log"
	"net/http"
	"todo_app/handlers"
	"todo_app/services"

	"github.com/rs/cors"
)

func main() {
	// Initialize database connection
	dsn := "avnadmin:AVNS_PmC-dmRPUSKgmMVcD49@tcp(mysql-c9ed2de-exampledemo46-9bf9.l.aivencloud.com:22180)/todo_app"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()

	// Initialize TodoService
	todoService := &services.TodoServiceImpl{DB: db}

	// Initialize TodoHandler with TodoService
	todoHandler := handlers.NewTodoHandler(todoService)

	// Define routes
	mux := http.NewServeMux()

	mux.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			todoHandler.GetAllTodos(w, r)
		case http.MethodPost:
			todoHandler.CreateTodo(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/todos/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			todoHandler.GetTodoByID(w, r)
		} else if r.Method == http.MethodPut {
			todoHandler.UpdateTodo(w, r)
		} else if r.Method == http.MethodDelete {
			todoHandler.DeleteTodo(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// CORS middleware
	corsHandler := cors.Default().Handler(mux)

	// Start the server
	port := ":8080"
	fmt.Println("Server started on port", port)
	log.Fatal(http.ListenAndServe(port, corsHandler))
}
