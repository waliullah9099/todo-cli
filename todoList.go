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

	if taskId < 1 || taskId > len(td.tasks) {
		fmt.Println("Invalid task id")
		return
	}
	td.tasks[taskId-1].taskAsDone()
}
