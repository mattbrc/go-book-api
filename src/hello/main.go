package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
}

// vars

var name string = "John"

// var age := 30 // type inference

// funcs

func greet(name string) string {
	return "Hello " + name
}

// structs

type Person struct {
	name string
	age  int
}

// slices - dynamic arrays

var numbers []int = []int{1, 2, 3, 4, 5}

// maps - key-value pairs

var person = map[string]int{
	"John": 30,
	"Jane": 25,
}

// control structures

func checkNumber(x int) {
	if x > 0 {
		fmt.Println("x is positive")
	} else if x < 0 {
		fmt.Println("x is negative")
	} else {
		fmt.Println("x is zero")
	}
}
