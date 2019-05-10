package main

import "fmt"

func main() {  
    var age int // variable declaration
    fmt.Println("my age is ", age)
    age = 29 //assignment
    fmt.Println("my age is", age)
    age = 54 //assignment
	fmt.Println("my new age is", age)
	
	// declare with inital value
	var defaultAge int = 29 // variable declaration with initial value

	fmt.Println("my age is", defaultAge)

    name, age := "naveen", 29 //short hand declaration

    fmt.Println("my name is", name, "age is", age)
}