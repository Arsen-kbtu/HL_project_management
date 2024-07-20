package repository

import (
	"HL_project_management/internal/model"
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log"
)

type Config struct {
	Port int
	Env  string
	Db   struct {
		Dsn string
	}
}

var db *sql.DB

func OpenDB(cfg Config) (*sql.DB, error) {
	// Use sql.Open() to create an empty connection pool, using the DSN from the config // struct.
	var err error
	db, err = sql.Open("postgres", cfg.Db.Dsn)
	fmt.Println(db)
	if err != nil {
		return nil, err
	}
	migrationUp(db)
	return db, nil
}
func migrationUp(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Путь к файлам миграции
	//m, err := migrate.NewWithDatabaseInstance(
	//	"file:///usr/src/app/internal/migrations",
	//	"postgres", driver)
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	if err != nil {
		log.Fatal(err)
	}

	// Применение миграций
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}
func CloseDB() {
	if db != nil {
		db.Close()
	}
}

// User functions
func GetAllUsers() ([]model.User, error) {
	rows, err := db.Query("SELECT id, name, email, registration_at, role FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.RegistrationAt, &user.Role); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func CreateUser(user model.User) (model.User, error) {
	err := db.QueryRow(
		"INSERT INTO users (name, email, registration_at, role) VALUES ($1, $2, $3, $4) RETURNING id",
		user.Name, user.Email, user.RegistrationAt, user.Role,
	).Scan(&user.ID)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func GetUserByID(id int) (model.User, error) {
	var user model.User
	err := db.QueryRow("SELECT id, name, email, registration_at, role FROM users WHERE id = $1", id).
		Scan(&user.ID, &user.Name, &user.Email, &user.RegistrationAt, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.User{}, err
		}
		return model.User{}, err
	}
	return user, nil
}

func UpdateUser(id int, user model.User) (model.User, error) {
	_, err := db.Exec(
		"UPDATE users SET name = $1, email = $2, role = $3 WHERE id = $4",
		user.Name, user.Email, user.Role, id,
	)
	if err != nil {
		return model.User{}, err
	}
	return GetUserByID(id)
}

func DeleteUser(id int) error {
	_, err := db.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}

func GetTasksByUserID(userID int) ([]model.Task, error) {
	rows, err := db.Query("SELECT id, title, description, priority, status, assignee_id, project_id, created_at, completed_at FROM tasks WHERE assignee_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var task model.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Priority, &task.Status, &task.AssigneeID, &task.ProjectID, &task.CreatedAt, &task.CompletedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func SearchUsers(name string, email string) ([]model.User, error) {
	var rows *sql.Rows
	var err error
	if name != "" && email != "" {
		rows, err = db.Query("SELECT id, name, email, registration_at, role FROM users WHERE name = $1 AND email = $2", name, email)
	} else if name != "" {
		rows, err = db.Query("SELECT id, name, email, registration_at, role FROM users WHERE name = $1", name)
	} else if email != "" {
		rows, err = db.Query("SELECT id, name, email, registration_at, role FROM users WHERE email = $1", email)
	} else {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.RegistrationAt, &user.Role); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// Task functions
func GetAllTasks() ([]model.Task, error) {
	rows, err := db.Query("SELECT id, title, description, priority, status, assignee_id, project_id, created_at, completed_at FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var task model.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Priority, &task.Status, &task.AssigneeID, &task.ProjectID, &task.CreatedAt, &task.CompletedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func CreateTask(task model.Task) (model.Task, error) {
	err := db.QueryRow(
		"INSERT INTO tasks (title, description, priority, status, assignee_id, project_id, created_at, completed_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id",
		task.Title, task.Description, task.Priority, task.Status, task.AssigneeID, task.ProjectID, task.CreatedAt, task.CompletedAt,
	).Scan(&task.ID)
	if err != nil {
		return model.Task{}, err
	}
	return task, nil
}

func GetTaskByID(id int) (model.Task, error) {
	var task model.Task
	err := db.QueryRow("SELECT id, title, description, priority, status, assignee_id, project_id, created_at, completed_at FROM tasks WHERE id = $1", id).
		Scan(&task.ID, &task.Title, &task.Description, &task.Priority, &task.Status, &task.AssigneeID, &task.ProjectID, &task.CreatedAt, &task.CompletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Task{}, err
		}
		return model.Task{}, err
	}
	return task, nil
}

func UpdateTask(id int, task model.Task) (model.Task, error) {
	_, err := db.Exec(
		"UPDATE tasks SET title = $1, description = $2, priority = $3, status = $4, assignee_id = $5, project_id = $6, completed_at = $7 WHERE id = $8",
		task.Title, task.Description, task.Priority, task.Status, task.AssigneeID, task.ProjectID, task.CompletedAt, id,
	)
	if err != nil {
		return model.Task{}, err
	}
	return GetTaskByID(id)
}

func DeleteTask(id int) error {
	_, err := db.Exec("DELETE FROM tasks WHERE id = $1", id)
	return err
}

func SearchTasks(title, priority, status string, assigneeID, projectID int) ([]model.Task, error) {
	query := fmt.Sprintf(
		`
		SELECT   *
		FROM tasks
		WHERE (STRPOS(LOWER(title), LOWER($1)) > 0 OR $1= '')
		AND (STRPOS(LOWER(priority), LOWER($2)) > 0 or $2 = '')
		AND (STRPOS(LOWER(status), LOWER($3)) > 0 or $3 = '')
		AND ($4 = 0 OR assignee_id = $4)
		AND ($5 = 0 OR project_id = $5)
		`)
	rows, err := db.Query(query, title, priority, status, assigneeID, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tasks []model.Task

	for rows.Next() {
		var task model.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Priority, &task.Status, &task.AssigneeID, &task.ProjectID, &task.CreatedAt, &task.CompletedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)

	}
	return tasks, nil
}

// Project functions
func GetAllProjects() ([]model.Project, error) {
	rows, err := db.Query("SELECT id, title, description, start_date, end_date, manager_id FROM projects")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []model.Project
	for rows.Next() {
		var project model.Project
		if err := rows.Scan(&project.ID, &project.Title, &project.Description, &project.StartDate, &project.EndDate, &project.ManagerID); err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}

	return projects, nil
}

func CreateProject(project model.Project) (model.Project, error) {
	var err error
	if project.EndDate.IsZero() {
		err = db.QueryRow(
			"INSERT INTO projects (title, description, start_date, end_date, manager_id) VALUES ($1, $2, $3, $4, $5) RETURNING id",
			project.Title, project.Description, project.StartDate, sql.NullTime{}, project.ManagerID,
		).Scan(&project.ID)

	} else {
		err = db.QueryRow(
			"INSERT INTO projects (title, description, start_date, end_date, manager_id) VALUES ($1, $2, $3, $4, $5) RETURNING id",
			project.Title, project.Description, project.StartDate, project.EndDate, project.ManagerID,
		).Scan(&project.ID)
	}
	if err != nil {
		return model.Project{}, err
	}
	return project, nil
}

func GetProjectByID(id int) (model.Project, error) {
	var project model.Project
	err := db.QueryRow("SELECT id, title, description, start_date, end_date, manager_id FROM projects WHERE id = $1", id).
		Scan(&project.ID, &project.Title, &project.Description, &project.StartDate, &project.EndDate, &project.ManagerID)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Project{}, err
		}
		return model.Project{}, err
	}
	return project, nil
}

func UpdateProject(id int, project model.Project) (model.Project, error) {
	_, err := db.Exec(
		"UPDATE projects SET title = $1, description = $2, start_date = $3, end_date = $4, manager_id = $5 WHERE id = $6",
		project.Title, project.Description, project.StartDate, project.EndDate, project.ManagerID, id,
	)
	if err != nil {
		return model.Project{}, err
	}
	return GetProjectByID(id)
}

func DeleteProject(id int) error {
	_, err := db.Exec("DELETE FROM projects WHERE id = $1", id)
	return err
}

func GetTasksByProjectID(projectID int) ([]model.Task, error) {
	rows, err := db.Query("SELECT id, title, description, priority, status, assignee_id, project_id, created_at, completed_at FROM tasks WHERE project_id = $1", projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var task model.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Priority, &task.Status, &task.AssigneeID, &task.ProjectID, &task.CreatedAt, &task.CompletedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func SearchProjects(title string, managerID int) ([]model.Project, error) {
	query := fmt.Sprintf(
		`
		SELECT   *
		FROM projects
		WHERE (STRPOS(LOWER(title), LOWER($1)) > 0 OR $1= '')
		AND ($2 = 0 OR manager_id = $2)
		`)
	rows, err := db.Query(query, title, managerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var projects []model.Project

	for rows.Next() {
		var project model.Project
		err := rows.Scan(&project.ID, &project.Title, &project.Description, &project.StartDate, &project.EndDate, &project.ManagerID)
		if err != nil {
			return nil, err
		}
		projects = append(projects, project)

	}
	return projects, nil

}
