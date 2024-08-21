# a Go project to create a cli tool to add perform some crud operations for a simple task tracking tool.

# CLI Package : [Cobra-Cli](https://github.com/spf13/cobra)
```bash
#add new  command to for cobra
$ cobra-cli init
$ cobra-cli add <task name>
```

#Go
```bash
#Run the project
$ go run main.go
#build the project
$ go build

#Run the project with the flag
$ go run main.go list -a
$ ./task list -a #built
```

# Todo
- [x] Add a flag to list all tasks
- - [x] load the tasks from a file(csv)
- - [x] function to display the tasks
- - [ ] function to display short task list
- [ ] Add task command
- - [x] open file and append new task
- - [x] add new task and update ID 
- [ ] Add a flag to show a specific task
- [ ] Add a flag to delete a specific task
- [ ] Add a flag to edit a specific task
- [ ] Add a flag to add a new task