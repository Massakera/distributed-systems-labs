package main

import (
	"time"
	"log"
	"net/rpc"
	"github.com/Massakera/MapReduce/mapreduce"
)

func main() {
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
			log.Printf("Worker %d is processing job %d", reply.WorkerID, reply.WorkerID)
			time.Sleep(1 * time.Second)
			args.WorkerID = reply.WorkerID
			err = client.Call("CoordinatorServer.TaskCompleted", args, reply)
			if err != nil {
				log.Fatalf("RPC error: %v", err)
			}
		} else {
			break
		}

		// Add a small delay before the next iteration
		time.Sleep(100 * time.Millisecond)
	}
}
