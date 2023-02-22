package golesson

import (
	"errors"
	"fmt"
)

const (
	STUDENT  string = "STUDENT"
	EMPLOYEE string = "EMPLOYEE"
	ROBOT    string = "ROBOT"
)

type ValidateStruct struct {
	IsOk  bool
	Error error
}

type People interface {
	SayHi() string
	Sing(lyrics string) string
	Validate() ValidateStruct
	Arrange() string
}

type Person struct {
	UUID string
	Name string
	Age  int
	Type string
}

type Robot struct {
	UUID    string
	Name    string
	Age     int
	Type    string
	Battery string
}

func (s Person) SayHi() string {
	return fmt.Sprintf("Hi My name is %s, I am %d years old, and I am a %s\n", s.Name, s.Age, s.Type)
}

func (r Robot) SayHi() string {
	return fmt.Sprintf("Hi My name is %s, I am %d years old, and I am a %s, battery %s\n", r.Name, r.Age, r.Type, r.Battery)
}

func (e Person) Sing(lyrics string) string {
	fmt.Printf("Hi My name is %s, and I am a %s, ======> %s\n", e.Name, e.Type, lyrics)
	return ""
}

func (r Robot) Sing(lyrics string) string {

	return "q12315t526"
}

func (p Person) Validate() ValidateStruct {
	if p.UUID == "" {
		return ValidateStruct{false, errors.New("uuid should be declared")}
	}

	if p.Name == "" {
		return ValidateStruct{false, errors.New("name should be declared")}
	}

	if p.Age <= 0 || p.Age >= 110 {
		return ValidateStruct{false, errors.New("age should large than 0 and less than 110")}
	}

	if p.Type == "" {
		return ValidateStruct{false, errors.New("type should be declared")}
	}
	return ValidateStruct{true, nil}
}

func (r Robot) Validate() ValidateStruct {
	if r.UUID == "" {
		return ValidateStruct{false, errors.New("uuid should be declared")}
	}

	if r.Name == "" {
		return ValidateStruct{false, errors.New("name should be declared")}
	}

	if r.Age <= 0 {
		return ValidateStruct{false, errors.New("age should large than 0")}
	}

	if r.Type == "" {
		return ValidateStruct{false, errors.New("type should be declared")}
	}
	return ValidateStruct{true, nil}
}

func (s Person) Arrange() string {
	return fmt.Sprintf("UUID: %s\n", s.UUID) + fmt.Sprintf("Name: %s\n", s.Name) + fmt.Sprintf("Age: %d\n", s.Age) + s.SayHi()
}

func (r Robot) Arrange() string {
	return fmt.Sprintf("UUID: %s\n", r.UUID) + fmt.Sprintf("Name: %s\n", r.Name) + fmt.Sprintf("Age: %d\n", r.Age) + r.SayHi()
}

func ArrangeOrder(people []People) (string, error) {
	result := ""
	for _, p := range people {
		if !p.Validate().IsOk {
			return "", p.Validate().Error
		}
		result += p.Arrange() + "\n----------\n"
	}

	return result, nil
}

func GetPersonInterface() {
	var student People = Person{UUID: "1", Name: "Sam", Age: 20, Type: STUDENT}

	var robot People = Robot{UUID: "2", Name: "Rampage", Age: 12, Type: ROBOT, Battery: "On"}

	result, err := ArrangeOrder([]People{robot, student})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
