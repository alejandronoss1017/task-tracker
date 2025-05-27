package models

import (
	"fmt"
	"time"
)

// Task represents a to-do item with its metadata (id, description, status, and timestamps).
type Task struct {
	Id          int       `json:"id"`          // Unique identifier for the task
	Description string    `json:"description"` // Short description of the task
	Status      string    `json:"status"`      // The status of the task (todo, in-progress, done)
	CreatedAt   time.Time `json:"created_at"`  // The date and time when the task was created
	UpdatedAt   time.Time `json:"updated_at"`  // The date and time when the task was last updated
}

// String returns a readable summary of the Task.
// It satisfies the fmt.Stringer interface so you can do fmt.Println(task).
func (t Task) String() string {
	return fmt.Sprintf(
		"Id: %d, Description: %s, Status: %s, Created at: %s, Updated at: %s",
		t.Id, t.Description, t.Status, t.CreatedAt, t.UpdatedAt)
}
