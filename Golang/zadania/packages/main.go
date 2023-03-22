package main

import (
	"github.com/google/uuid"
	"github.com/grupawp/appdispatcher"
	"log"
)

type Student struct {
	FirstName     string
	LastName      string
	applicationID uuid.UUID
}

func (s Student) FullName() string {
	return s.FirstName + " " + s.LastName
}

func (s Student) ApplicationID() string {
	return s.applicationID.String()
}

func main() {
	code, err := appdispatcher.Submit(Student{
		FirstName:     "Jędrzej",
		LastName:      "Romankiewicz",
		applicationID: uuid.New(),
	})
	log.Println(code)
	log.Println(err)
}
