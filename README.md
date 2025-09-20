# Inspired by [roadmap.sh](https://roadmap.sh/projects/task-tracker)
# Task Tracker CLI

Task Tracker CLI is a simple command-line tool written in Go for managing tasks using a local JSON file as a database. It allows you to add, update, delete, and list tasks directly from your terminal.

## Features
- Add new tasks with descriptions
- Update task descriptions and status
- Delete tasks
- List all tasks or filter by status
- Persistent storage in a JSON file

## Usage

### Build
To build the CLI tool, run:

```
go build -o task-cli cmd/server/server.go
```

This will create an executable named `task-cli` in your current directory.

### Run
Basic usage:

```
./task-cli <action> [parameters]
```

#### Actions
- `add <description>`: Add a new task with the given description
- `update <id> <description>`: Update the description of a task by ID
- `delete <id>`: Delete a task by ID
- `update-status <id> <status>`: Update the status of a task
- `list`: List all tasks
- `list <status>`: List tasks filtered by status

### Example
```
./task-cli add "Buy groceries"
./task-cli update 1 "Buy groceries and cook dinner"
./task-cli update-status 1 done
./task-cli list
./task-cli list todo
```

## Data Storage
Tasks are stored in `db/tasks.json` as a JSON object. The file is created automatically if it does not exist.

## Requirements
- Go 1.18 or newer

## License
MIT
