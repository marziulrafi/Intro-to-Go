package main

import "fmt"

// Global variable — accessible from any function in this package
var globalVar = "I'm a global variable!"

// Function to demonstrate scope concepts
func Scopes() {
    fmt.Println("===== SCOPE IN GOLANG =====")

    // Accessing global variable
    fmt.Println("Accessing Global Variable:", globalVar)

    // Local variable (only accessible inside this function)
    localVar := "I'm a local variable"
    fmt.Println("Accessing Local Variable:", localVar)

    // Block scope — variable inside if-block is limited to that block
    x := 10
    if x > 5 {
        y := 20 // y is only accessible inside this if block
        fmt.Println("Inside if block:", x, y)
    }
    // fmt.Println(y) ❌ Uncommenting this will cause error (undefined: y)

    // Variable shadowing — local variable hides global one
    globalVar := "I'm a local variable shadowing the global one!"
    fmt.Println("Shadowing Example:", globalVar)

    // Function scope example
    showFunctionScope()

    fmt.Println("===== END OF SCOPE EXAMPLES =====")
}

// Another function (separate scope)
func showFunctionScope() {
    // This function cannot access localVar from ScopeExample()
    fmt.Println("Inside another function:", globalVar)
    // fmt.Println(localVar) ❌ Not accessible — different function scope
}
