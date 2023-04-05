package main

import (
	"time"
	"log"
	"net/rpc"
	"github.com/Massakera/MapReduce/mapreduce" // Replace with your actual GitHub username
)

func main() {
	// Connect to the coordinator
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatalf("Error connecting to the coordinator: %v", err)
	}
	defer client.Close()

	for {
		args := &mapreduce.TaskArgs{}
		reply := &mapreduce.TaskReply{}
		err := client.Call("CoordinatorServer.GetTask", args, reply)
		if err != nil {
			log.Fatalf("RPC error: %v", err)
		}

		if reply.Success {
			// Simulate processing the job
			log.Printf("Worker %d is processing job %d", reply.WorkerID, reply.WorkerID)
			time.Sleep(1 * time.Second)

			// Notify the coordinator that the job is completed
			args.WorkerID = reply.WorkerID
			err = client.Call("CoordinatorServer.TaskCompleted", args, reply)
			if err != nil {
				log.Fatalf("RPC error: %v", err)
			}
		} else {
			// No more jobs available
			break
		}

		// Add a small delay before the next iteration
		time.Sleep(100 * time.Millisecond)
	}
}
