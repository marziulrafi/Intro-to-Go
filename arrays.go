package main

import "fmt"

func Arrays() {
    fmt.Println("===== ARRAYS IN GOLANG =====")

    // Declare an array with fixed size
    var numbers [3]int
    numbers[0] = 10
    numbers[1] = 20
    numbers[2] = 30
    fmt.Println("Numbers:", numbers)

    // Short declaration with values
    fruits := [3]string{"Apple", "Banana", "Cherry"}
    fmt.Println("Fruits:", fruits)

    // Access specific element
    fmt.Println("First fruit:", fruits[0])

    // Change an element
    fruits[1] = "Mango"
    fmt.Println("After change:", fruits)

    // Array with auto length
    colors := [...]string{"Red", "Green", "Blue"}
    fmt.Println("Colors:", colors)

    // Iterate over an array using for loop
    fmt.Println("Numbers in array:")
    for i := 0; i < len(numbers); i++ {
        fmt.Println(numbers[i])
    }

    // Iterate using range (simpler way)
    fmt.Println("Fruits in array:")
    for index, fruit := range fruits {
        fmt.Println(index, "→", fruit)
    }

    fmt.Println("===== END OF ARRAYS =====")
}
