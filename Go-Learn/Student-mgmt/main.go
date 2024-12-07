package main

import (
	"fmt"
)

type Student struct {
	Id    uint32
	Name  string
	Marks map[string]int
}

var students = []Student{}
var choice int
var st Student
var id uint32
var updatedName string

func main() {

	for {
		if choice == 6 {
			panic("Exiting the programe")
		}
		fmt.Println(`
			1 for adding new student
			2 for displaying all students
			3 for searching student with thier id
			4 for updating student name
			5 for deleting student
			6 for exit

			Enter your choice:
		`)

		fmt.Scan(&choice)
		switch choice {
			case 1:
				fmt.Println(addStudents())
			case 2:
				displayAllStudents()
			case 3:
				inputStudent()
				displayStudent(id)
			case 4:
				inputStudent()
				fmt.Println("Enter updated name:")
				fmt.Scan(&updatedName)
				updateStudent(id, updatedName)
			case 5:
				inputStudent()
				fmt.Println(deleteStudent(id))
		}
	}
}