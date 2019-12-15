package testmethod

import (
	"fmt"
	"testmethod/user"
)

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

func (s *Student) modifyName(name string) {
	s.name = name
}

func TestMethod() {
	s1 := new(Student)
	s1.age = 20
	s1.name = "xxx"

	s2 := new(Student)
	s2.age = 22
	s2.name = "yyy"
	fmt.Println("s1 compare with s2: ", s1.Compare(*s2))

	s2.modifyName("xingtianyu")
	fmt.Println("student 2 is:", s2)

	user := new(user.User)
	user.Init()
	fmt.Println("init user:", user)
	fmt.Println("user password is:", user.GetPassword())
	user.SetPassword("xxccvv")
	fmt.Println("user now is:", user)
}
