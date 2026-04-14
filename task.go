package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Priority int

const (
	PriorityLow Priority = iota
	PriorityMedium
	PriorityHigh
)

func (p Priority) String() string {
	switch p {
	case PriorityLow:
		return "low"
	case PriorityMedium:
		return "medium"
	case PriorityHigh:
		return "high"
	default:
		return "unknown"
	}
}

type Task struct {
	ID          uuid.UUID
	Title       string
	Description string
	Priority    Priority
	CreatedAt   time.Time
	DueAt       time.Time
	Done        bool
}

func NewTask(title, description string, priority Priority, dueAt time.Time) *Task {
	return &Task{
		ID:          uuid.New(),
		Title:       title,
		Description: description,
		Priority:    priority,
		CreatedAt:   time.Now(),
		DueAt:       dueAt,
	}
}

func (t *Task) Complete() {
	t.Done = true
}

func (t *Task) IsOverdue() bool {
	return !t.Done && time.Now().After(t.DueAt)
}

func (t *Task) String() string {
	status := "pending"
	if t.Done {
		status = "done"
	} else if t.IsOverdue() {
		status = "overdue"
	}
	return fmt.Sprintf("[%s] %s (priority: %s, status: %s, due: %s)",
		t.ID, t.Title, t.Priority, status, t.DueAt.Format(time.RFC3339))
}
