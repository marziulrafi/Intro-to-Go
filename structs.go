package main

import "fmt"

// Define a struct named Student
type Student struct {
	name string
	age  int
	roll int
}


func Structs() {

	// Initialize struct using struct literal
	student1 := Student{name: "Marziul", age: 23, roll: 101}
	fmt.Println("Student 1:", student1)

	// Assign values to fields individually
	var student2 Student
	student2.name = "Rafi"
	student2.age = 22
	student2.roll = 102
	fmt.Println("Student 2:", student2)

	// Access individual fields
	fmt.Println("\nName of student1:", student1.name)
	fmt.Println("Age of student1:", student1.age)

	// Pointer to struct (Go automatically dereferences)
	ptr := &student1
	fmt.Println("\nPointer accessing student1 name:", ptr.name)

	// Modify value using pointer
	ptr.age = 25
	fmt.Println("Updated age of student1:", student1.age)

	// Pass struct to a function
	printStudent(student2)

	// Pass pointer to function to modify struct
	updateAge(&student2, 30)
	fmt.Println("\nAfter updating age of student2:", student2)

	fmt.Println("===== END OF STRUCTS =====")
}


// Function that receives struct as a parameter (copy)
func printStudent(s Student) {
	fmt.Println("\nStudent details from function:", s)
}

// Function that receives pointer to struct (modifies original)
func updateAge(s *Student, newAge int) {
	s.age = newAge
}