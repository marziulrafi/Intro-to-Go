package main

import "fmt"


func Pointers() {
	// Declare a normal variable
	var num int = 10
	fmt.Println("Value of num:", num)
	fmt.Println("Address of num (&num):", &num) // &num gives the memory address of num

	// Declare a pointer variable that stores the address of num
	var ptr *int = &num
	fmt.Println("\nPointer (ptr) stores:", ptr)    // address of num
	fmt.Println("Value at that address (*ptr):", *ptr) // *ptr gives the value at the memory address

	// Change the value using the pointer
	*ptr = 20 // modifies num directly through its address
	fmt.Println("\nAfter changing *ptr = 20:")
	fmt.Println("Value of num:", num)
	fmt.Println("Value through pointer (*ptr):", *ptr)

	// Pointer with new() function (creates a pointer to a zero value)
	newPtr := new(int)
	fmt.Println("\nnewPtr initially points to:", *newPtr) // 0 (default int value)
	*newPtr = 50
	fmt.Println("After assigning *newPtr = 50, value is:", *newPtr)

	// Nil pointer example
	var nilPtr *int
	if nilPtr == nil {
		fmt.Println("\nPointer is nil (doesn't point anywhere yet).")
	}

	// Function with pointer parameter
	changeValue(ptr)
	fmt.Println("\nAfter calling changeValue(ptr):", num)
}



// Function that takes a pointer and modifies its value
func changeValue(x *int) {
	*x = 100 // directly changes the value in memory
}
