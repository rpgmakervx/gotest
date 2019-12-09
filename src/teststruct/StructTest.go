package teststruct

import (
	"fmt"
	"time"
	"unsafe"
)

type JobInfo struct {
	id      string
	jobType int
	time    int64
}
type Object struct {
	name string
}

func TestStruct() {
	testDefine()
	testVisit()
	testEquals()
	testSize()
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

func testEquals() {
	var job1 = JobInfo{"job-1", 1, 1}
	var job2 = JobInfo{"job-1", 1, 1}
	var job3 = JobInfo{"job-3", 3, 1}
	fmt.Printf("job1 addr:%p job2 addr2:%p, job3 addr:%p\n", &job1, &job2, &job3)
	fmt.Println("job1 == job2? :", job1 == job2)
	fmt.Println("job1 == job3? :", job1 == job3)
	fmt.Println("job1 addr == job2 addr ? :", &job1 == &job3)
}

func testSize() {
	var object = Object{name: "12345678901234567890"}
	var p uintptr
	fmt.Printf("object len:%d\n", unsafe.Sizeof(object))
	fmt.Printf("string len:%d\n", unsafe.Sizeof(""))
	fmt.Printf("uintptr len:%d\n", unsafe.Sizeof(p))
}
