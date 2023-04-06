# MapReduce in Go

This project implements a simple MapReduce system in Go using the net/rpc package. It consists of a coordinator and worker nodes. The coordinator is responsible for distributing tasks to worker nodes, tracking their status, and managing the completion of jobs. The worker nodes connect to the coordinator, request tasks, process the tasks, and notify the coordinator once the tasks are completed.

## Project Structure

- `coordinator.go`: The coordinator server implementation. It listens for connections from worker nodes, assigns tasks, and tracks the status of both workers and jobs.
- `worker.go`: The worker node implementation. It connects to the coordinator, requests tasks, processes the tasks, and notifies the coordinator once the tasks are completed.
- `mapreduce.go`: The common data structures and types used by both the coordinator and worker nodes.

## Getting Started

### Prerequisites

- Install [Go](https://golang.org/doc/install) (version 1.16 or later).

### Running the Project

1. Clone the repository.
```bash
git clone https://github.com/Massakera/distributed-systems-go.git
```
2. Change of directory and start the coordinator 
```bash
cd MapReduce
go run coordinator.go
```

3. Start the worker in another window
```bash
go run worker.go
```

### Customization

You can customize the project by modifying the Job struct and the mapreduce.go file to handle specific tasks or data processing logic.
