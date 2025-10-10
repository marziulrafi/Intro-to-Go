package main

import "fmt"

func Slices() {
    fmt.Println("===== SLICES IN GOLANG =====")

    // Create a slice directly
    fruits := []string{"Apple", "Banana", "Cherry"}
    fmt.Println("Fruits:", fruits)

    // Add elements using append()
    fruits = append(fruits, "Mango", "Orange")
    fmt.Println("After append:", fruits)

    // Create a slice from an existing array
    numbers := [5]int{10, 20, 30, 40, 50}
    part := numbers[1:4] // includes index 1 to 3
    fmt.Println("Array:", numbers)
    fmt.Println("Sliced part:", part)

    // Changing the slice also changes the original array
    part[0] = 200
    fmt.Println("Modified slice:", part)
    fmt.Println("Original array after slice change:", numbers)

    // Length and Capacity
    fmt.Println("Length:", len(part))   // number of elements
    fmt.Println("Capacity:", cap(part)) // how many it can hold before new memory is allocated

    // Create an empty slice using make()
    scores := make([]int, 3) // length = 3, capacity = 3
    fmt.Println("Empty slice:", scores)

    scores[0] = 50
    scores[1] = 70
    scores[2] = 90
    fmt.Println("After adding values:", scores)

    // Append beyond capacity (automatically expands)
    scores = append(scores, 100)
    fmt.Println("After appending extra element:", scores)

    // Copy slices
    newScores := make([]int, len(scores))
    copy(newScores, scores)
    fmt.Println("Copied slice:", newScores)

    // Slice of slice
    sub := scores[1:3]
    fmt.Println("Sub-slice (index 1 to 2):", sub)

    // Iterate through a slice
    fmt.Println("Iterating through slice:")
    for index, value := range fruits {
        fmt.Println(index, "→", value)
    }

    fmt.Println("===== END OF SLICES =====")
}
