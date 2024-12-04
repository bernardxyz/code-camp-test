package main

import (
	"code-camp-test/task-1"
	"fmt"
)

var cars = map[string]int{
	"audi":     100,
	"bmw":      200,
	"mercedes": 400,
	"toyota":   120,
}

func printMap(m map[string]int) map[string]int {
	for key, value := range m {
		if value >= 150 {
			fmt.Println("Proizvodi s cijenom vecom od 150", value)
		}
		fmt.Println(key, value)
	}

	return m
}

func main() {

	person := task_1.Employee{
		Name:     "Jon",
		Position: "Completed",
		Tasks:    "",
	}

	person2 := task_1.Employee{
		Name:     "Mark",
		Position: "In progress",
		Tasks:    "",
	}

	fmt.Println(person.AddTask("Running"), person2.AddTask("Driving"))

	fmt.Println(person.CompleteTasks(), person2.CompleteTasks())

	fmt.Println(printMap(cars))
}
