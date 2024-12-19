package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"os"
)

var students = []Student{}
var st Student

func inputStudent() {
	fmt.Println("Enter id:")
	fmt.Scan(&id)
}

func addStudents() string {
	var name string
	var engMarks, hindiMarks, mathsMarks, csMarks, sciMarks int
	fmt.Println("Enter name and marks of student")
	fmt.Scan(&name, &engMarks, &hindiMarks, &mathsMarks, &csMarks, &sciMarks)

	st = Student{
		Id:   int(uuid.New().ID()),
		Name: name,
		Marks: map[string]int{
			"eng":   engMarks,
			"hindi": hindiMarks,
			"maths": mathsMarks,
			"cs":    csMarks,
			"sci":   sciMarks,
		},
	}
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err.Error())
	}
	err = json.Unmarshal(data, &students)
	if err != nil {
		panic(err.Error())
	}
	students = append(students, st)
	marshelledData, err := json.Marshal(students)
	if err != nil {
		panic(err.Error())
	}
	err = os.WriteFile(fileName, marshelledData, 0644)
	if err != nil {
		panic(err.Error())
	}
	return "Student added"
}

func displayAllStudents() {
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err.Error())
	}
	err = json.Unmarshal(data, &students)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(students)
}

func displayStudent(id int) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err.Error())
	}
	err = json.Unmarshal(data, &students)
	if err != nil {
		panic(err.Error())
	}

	for _, val := range students {
		if val.Id == id {
			fmt.Println(val)
			return
		}
	}
	fmt.Println("No result found")
}

func updateStudent(id int, updName string) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err.Error())
	}
	err = json.Unmarshal(data, &students)
	if err != nil {
		panic(err.Error())
	}
	for i, val := range students {
		if val.Id == id {
			students[i].Name = updName
			return
		}
	}
}

func deleteStudent(id int) string {
	newStudents := []Student{}
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err.Error())
	}
	err = json.Unmarshal(data, &students)
	if err != nil {
		panic(err.Error())
	}
	for i := 0; i < len(students); i++ {
		if students[i].Id != id {
			newStudents = append(newStudents, students[i])
		}
	}

	marshelledData, err := json.Marshal(newStudents)
	if err != nil {
		panic(err.Error())
	}
	err = os.WriteFile(fileName, marshelledData, 0644)
	if err != nil {
		panic(err.Error())
	}
	return "Student deleted"
}
