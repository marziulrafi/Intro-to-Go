package main

import "fmt"

func Loops() {
    fmt.Println("---- FOR LOOPS ----")

    // Classic for loop
    for i := 0; i < 3; i++ {
        fmt.Println("Classic for loop:", i)
    }

    // While-style for loop
    j := 0
    for j < 3 {
        fmt.Println("While-style loop:", j)
        j++
    }

    // Infinite loop with break
    count := 0
    for {
        fmt.Println("Infinite loop iteration:", count)
        count++
        if count == 2 {
            break
        }
    }

    // For with continue
    fmt.Println("\n---- FOR WITH CONTINUE ----")
    for k := 0; k < 5; k++ {
        if k%2 == 0 {
            continue
        }
        fmt.Println("Odd number:", k)
    }

    // For-range over slice
    fmt.Println("\n---- FOR-RANGE OVER SLICE ----")
    nums := []int{10, 20, 30}
    for i, val := range nums {
        fmt.Println("Index:", i, "Value:", val)
    }

    // For-range over map
    fmt.Println("\n---- FOR-RANGE OVER MAP ----")
    person := map[string]string{"name": "Alice", "city": "Dhaka"}
    for key, value := range person {
        fmt.Println(key, "=", value)
    }

    // For-range over string
    fmt.Println("\n---- FOR-RANGE OVER STRING ----")
    text := "GoLang"
    for index, char := range text {
        fmt.Printf("Index: %d, Character: %c\n", index, char)
    }

    fmt.Println("---- END OF LOOPS ----")
}
