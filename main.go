package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TaskStatus string

var taskList TaskList

func processCommand(input string) {
	splittedInput := strings.Fields(input)
	keyword := strings.ToLower(splittedInput[0])
	switch keyword {
	case "add":
		addCommand(splittedInput)
	case "update":
		updateCommand(splittedInput)
	case "delete":
		deleteCommand(splittedInput)
	case "list":
		listCommand(splittedInput)
	default:
		fmt.Println("Unknown command")
	}
}

func addCommand(splittedInput []string) {
	taskDescription := concatenateAndCleanQuotedString(splittedInput[1:])

	newTask := Task{
		description: taskDescription,
	}
	taskList.AddTask(newTask)
	fmt.Println("Task added successfully!")
}

func updateCommand(splittedInput []string) {
	if len(splittedInput) < 3 {
		fmt.Println("Invalid update command. Usage: update <task_id> <new_description>")
		return
	}

	taskId := splittedInput[1]
	newDescription := concatenateAndCleanQuotedString(splittedInput[2:])

	id, err := strconv.Atoi(taskId)
	if err != nil {
		fmt.Println("Invalid task ID. Please provide a valid integer.")
		return
	}

	taskList.UpdateTaskDescription(id, newDescription)
	fmt.Println("Task updated successfully!")
}

func deleteCommand(splittedInput []string) {
	if len(splittedInput) < 2 {
		fmt.Println("Invalid delete command. Usage: delete <task_id>")
		return
	}

	taskId := splittedInput[1]

	id, err := strconv.Atoi(taskId)
	if err != nil {
		fmt.Println("Invalid task ID. Please provide a valid integer.")
		return
	}

	taskList.DeleteTaskById(id)
	fmt.Println("Task deleted successfully!")
}

func concatenateAndCleanQuotedString(quotedString []string) string {
	concatenatedString := strings.Join(quotedString, " ")
	return strings.ReplaceAll(concatenatedString, "\"", "")
}

func listCommand(splittedInput []string) {
	if len(splittedInput) < 2 {
		taskList.PrintAllTask()
	} else {
		status := TaskStatus(splittedInput[1])
		taskList.PrintAllTaskByStatus(status)
	}
}

func main() {
	fmt.Println("Welcome to the Task List Application!")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()

		processCommand(input)

	}
}
