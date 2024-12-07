package main

import (
	"fmt"
	"github.com/google/uuid"
)

func inputStudent(){
	fmt.Println("Enter id:")
	fmt.Scan(&id)
}

func addStudents()[]Student{
	var name string
	var engMarks, hindiMarks, mathsMarks, csMarks, sciMarks int
	fmt.Println("Enter name and marks of student")
	fmt.Scan(&name, &engMarks, &hindiMarks, &mathsMarks, &csMarks, &sciMarks)

	st = Student{
		Id: uuid.New().ID(),
		Name: name,
		Marks: map[string]int{
			"eng": engMarks,
			"hindi": hindiMarks,
			"maths": mathsMarks,
			"cs": csMarks,
			"sci": sciMarks,
		},
	}
	students = append(students, st)
	return students
}

func displayAllStudents(){
	for _, val:= range students{
		fmt.Println(val)
	}
}

func displayStudent(id uint32){
	for i:=0;i<len(students);i++{
		if students[i].Id == uint32(id){
			fmt.Println(students[i])
			return
		}
	}
}

func updateStudent(id uint32, updName string){
		for i:= 0;i<len(students);i++{
			if students[i].Id == id{
				students[i].Name = updName
			}
		}
}

func deleteStudent(id uint32) []Student {
	newStudents := []Student{} 
	for i := 0; i < len(students); i++ {
		if students[i].Id != id {
			newStudents = append(newStudents, students[i]) 
		}
	}
	students = newStudents 
	return students
}
