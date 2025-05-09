package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	todoList := TodoList{}
	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Println("======== TODO LIST===========")

		fmt.Println("1. Add Task")
		fmt.Println("2. View Task")
		fmt.Println("3. Mark Task as done")
		fmt.Println("4. Exit")

		fmt.Println("===================")
		fmt.Print("Enter Your Choice: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter task name: ")
			scanner.Scan()
			newTaskName := scanner.Text()
			todoList.AddTask(newTaskName)
		case "2":
			todoList.ViewTask()
		case "3":
			fmt.Print("Choice task ID: ")
			scanner.Scan()
			taskIdStr := scanner.Text()

			taskId, err := strconv.Atoi(taskIdStr)
			if err != nil {
				fmt.Println("Invalid Task Id")
				continue
			}

			todoList.MarkAsDone(taskId)
		case "4":
			fmt.Println("Exiting.......")
			return
		default:
			fmt.Println("Invalid Choice, Please Try Again......")
		}

	}

}
