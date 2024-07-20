package router

import (
	"HL_project_management/internal/handler"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// Swagger docs
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	r.HandleFunc("/health", handler.HealthCheck).Methods("GET")

	r.HandleFunc("/users", handler.GetAllUsers).Methods("GET")
	r.HandleFunc("/users", handler.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", handler.GetUserByID).Methods("GET")
	r.HandleFunc("/users/{id}", handler.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", handler.DeleteUser).Methods("DELETE")
	r.HandleFunc("/users/{id}/tasks", handler.GetTasksByUserID).Methods("GET")
	r.HandleFunc("/search/users", handler.SearchUsers).Methods("GET")

	r.HandleFunc("/tasks", handler.GetAllTasks).Methods("GET")
	r.HandleFunc("/tasks", handler.CreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", handler.GetTaskByID).Methods("GET")
	r.HandleFunc("/tasks/{id}", handler.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", handler.DeleteTask).Methods("DELETE")
	r.HandleFunc("/search/tasks", handler.SearchTasks).Methods("GET")
	//
	r.HandleFunc("/projects", handler.GetAllProjects).Methods("GET")
	r.HandleFunc("/projects", handler.CreateProject).Methods("POST")
	r.HandleFunc("/projects/{id}", handler.GetProjectByID).Methods("GET")
	r.HandleFunc("/projects/{id}", handler.UpdateProject).Methods("PUT")
	r.HandleFunc("/projects/{id}", handler.DeleteProject).Methods("DELETE")
	r.HandleFunc("/projects/{id}/tasks", handler.GetTasksByProjectID).Methods("GET")
	r.HandleFunc("/search/projects", handler.SearchProjects).Methods("GET")

	// Default handler for unsupported methods
	r.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	})

	return r
}
