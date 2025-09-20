package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type Tasks struct {
	Tasks map[string]Task `json:"tasks"`
}

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

func main() {
	arguments := os.Args[1:]
	if len(arguments) < 1 {
		usage()
		return
	}
	action := arguments[0]
	log.Println("Action:", action)

	dbLoc := "db/tasks.json"
	tasks := initJSONdb(dbLoc)

	switch action {
	case "add":
		if len(arguments) < 2 {
			usage()
			log.Fatal("Missing description")
		}
		addTask(tasks, arguments[1])
		commitToDB(tasks, dbLoc)
	case "update":
		if len(arguments) < 3 {
			usage()
			log.Fatal("Missing something")
		}
		updateTask(tasks, arguments[1], arguments[2])
		commitToDB(tasks, dbLoc)
	case "delete":
		if len(arguments) < 2 {
			usage()
			log.Fatal("Missing id")
		}
		deleteTask(tasks, arguments[1])
		commitToDB(tasks, dbLoc)
	case "mark-in-progress":
		if len(arguments) < 2 {
			usage()
			log.Fatal("Missing id")
		}
		updateStatus(tasks, arguments[1], "in-progress")
		commitToDB(tasks, dbLoc)
	case "mark-done":
		if len(arguments) < 2 {
			usage()
			log.Fatal("Missing id")
		}
		updateStatus(tasks, arguments[1], "done")
		commitToDB(tasks, dbLoc)
	case "list":
		if len(arguments) < 2 {
			listAllTasks(tasks)
		} else {
			listTasksByStatus(tasks, arguments[1])
		}

	default:
		usage()
	}
}

func usage() {
	fmt.Println("Usage: task-cli <action> [parameters]")
	fmt.Println("Actions:")
	fmt.Println("  add <description>                Add a new task")
	fmt.Println("  update <id> <description>        Update a task's description")
	fmt.Println("  delete <id>                      Delete a task")
	fmt.Println("  mark-in-progress <id>            Mark a task as in-progress")
	fmt.Println("  mark-done <id>                   Mark a task as done")
	fmt.Println("  list                             List all tasks")
	fmt.Println("  list <status>                    List tasks filtered by status")
}

func initJSONdb(path string) Tasks {

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		log.Println("File does not exists, creating one....")
		err := os.WriteFile(path, []byte("{\"tasks\": {  } }"), 0666)
		if err != nil {
			log.Fatal(err)
		}
	}

	db, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("Error loading DB")
	}

	tasks := Tasks{}
	err = json.Unmarshal(db, &tasks)
	if err != nil {
		log.Fatal(err)
	}

	return tasks

}

func commitToDB(tasks Tasks, path string) {

	log.Println("Commiting tasks to permanent storage.......")
	data, err := json.Marshal(tasks)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(path, data, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB save completed!")
}

func addTask(tasks Tasks, description string) {
	log.Println("Adding a new task...")
	id := len(tasks.Tasks) + 1
	task := Task{}
	task.ID = id
	task.Status = "todo"
	task.Description = description
	task.CreatedAt = time.Now().String()
	task.UpdatedAt = ""

	tasks.Tasks[strconv.Itoa(id)] = task
	log.Printf("Task added successfully (ID: %d)", id)
}

func updateTask(tasks Tasks, id string, description string) {
	if taskExists(tasks, id) {
		task := tasks.Tasks[id]
		task.Description = description
		task.UpdatedAt = time.Now().String()
		tasks.Tasks[id] = task
	} else {
		log.Println("Task does not exist")
	}

}

func deleteTask(tasks Tasks, id string) {
	if taskExists(tasks, id) {
		delete(tasks.Tasks, id)
	} else {
		log.Println("Task does not exist")
	}

}

func updateStatus(tasks Tasks, id string, status string) {
	if taskExists(tasks, id) {
		task := tasks.Tasks[id]
		task.Status = status
		task.UpdatedAt = time.Now().String()
		tasks.Tasks[id] = task
	} else {
		log.Println("Task does not exist")
	}

}

func listAllTasks(tasks Tasks) {
	for id, task := range tasks.Tasks {
		fmt.Println(id, task)

	}
}

func listTasksByStatus(tasks Tasks, status string) {
	for id, task := range tasks.Tasks {
		if task.Status == status {
			fmt.Println(id, task)
		}

	}
}

func taskExists(tasks Tasks, id string) bool {
	_, ok := tasks.Tasks[id]
	return ok
}
