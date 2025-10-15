package main
import "fmt"



// 🌍 PACKAGE SCOPE
// Declared outside any function — accessible to all functions in this package.
var packageVar = "I'm a package-level variable!"

// Capitalized means "exported" → can be used from other packages if imported.
var Version = "1.0.0"

func ScopeTypes() {
    fmt.Println("===== SCOPE TYPES IN GOLANG =====")

    // 🧠 LOCAL SCOPE
    
    localVar := "I'm local to ScopeTypes()"
    fmt.Println(localVar)      // ✅ Accessible here
    showLocalExample()         // Calling another function
    // fmt.Println(localExample) ❌ Not accessible here (declared in another function)

    



    // 🧱 BLOCK SCOPE

    number := 10
    if number > 5 {
        msg := "I'm inside an if-block"
        fmt.Println("Inside block:", msg)
    }
    // fmt.Println(msg) ❌ ERROR: not accessible outside the if block

    for i := 1; i <= 3; i++ {
        fmt.Println("Inside loop:", i)
    }
    // fmt.Println(i) ❌ ERROR: i not accessible outside loop

    



    // 📦 PACKAGE SCOPE
    
    fmt.Println("Accessing packageVar:", packageVar) // ✅ Accessible here
    fmt.Println("App Version:", Version)
    modifyPackageVar()

    fmt.Println("===== END OF SCOPE TYPES =====")
}

// Demonstrates another local scope
func showLocalExample() {
    localExample := "I'm local to showLocalExample()"
    fmt.Println(localExample)
}

// Demonstrates modifying package-level variable
func modifyPackageVar() {
    packageVar = "Modified inside modifyPackageVar()"
    fmt.Println("Modified packageVar:", packageVar)
}
