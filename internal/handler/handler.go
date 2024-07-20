package handler

import (
	"HL_project_management/internal/model"
	"HL_project_management/internal/repository"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

var validate = validator.New()

// @Summary Health check
// @Description Health check
// @Tags health
// @Produce plain
// @Success 200 {string} string "OK"
// @Router /health [get]
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// @Summary Get all users
// @Description Get all users
// @Tags users
// @Produce json
// @Success 200 {array} model.User
// @Failure 500 {string} string "Internal server error"
// @Router /users [get]
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := repository.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

// @Summary Create a new user
// @Description Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.User true "User data"
// @Success 201 {object} model.User
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Internal server error"
// @Router /users [post]
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	err := validate.Struct(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.RegistrationAt = time.Now()
	createdUser, err := repository.CreateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}

// @Summary Get user by ID
// @Description Get user by ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} model.User
// @Failure 400 {string} string "Invalid ID"
// @Failure 404 {string} string "User not found"
// @Router /users/{id} [get]
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	user, err := repository.GetUserByID(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

// @Summary Update user
// @Description Update user
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body model.User true "User data"
// @Success 200 {object} model.User
// @Failure 400 {string} string "Invalid input"
// @Failure 404 {string} string "User not found"
// @Failure 500 {string} string "Internal server error"
// @Router /users/{id} [put]
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	err = validate.Struct(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}

	updatedUser, err := repository.UpdateUser(id, user)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(updatedUser)
}

// @Summary Delete user
// @Description Delete user
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {string} string "Deleted successfully"
// @Failure 400 {string} string "Invalid ID"
// @Failure 404 {string} string "User not found"
// @Router /users/{id} [delete]
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	err = repository.DeleteUser(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted successfully")
}

// @Summary Get tasks by user ID
// @Description Get tasks by user ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {array} model.Task
// @Failure 400 {string} string "Invalid ID"
// @Failure 404 {string} string "Tasks not found"
// @Router /users/{id}/tasks [get]
func GetTasksByUserID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	_, err = repository.GetUserByID(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	tasks, err := repository.GetTasksByUserID(id)
	if err != nil {
		http.Error(w, "Tasks not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

// @Summary Search users by name or email
// @Description Search users by name or email
// @Tags users
// @Produce json
// @Param name query string false "User name"
// @Param email query string false "User email"
// @Success 200 {array} model.User
// @Failure 400 {string} string "Invalid input"
// @Router /search/users [get]
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	email := r.URL.Query().Get("email")
	users, err := repository.SearchUsers(name, email)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(users)
}

// @Summary Get all tasks
// @Description Get all tasks
// @Tags tasks
// @Produce json
// @Success 200 {array} model.Task
// @Failure 500 {string} string "Internal server error"
// @Router /tasks [get]
func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := repository.GetAllTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

// @Summary Create a new task
// @Description Create a new task
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body model.Task true "Task data"
// @Success 201 {object} model.Task
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Internal server error"
// @Router /tasks [post]
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task model.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validate.Struct(task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	task.CreatedAt = time.Now()
	if task.CompletedAt.Before(task.CreatedAt) && !task.CompletedAt.IsZero() {
		http.Error(w, "Completed date should be after created date", http.StatusBadRequest)
		return

	}
	if task.CompletedAt.IsZero() {
		task.CompletedAt = time.Now().AddDate(0, 1, 0)

	}
	createdTask, err := repository.CreateTask(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdTask)
}

// @Summary Get task by ID
// @Description Get task by ID
// @Tags tasks
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} model.Task
// @Failure 400 {string} string "Invalid ID"
// @Failure 404 {string} string "Task not found"
// @Router /tasks/{id} [get]
func GetTaskByID(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	task, err := repository.GetTaskByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(task)
}

// @Summary Update task
// @Description Update task
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Param task body model.Task true "Task data"
// @Success 200 {object} model.Task
// @Failure 400 {string} string "Invalid input"
// @Failure 404 {string} string "Task not found"
// @Failure 500 {string} string "Internal server error"
// @Router /tasks/{id} [put]
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var task model.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validate.Struct(task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = repository.GetTaskByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return

	}

	task, err = repository.UpdateTask(id, task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(task)
}

// @Summary Delete task
// @Description Delete task
// @Tags tasks
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {string} string "Deleted successfully"
// @Failure 400 {string} string "Invalid ID"
// @Failure 404 {string} string "Task not found"
// @Router /tasks/{id} [delete]
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	_, err = repository.GetTaskByID(id)
	if err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	if err := repository.DeleteTask(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted successfully")
}

// @Summary Search tasks
// @Description Search tasks by title, priority, status, assignee, or project
// @Tags tasks
// @Produce json
// @Param title query string false "Task title"
// @Param priority query string false "Task priority"
// @Param status query string false "Task status"
// @Param assignee query int false "Assignee ID"
// @Param project query int false "Project ID"
// @Success 200 {array} model.Task
// @Failure 400 {string} string "Invalid input"
// @Router /search/tasks [get]
func SearchTasks(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	priority := r.URL.Query().Get("priority")
	status := r.URL.Query().Get("status")
	assigneeID, err := strconv.Atoi(r.URL.Query().Get("assignee"))
	projectID, err := strconv.Atoi(r.URL.Query().Get("project"))
	tasks, err := repository.SearchTasks(title, priority, status, assigneeID, projectID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tasks)

}

// @Summary Get all projects
// @Description Get all projects
// @Tags projects
// @Produce json
// @Success 200 {array} model.Project
// @Failure 500 {string} string "Internal server error"
// @Router /projects [get]
func GetAllProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := repository.GetAllProjects()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(projects)
}

// @Summary Create a new project
// @Description Create a new project
// @Tags projects
// @Accept json
// @Produce json
// @Param project body model.Project true "Project data"
// @Success 201 {object} model.Project
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Internal server error"
// @Router /projects [post]
func CreateProject(w http.ResponseWriter, r *http.Request) {
	var project model.Project
	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validate.Struct(project); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	project.StartDate = time.Now()
	if project.EndDate.Before(project.StartDate) && !project.EndDate.IsZero() {
		http.Error(w, "End date should be after start date", http.StatusBadRequest)
		return
	}
	if project.EndDate.IsZero() {
		project.EndDate = time.Now().AddDate(1, 0, 0)
	}

	project, err := repository.CreateProject(project)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(project)
}

// @Summary Get project by ID
// @Description Get project by ID
// @Tags projects
// @Produce json
// @Param id path int true "Project ID"
// @Success 200 {object} model.Project
// @Failure 400 {string} string "Invalid ID"
// @Failure 404 {string} string "Project not found"
// @Router /projects/{id} [get]
func GetProjectByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	project, err := repository.GetProjectByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(project)
}

