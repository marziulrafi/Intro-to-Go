package main

import "fmt"

func Maps() {
	
	// Creating an empty map using make()
	studentMarks := make(map[string]int)

	// Adding key-value pairs
	studentMarks["Marziul"] = 95
	studentMarks["Rafi"] = 90
	studentMarks["Karim"] = 85

	fmt.Println("Student Marks:", studentMarks)

	// Accessing values using keys
	fmt.Println("\nMarziul's mark:", studentMarks["Marziul"])

	// Updating a value
	studentMarks["Rafi"] = 92
	fmt.Println("Updated Rafi's mark:", studentMarks["Rafi"])

	// Deleting a key-value pair
	delete(studentMarks, "Karim")
	fmt.Println("\nAfter deleting Karim:", studentMarks)

	// Checking if a key exists
	value, exists := studentMarks["Karim"]
	if exists {
		fmt.Println("Karim's mark:", value)
	} else {
		fmt.Println("Karim not found in map")
	}

	// Declaring and initializing directly
	countries := map[string]string{
		"BD": "Bangladesh",
		"US": "United States",
		"JP": "Japan",
	}

	fmt.Println("\nCountries Map:", countries)

	// Iterating over map
	fmt.Println("\nIterating over student marks:")
	for name, mark := range studentMarks {
		fmt.Println(name, "=>", mark)
	}

	// Length of map
	fmt.Println("\nNumber of students:", len(studentMarks))

	fmt.Println("===== END OF MAPS =====")
}
