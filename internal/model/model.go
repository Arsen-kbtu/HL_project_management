package model

import "time"

type User struct {
	ID             int       `json:"id" readOnly:"true"`
	Name           string    `json:"name" validate:"required"`
	Email          string    `json:"email" validate:"required,email" example:"string@gmail.com"`
	RegistrationAt time.Time `json:"registrationAt" readOnly:"true"`
	Role           string    `json:"role" validate:"required"`
}

type Task struct {
	ID          int       `json:"id" readOnly:"true`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description"`
	Priority    string    `json:"priority"  validate:"required,oneof=low medium high"`
	Status      string    `json:"status"`
	AssigneeID  int       `json:"assigneeId" validate:"required"`
	ProjectID   int       `json:"projectId" validate:"required"`
	CreatedAt   time.Time `json:"createdAt" readOnly:"true"`
	CompletedAt time.Time `json:"completedAt"`
}

type Project struct {
	ID          int       `json:"id" readOnly:"true`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"startDate" readOnly:"true`
	EndDate     time.Time `json:"endDate"`
	ManagerID   int       `json:"managerId" validate:"required"`
}
