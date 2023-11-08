package documents

import "time"

type Task struct {
	Name        string    `bson:"name"`
	Category    string    `bson:"category"`
	Description string    `bson:"description"`
	DueDate     time.Time `bson:"due_date"`
}
