package main

import "fmt"

type TodoList struct {
	tasks []Task
}

func (td *TodoList) AddTask(taskName string) {
	taskId := len(td.tasks) + 1
	task := NewTask(taskId, taskName)
	td.tasks = append(td.tasks, task)
	fmt.Println("Added task successfully...")
}

func (td *TodoList) ViewTask() {
	fmt.Println("Tasks are viewed...", td.tasks)
}

func (td *TodoList) MarkAsDone(taskId int) {
	fmt.Println("Task is done:", taskId)
}
