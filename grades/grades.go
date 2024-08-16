package grades

import (
	"fmt"
	"sync"
)

type Grade struct {
	Title string
	Type  GradeType
	Score float32
}

type GradeType string

const (
	GradeTest     = GradeType("Test")
	GradeHomework = GradeType("Homework")
	GradeQuiz     = GradeType("Quiz")
)

type Student struct {
	ID        int
	FirstName string
	LastName  string
	Grades    []Grade
}

func (s Student) Average() float32 {
	var res float32
	for _, grade := range s.Grades {
		res += grade.Score
	}
	return res / float32(len(s.Grades))
}

type Students []Student

func (s Students) GetById(id int) (*Student, error) {
	for i := range s {
		if s[i].ID == id {
			return &s[i], nil
		}
	}
	return nil, fmt.Errorf("Student with id %v not found", id)
}

var (
	students      Students
	studentsMutex sync.Mutex
)
