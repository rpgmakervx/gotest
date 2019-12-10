package testmethod

import "fmt"

type Student struct {
	age  int
	name string
}

func (student Student) Compare(s Student) int {
	if s.age > student.age {
		return -1
	} else if s.age < student.age {
		return 1
	}
	return 0
}

func TestMethod() {
	s1 := new(Student)
	s1.age = 20
	s1.name = "xxx"

	s2 := new(Student)
	s2.age = 22
	s2.name = "yyy"
	fmt.Println("s1 compare with s2: ", s1.Compare(*s2))
}