// @Summary Update project
// @Description Update project
// @Tags projects
// @Accept json
// @Produce json
// @Param id path int true "Project ID"
// @Param project body model.Project true "Project data"
// @Success 200 {object} model.Project
// @Failure 400 {string} string "Invalid input"
// @Failure 404 {string} string "Project not found"
// @Failure 500 {string} string "Internal server error"
// @Router /projects/{id} [put]
func UpdateProject(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var project model.Project
	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validate.Struct(project); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = repository.GetProjectByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	project, err = repository.UpdateProject(id, project)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(project)
}

// @Summary Delete project
// @Description Delete project
// @Tags projects
// @Produce json
// @Param id path int true "Project ID"
// @Success 200 {string} string "Deleted successfully"
// @Failure 400 {string} string "Invalid ID"
// @Failure 404 {string} string "Project not found"
// @Router /projects/{id} [delete]
func DeleteProject(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	_, err = repository.GetProjectByID(id)
	if err != nil {
		http.Error(w, "Project not found", http.StatusNotFound)
		return

	}

	if err := repository.DeleteProject(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted successfully")
}

// @Summary Get tasks by project ID
// @Description Get tasks by project ID
// @Tags projects
// @Produce json
// @Param id path int true "Project ID"
// @Success 200 {array} model.Task
// @Failure 400 {string} string "Invalid ID"
// @Failure 404 {string} string "Tasks not found"
// @Router /projects/{id}/tasks [get]
func GetTasksByProjectID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	_, err = repository.GetProjectByID(id)
	if err != nil {
		http.Error(w, "Project not found", http.StatusNotFound)
		return
	}
	tasks, err := repository.GetTasksByProjectID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

// @Summary Search projects
// @Description Search projects by title or manager
// @Tags projects
// @Produce json
// @Param title query string false "Project title"
// @Param manager query int false "Manager ID"
// @Success 200 {array} model.Project
// @Failure 400 {string} string "Invalid input"
// @Router /search/projects [get]
func SearchProjects(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	managerID, err := strconv.Atoi(r.URL.Query().Get("manager"))
	projects, err := repository.SearchProjects(title, managerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(projects)
}
