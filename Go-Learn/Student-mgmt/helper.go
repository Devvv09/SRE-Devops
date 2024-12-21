package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"os"
)

const addStudentMsg = "student added!"
const deleteStudentMsg = "student deleted!"
const updateStudentMsg = "student updated!"

var students = []*Student{}
var st *Student

func inputStudent() {
	fmt.Println("Enter id:")
	fmt.Scan(&id)
}

func readFromJSON(fileName string) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err.Error())
	}
	err = json.Unmarshal(data, &students)
	if err != nil {
		panic(err.Error())
	}
}

func writeToJSON(fileName string) {
	marshelledData, err := json.Marshal(students)
	if err != nil {
		panic(err.Error())
	}
	err = os.WriteFile(fileName, marshelledData, 0644)
	if err != nil {
		panic(err.Error())
	}
}

func addStudents(name string,engMarks,hindiMarks,mathsMarks,csMarks,sciMarks int) string{
	
	st = &Student{
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
	readFromJSON(fileName)
	students = append(students, st)
	writeToJSON(fileName)
	return addStudentMsg
}

func displayAllStudents() {
    readFromJSON(fileName)
    for _, student := range students {
        fmt.Printf("Name: %s, ID: %d, Marks: %+v\n", student.Name, student.Id, student.Marks)
    }
}


func displayStudentById(id int) *Student{
	readFromJSON(fileName)
	for _, val := range students {
		if val.Id == id {
			return val
		}
	}
	return nil
}

func updateStudent(id int, updName string) string{
	readFromJSON(fileName)
	for i, val := range students {
		if val.Id == id {
			students[i].Name = updName
		}
	}
	writeToJSON(fileName)
	return updateStudentMsg
}

func deleteStudent(id int) string {
	newStudents := []Student{}
	readFromJSON(fileName)
	for i := 0; i < len(students); i++ {
		if students[i].Id != id {
			newStudents = append(newStudents, *students[i])
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
	return deleteStudentMsg
}
