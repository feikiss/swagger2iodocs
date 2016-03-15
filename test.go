package main

import "fmt"

type Student struct  {
	name,age string
}

type class struct {
	stu *Student
	level int
}

func main() {
	student := &Student{
		name:"Jack",
	}

	cla := class{
		stu:student,
		level:2,
	}

	fmt.Println(cla.stu.name)

	update(student)
	fmt.Println(cla.stu.name)

}

func update(stu *Student)  {
	stu.name = "from update."
}