package main

import (
	"fmt"
	"time"
)

const (
	TaskStatusTodo       TaskStatus = "TODO"
	TaskStatusInProgress TaskStatus = "in-progress"
	TaskStatusDone       TaskStatus = "done"
)

type Task struct {
	id          int
	description string
	status      TaskStatus
	createdAt   time.Time
	updatedAt   time.Time
}

func (t *Task) printTask() {
	fmt.Printf("ID: %d, Description: %s, Status: %s, Created At: %s, Updated At: %s\n",
		t.id,
		t.description,
		t.status,
		t.createdAt.Format(time.RFC3339),
		t.updatedAt.Format(time.RFC3339))
}
