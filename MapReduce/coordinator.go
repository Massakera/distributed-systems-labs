package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"sync"
	"time"
	"github.com/Massakera/MapReduce/mapreduce"
)

type CoordinatorServer struct {
	coordinator  *mapreduce.Coordinator
	workerStatus map[int]bool
	mu           sync.Mutex
	completedJobs map[int]bool
}

func (c *CoordinatorServer) GetTask(args *mapreduce.TaskArgs, reply *mapreduce.TaskReply) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	for i, job := range c.coordinator.Jobs {
		if !c.workerStatus[i] && !c.completedJobs[i] {
			reply.WorkerID = i
			reply.Job = job
			reply.Success = true
			c.workerStatus[i] = true
			return nil
		}
	}

	reply.Success = false
	return nil
}


func (c *CoordinatorServer) waitForCompletion(workerID int) {
	time.Sleep(10 * time.Second)
	c.mu.Lock()
	c.workerStatus[workerID] = false
	c.mu.Unlock()
}

func (c *CoordinatorServer) TaskCompleted(args *mapreduce.TaskArgs, reply *mapreduce.TaskReply) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.workerStatus[args.WorkerID] = false
	c.completedJobs[args.WorkerID] = true
	reply.Success = true

	return nil
}


func main() {
	coordinator := mapreduce.NewCoordinator(5)
	coordinatorServer := &CoordinatorServer{
		coordinator:  coordinator,
		workerStatus: make(map[int]bool),
    completedJobs: make(map[int]bool),
	}

	err := rpc.Register(coordinatorServer)
	if err != nil {
		log.Fatal("Failed to register RPC server:", err)
	}

	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Failed to listen on :1234:", err)
	}

	fmt.Println("Coordinator server is running on :1234")
	http.Serve(listener, nil)
}
