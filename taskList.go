package main

import (
	"time"
)

type TaskList struct {
	tasks []Task
}

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

func (tl *TaskList) GetTaskById(id int) Task {
	index := tl.GetTaskIndexById(id)
	if index != -1 {
		return tl.tasks[index]
	}
	return Task{}
}

func (tl *TaskList) GetTaskIndexById(id int) int {
	for i, task := range tl.tasks {
		if task.id == id {
			return i
		}
	}
	return -1
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