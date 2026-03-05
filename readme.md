# Go CLI Task Manager

A simple **terminal-based task manager built with Go**.
This project allows users to manage tasks directly from the command line.

It demonstrates core Go concepts such as:

* Structs
* File handling
* JSON encoding/decoding
* CLI arguments
* Goroutines
* Modular project structure

---

# Features

* Add new tasks
* List all tasks
* Mark tasks as completed
* Delete tasks
* Persistent storage using JSON
* Asynchronous logging using Goroutines
* Simple CLI interface

---

# Project Structure

```
go-task-manager
│
├── main.go        # CLI command handling
├── task.go        # Task structure
├── storage.go     # Load and save tasks
├── logger.go      # Async logging using goroutines
├── tasks.json     # Local storage file
├── go.mod
└── README.md
```

---

# Requirements

* Go 1.20 or higher

Check Go version:

```
go version
```

---

# Installation

Clone the repository:

```
git clone https://github.com/your-username/go-task-manager.git
```

Navigate to the project folder:

```
cd go-task-manager
```

Initialize modules (if needed):

```
go mod init taskmanager
```

---

# Running the Project

Run the application using:

```
go run .
```

---

# Usage

## Add Task

```
go run . add "Learn Go"
```

## List Tasks

```
go run . list
```

Example output:

```
1 | Learn Go | false
2 | Build CLI App | false
```

## Mark Task Completed

```
go run . done 1
```

## Delete Task

```
go run . delete 1
```

---

# Build Executable

Build the binary:

```
go build
```

This creates an executable file:

```
go-task-manager.exe
```

Run it directly:

```
./go-task-manager add "Learn Go"
./go-task-manager list
```

---

# Data Storage

Tasks are stored locally in:

```
tasks.json
```

Example structure:

```
[
 {
  "id": 1,
  "title": "Learn Go",
  "completed": false
 }
]
```

---

# Concepts Covered

This project demonstrates several important Go concepts:

* Structs
* JSON encoding and decoding
* File operations
* Command line arguments
* Goroutines
* Modular project design
* Error handling

---

# Future Improvements

Possible enhancements:

* Add task priorities
* Add due dates
* Search tasks
* Use SQLite database
* Add colored CLI output
* Use Cobra CLI framework
* Docker support

---

# Author

Sahil Raj

MERN Stack Developer
Frontend & Backend Development
