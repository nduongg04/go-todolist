package main

import (
	"log"
	"net/http"
	"todolist-api/internal/config"
	"todolist-api/internal/handlers"
	"todolist-api/internal/middleware"
	"todolist-api/internal/repository"
	"todolist-api/internal/services"
	"todolist-api/prisma/db"

	"github.com/gorilla/mux"
)

func initDB() *db.PrismaClient {
	db := db.NewClient()
	if err := db.Prisma.Connect(); err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	return db
}

func main() {
	r := mux.NewRouter()
	db := initDB()

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config: ", err)
	}

	userService := services.NewUserService(repository.NewUserRepository(db))
	userHandler := handlers.NewUserHandler(userService)

	todoService := services.NewTodoService(repository.NewTodoRepository(db))
	todoHandler := handlers.NewTodoHandler(todoService)

	v1 := r.PathPrefix("/api/v1").Subrouter()

	auth := v1.PathPrefix("/auth").Subrouter()
	// Routes
	auth.HandleFunc("/register", userHandler.Register).Methods(http.MethodPost)
	auth.HandleFunc("/login", userHandler.Login).Methods(http.MethodPost)

	api := v1.PathPrefix("").Subrouter()
	api.Use(middleware.AuthMiddleWare(config))

	api.HandleFunc("/todos", todoHandler.GetTodoByUserID).Methods(http.MethodGet)
	api.HandleFunc("/todos", todoHandler.CreateTodo).Methods(http.MethodPost)
	api.HandleFunc("/todos/{id}", todoHandler.GetTodo).Methods(http.MethodGet)
	api.HandleFunc("/todos/{id}", todoHandler.UpdateTodo).Methods(http.MethodPatch)
	api.HandleFunc("/todos/{id}", todoHandler.DeleteTodo).Methods(http.MethodDelete)
	// CORS middleware
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
			next.ServeHTTP(w, r)
		})
	})

	log.Println("Server starting on port " + config.Port + "...")
	log.Fatal(http.ListenAndServe(":"+config.Port, r))
}
