package main

import "fmt"

func Conditionals() {
    fmt.Println("---- IF CONDITION ----")

    a := 10
    b := 20

    // Simple if condition
    if a < b {
        fmt.Println("a is less than b")
    }

    // if-else condition
    if a > b {
        fmt.Println("a is greater than b")
    } else {
        fmt.Println("a is NOT greater than b")
    }

    // if-else if-else chain
    if a > b {
        fmt.Println("a is greater")
    } else if a == b {
        fmt.Println("a and b are equal")
    } else {
        fmt.Println("a is smaller")
    }

    // if with short statement
    if diff := b - a; diff > 5 {
        fmt.Println("Difference is greater than 5")
    } else {
        fmt.Println("Difference is 5 or less")
    }

    fmt.Println("\n---- SWITCH STATEMENTS ----")

    // Basic switch
    day := 3
    switch day {
    case 1:
        fmt.Println("Monday")
    case 2:
        fmt.Println("Tuesday")
    case 3:
        fmt.Println("Wednesday")
    default:
        fmt.Println("Another day")
    }

    // Multiple case values
    fruit := "apple"
    switch fruit {
    case "banana", "apple":
        fmt.Println("Fruit is either banana or apple")
    default:
        fmt.Println("Some other fruit")
    }

    // Switch without expression
    num := 15
    switch {
    case num < 0:
        fmt.Println("Negative number")
    case num == 0:
        fmt.Println("Zero")
    case num > 0 && num < 10:
        fmt.Println("Small positive number")
    default:
        fmt.Println("Large positive number")
    }

    // Switch with fallthrough
    x := 2
    switch x {
    case 1:
        fmt.Println("Case 1")
    case 2:
        fmt.Println("Case 2")
        fallthrough
    case 3:
        fmt.Println("Case 3 (executed due to fallthrough)")
    default:
        fmt.Println("Default case")
    }

    fmt.Println("---- END OF CONDITIONALS ----")
}
