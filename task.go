package main

import "fmt"

type Task struct {
	ID   int
	Name string
	Done bool
}

func (t *Task) taskAsDone() {
	if t.Done {
		fmt.Println("Task already done...")
	}
	t.Done = true
}

func NewTask(taskId int, taskName string) Task {
	return Task{
		Name: taskName,
		ID:   taskId,
	}
}
