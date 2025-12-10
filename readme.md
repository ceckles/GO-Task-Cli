# GO-Task-Cli

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Cobra](https://img.shields.io/badge/Cobra-000000?style=for-the-badge&logo=cobra&logoColor=white)

A simple and efficient command-line task management tool built with Go. Manage your tasks with full CRUD operations directly from your terminal.

## 🚀 Features

- ✅ **Add tasks** - Quickly add new tasks to your list
- 📋 **List tasks** - View all tasks or filter by status
- ✓ **Complete tasks** - Mark tasks as done
- 🗑️ **Delete tasks** - Remove tasks and auto-resequence IDs
- 💾 **CSV storage** - Simple file-based storage
- ⏰ **Time tracking** - See when tasks were created with relative time

## 📋 Table of Contents

- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [Commands](#commands)
- [Development](#development)
- [Project Structure](#project-structure)
- [Tech Stack](#tech-stack)

## 🔧 Prerequisites

- **Go 1.18+** - [Download Go](https://golang.org/dl/)
- **Git** - For cloning the repository

## 📦 Installation

### Option 1: Build from Source

1. **Clone the repository:**
   ```bash
   git clone https://github.com/ceckles/GO-Task-Cli.git
   cd GO-Task-Cli
   ```

2. **Install dependencies:**
   ```bash
   go mod download
   ```

3. **Build the project:**
   ```bash
   go build -o task
   ```

4. **Install globally (optional):**
   ```bash
   # Linux/macOS
   sudo mv task /usr/local/bin/
   
   # Or add to your PATH
   export PATH=$PATH:$(pwd)
   ```

### Option 2: Run Directly

You can run the tool directly without building:

```bash
go run main.go [command]
```

## 🎯 Usage

### Basic Commands

**List all tasks (short view):**
```bash
task list
# or
go run main.go list
```

**List all tasks with full details:**
```bash
task list -a
# or
go run main.go list -a
```

**Add a new task:**
```bash
task add "Complete the project documentation"
# or
go run main.go add "Complete the project documentation"
```

**Mark a task as complete:**
```bash
task complete 1
# or
go run main.go complete 1
```

**Delete a task:**
```bash
task delete 1
# or
go run main.go delete 1
```

**Show help:**
```bash
task --help
task [command] --help
```

## 📚 Commands

### `list`

List all tasks. By default shows a short view (ID, Task, Created time). Use the `-a` or `--all` flag to see all details including completion status.

```bash
task list          # Short view
task list -a       # Full view with all details
task list --all    # Same as -a
```

### `add`

Add a new task to your list. The task will be assigned an auto-incremented ID.

```bash
task add "Your task description here"
```

**Example:**
```bash
$ task add "Review pull request #42"
Task added: Review pull request #42
```

### `complete`

Mark a task as complete by providing its ID. The command will check if the task is already completed.

```bash
task complete [task-id]
```

**Example:**
```bash
$ task complete 1
Task completed successfully
```

### `delete`

Delete a task by its ID. All remaining task IDs will be automatically re-sequenced.

```bash
task delete [task-id]
```

**Example:**
```bash
$ task delete 2
Task ID 2 deleted and remaining IDs re-sequenced.
```

## 🛠️ Development

### Project Setup

1. **Install Go dependencies:**
   ```bash
   go mod download
   ```

2. **Run the application:**
   ```bash
   go run main.go
   ```

3. **Build the binary:**
   ```bash
   go build -o task
   ```

### Adding New Commands with Cobra

This project uses [Cobra](https://github.com/spf13/cobra) for CLI command management.

**Install Cobra CLI (if not already installed):**
```bash
go install github.com/spf13/cobra-cli@latest
```

**Add a new command:**
```bash
cobra-cli add <command-name>
```

**Example:**
```bash
cobra-cli add edit
```

This will create a new command file in the `cmd/` directory that you can customize.

### Running Tests

```bash
go test ./...
```

### File Structure

The tasks are stored in a CSV file (`tasks.csv`) with the following format:
```
ID,Task,Created,Done
1,Example task,2024-01-15 10:30:00 -0500 EST,false
```

## 📁 Project Structure

```
GO-Task-Cli/
├── cmd/
│   ├── root.go      # Root command and CLI setup
│   ├── add.go       # Add task command
│   ├── list.go      # List tasks command
│   ├── complete.go  # Complete task command
│   └── delete.go    # Delete task command
├── utils/
│   ├── file_utils.go  # File operations utilities
│   └── time_utils.go  # Time formatting utilities
├── main.go          # Application entry point
├── go.mod           # Go module dependencies
├── go.sum           # Go module checksums
├── tasks.csv        # Task storage file (created on first use)
└── readme.md        # This file
```

## 🏗️ Tech Stack

- **Go** - Programming language
- **Cobra** - CLI framework for building command-line applications
- **timediff** - Library for human-readable time differences

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
