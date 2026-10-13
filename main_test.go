package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestCommandsPersistTaskChanges(t *testing.T) {
	path := filepath.Join(t.TempDir(), "tasks.json")
	store, err := loadTaskStore(path)
	if err != nil {
		t.Fatal(err)
	}

	if err := run([]string{"add", "write tests"}, store); err != nil {
		t.Fatal(err)
	}
	if err := run([]string{"mark-in-progress", "1"}, store); err != nil {
		t.Fatal(err)
	}
	if err := run([]string{"update", "1", "write thorough tests"}, store); err != nil {
		t.Fatal(err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	var tasks []task
	if err := json.Unmarshal(data, &tasks); err != nil {
		t.Fatal(err)
	}
	if len(tasks) != 1 {
		t.Fatalf("got %d tasks, want 1", len(tasks))
	}
	if tasks[0].Description != "write thorough tests" || tasks[0].Status != statusInProgress {
		t.Fatalf("got task %+v, want updated description and in-progress status", tasks[0])
	}
}

func TestDeleteTask(t *testing.T) {
	store := &taskStore{path: filepath.Join(t.TempDir(), "tasks.json")}
	if err := run([]string{"add", "temporary task"}, store); err != nil {
		t.Fatal(err)
	}
	if err := run([]string{"delete", "1"}, store); err != nil {
		t.Fatal(err)
	}
	if len(store.tasks) != 0 {
		t.Fatalf("got %d tasks after delete, want 0", len(store.tasks))
	}
}

func TestRunRejectsUnknownTask(t *testing.T) {
	store := &taskStore{path: filepath.Join(t.TempDir(), "tasks.json")}
	if err := run([]string{"delete", "42"}, store); err == nil {
		t.Fatal("expected an error when deleting a missing task")
	}
}
