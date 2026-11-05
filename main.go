package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type TaskStatus string

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

type TaskList struct {
	tasks []Task
}

var taskList TaskList

func (tl *TaskList) AddTask(task Task) int {
	taskId := 0

	if len(tl.tasks) > 0 {
		lastTask := tl.tasks[len(tl.tasks)-1]
		taskId = lastTask.id
	}

	task.id = taskId + 1
	task.status = TaskStatusTodo
	task.createdAt = time.Now()
	task.updatedAt = time.Now()
	tl.tasks = append(tl.tasks, task)
	return task.id
}

func (tl *TaskList) GetTaskIndexById(id int) int {
	for i, task := range tl.tasks {
		if task.id == id {
			return i
		}
	}
	return -1
}

func (tl *TaskList) GetTaskById(id int) Task {
	index := tl.GetTaskIndexById(id)
	if index != -1 {
		return tl.tasks[index]
	}
	return Task{}
}

func (tl *TaskList) UpdateTaskDescription(id int, updatedDescription string) {
	index := tl.GetTaskIndexById(id)
	if index != -1 {
		tl.tasks[index].description = updatedDescription
		tl.tasks[index].updatedAt = time.Now()
	}
}

func (tl *TaskList) DeleteTaskById(id int) {
	index := tl.GetTaskIndexById(id)
	if index != -1 {
		tl.tasks = append(tl.tasks[:index], tl.tasks[index+1:]...)
	}
}

func (tl *TaskList) UpdateTaskStatus(id int, updatedStatus TaskStatus) {
	index := tl.GetTaskIndexById(id)
	if index != -1 {
		tl.tasks[index].status = updatedStatus
		tl.tasks[index].updatedAt = time.Now()
	}
}

func (tl *TaskList) UpdateTaskStatusInProgress(id int) {
	tl.UpdateTaskStatus(id, TaskStatusInProgress)
}

func (tl *TaskList) UpdateTaskStatusDone(id int) {
	tl.UpdateTaskStatus(id, TaskStatusDone)
}

func (t *Task) printTask() {
	fmt.Printf("ID: %d, Description: %s, Status: %s, Created At: %s, Updated At: %s\n",
		t.id,
		t.description,
		t.status,
		t.createdAt.Format(time.RFC3339),
		t.updatedAt.Format(time.RFC3339))
}

func (tl *TaskList) PrintAllTask() {
	for _, task := range tl.tasks {
		task.printTask()
	}
}

func (tl *TaskList) PrintAllTaskByStatus(status TaskStatus) {
	for _, task := range tl.tasks {
		if task.status == status {
			task.printTask()
		}
	}
}

func processCommand(input string) {
	splittedInput := strings.Fields(input)
	keyword := strings.ToLower(splittedInput[0])
	switch keyword {
		case "add":
			// Handle add command
		case "update":
			// Handle update command
		case "delete":
			// Handle delete command
		case "list":
			// Handle list command
		default:
			fmt.Println("Unknown command")
	}
}

func addCommand(splittedInput []string) {
	taskDescription := strings.Join(splittedInput[1:], " ")
	taskDescription = strings.ReplaceAll(taskDescription, "\"", "")

	newTask := Task{
		description: taskDescription,
	}
	taskList.AddTask(newTask)
	fmt.Println("Task added successfully!")
}

func





func main() {
	fmt.Println("Welcome to the Task List Application!")
	
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()

		processCommand(input)
	
	
	}
}
	
