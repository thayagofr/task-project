package domain

import "time"

type Category string

const (
	Work      Category = "Work"
	Studies   Category = "Studies"
	Home      Category = "Home"
	Health    Category = "Health"
	Leisure   Category = "Leisure"
	Shopping  Category = "Shopping"
	Urgent    Category = "Urgent"
	Important Category = "Important"
	Social    Category = "Social"
	Projects  Category = "Projects"
)

type Status string

const (
	Open       Status = "Open"
	InProgress Status = "In Progress"
	Completed  Status = "Completed"
	Cancelled  Status = "Cancelled"
)

type TaskToRegister struct {
	OwnerID     string
	Name        string
	Description *string
	Category    Category
	DueDate     time.Time
}

type RegisteredTask struct {
	ID             string
	RegisteredDate time.Time
	Name           string
	Description    *string
	Category       Category
	Status         Status
	DueDate        time.Time
}
