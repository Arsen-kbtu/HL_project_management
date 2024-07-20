package model

import "time"

type User struct {
	ID             int       `json:"id" readonly:"true"`
	Name           string    `json:"name" validate:"required"`
	Email          string    `json:"email" validate:"required,email" example:"string@gmail.com"`
	RegistrationAt time.Time `json:"registrationAt" readonly:"true"`
	Role           string    `json:"role" validate:"required"`
}

type Task struct {
	ID          int       `json:"id" readonly:"true" `
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description"`
	Priority    string    `json:"priority"  validate:"required,oneof=low medium high"`
	Status      string    `json:"status"`
	AssigneeID  int       `json:"assigneeId" validate:"required" example:"1"`
	ProjectID   int       `json:"projectId" validate:"required" example:"1"`
	CreatedAt   time.Time `json:"createdAt" readonly:"true"`
	CompletedAt time.Time `json:"completedAt" example:"2024-09-20T15:04:05Z"`
}

type Project struct {
	ID          int       `json:"id" readonly:"true"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"startDate" readonly:"true"`
	EndDate     time.Time `json:"endDate" example:"2024-09-20T15:04:05Z"`
	ManagerID   int       `json:"managerId" validate:"required" example:"1"`
}
