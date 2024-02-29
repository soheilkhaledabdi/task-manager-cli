package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const taskFileName = "tasks.txt"

func main() {
	start()

	reader := bufio.NewReader(os.Stdin)
	command, _ := reader.ReadString('\n')
	processTask(strings.TrimSpace(command))
}

func start() {
	fmt.Println("Welcome to the task manager")
	fmt.Println("Enter a number to perform a task:")
	fmt.Println("1. Add a task")
	fmt.Println("2. Delete a task by ID")
	fmt.Println("3. List tasks")
	fmt.Println("4. Exit")
}

func processTask(command string) {
		switch command {
		case "1":
			addTask()
		case "2":
			deleteTask()
		case "3":
			listTasks()
		case "4":
			exit()
		default:
			fmt.Println("Invalid command. Please enter a valid number.")
		}
}

func addTask() {
	fmt.Print("Enter task description: ")
	reader := bufio.NewReader(os.Stdin)
	description, _ := reader.ReadString('\n')
	if err := appendToFile(taskFileName, description); err != nil {
		fmt.Println("Error adding task:", err)
		return
	}
	fmt.Println("Task added successfully!")
}

func deleteTask() {
	listTasks()

	fmt.Print("Enter the ID of the task to delete: ")
	reader := bufio.NewReader(os.Stdin)
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID. Please enter a valid number.")
		return
	}

	tasks, err := readTasksFromFile(taskFileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if id < 1 || id > len(tasks) {
		fmt.Println("Invalid ID. Please enter a valid ID.")
		return
	}

	deletedTask := tasks[id-1]
	tasks = append(tasks[:id-1], tasks[id:]...)

	if err := writeTasksToFile(taskFileName, tasks); err != nil {
		fmt.Println("Error deleting task:", err)
		return
	}

	fmt.Printf("Task \"%s\" deleted successfully!\n", deletedTask)
}

func listTasks() {
	tasks, err := readTasksFromFile(taskFileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}

	fmt.Println("Tasks:")
	for i, task := range tasks {
		fmt.Printf("%d. %s\n", i+1, task)
	}
}

func appendToFile(fileName, content string) error {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.WriteString(content); err != nil {
		return err
	}
	return nil
}

func readTasksFromFile(fileName string) ([]string, error) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}, nil
		}
		return nil, err
	}
	tasks := strings.Split(string(content), "\n")
	return tasks[:len(tasks)-1], nil
}

func writeTasksToFile(fileName string, tasks []string) error {
	content := strings.Join(tasks, "\n") + "\n"
	if err := ioutil.WriteFile(fileName, []byte(content), 0644); err != nil {
		return err
	}
	return nil
}

func exit() {
	fmt.Println("Bye :)")
}
