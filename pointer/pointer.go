package main

import "fmt"

func main() {
	var age, age2 int
	age = 34
	age2 = 35
	fmt.Println("Age is: ", age)
	fmt.Println("Age2 is: ", age2)
	var agePointer *int
	agePointer = &age
	fmt.Println("Age Pointer is: ", &agePointer)
	fmt.Println("Age Pointer value is: ", *agePointer)
	fmt.Println("Age Pointer : ", agePointer)
	fmt.Println("Age2 is: ", age2)
	age = 35
	fmt.Println("Age is: ", age)
	fmt.Println("Age Pointer is: ", &agePointer)
	fmt.Println("Age Pointer value is: ", *agePointer)
	fmt.Println("Age Pointer : ", agePointer)
	fmt.Println("Age2 is: ", age2)
	*agePointer = 42
	age2 = 43
	fmt.Println("Age is: ", age)
	fmt.Println("Age Pointer is: ", &agePointer)
	fmt.Println("Age Pointer value is: ", *agePointer)
	fmt.Println("Age Pointer : ", agePointer)
	fmt.Println("Age2 is: ", age2)
}
