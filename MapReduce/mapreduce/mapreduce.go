package mapreduce

import "time"

const TaskTimeout = 10 * time.Second

type Job struct {
  Id int
  Filename string
}

type Coordinator struct {
  Jobs []*Job
  WorkerCount int
}

type TaskArgs struct {
  Job *Job
  WorkerID int
}

func NewCoordinator(numJobs int) *Coordinator {
	jobs := make([]*Job, numJobs)
	for i := 0; i < numJobs; i++ {
		jobs[i] = &Job{Id: i}
	}

	return &Coordinator{
		Jobs: jobs,
	}
}

type TaskReply struct {
  Success bool
  Job *Job
  WorkerID  int
}


