package teststruct

import (
	"fmt"
	"time"
)

type JobInfo struct {
	id      string
	jobType int
	time    int64
}

func TestStruct() {
	testDefine()
	testVisit()
}

func testDefine() {
	var job = JobInfo{"job-1", 1, time.Now().Unix() * 1e3}
	fmt.Println("job info:", job)
	var anotherJob = JobInfo{id: "job-2", jobType: 1}
	fmt.Println("anotherJob info:", anotherJob)

	var pJob *JobInfo = &JobInfo{"hi", 10, time.Now().Unix() * 1e3}
	fmt.Println("pJob info:", *pJob)
	// pJob.id 和 (*pJob).id 等价的
	fmt.Println("pJob jobId:", pJob.id)

	var newJob = new(JobInfo)
	newJob.id = "newJob-1"
	newJob.time = time.Now().Unix() * 1e3
	newJob.jobType = 15
	fmt.Println("new job:", newJob)
}

func testVisit() {
	var job JobInfo
	job.time = time.Now().Unix() * 1e3
	job.id = "job-2"
	job.jobType = 2
	fmt.Println("job is:", job)
}
