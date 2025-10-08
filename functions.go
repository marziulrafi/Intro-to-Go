package main

import "fmt"

func Functions() {
    fmt.Println("===== FUNCTIONS IN GOLANG =====")

    // 1️⃣ Basic function call
    sayHello()

    // 2️⃣ Function with parameters
    greet("Marziul", "Rafi")

    // 3️⃣ Function with return value
    sum := add(10, 20)
    fmt.Println("Sum:", sum)

    // 4️⃣ Function with multiple return values
    total, avg := calculate(10, 20, 30)
    fmt.Println("Total:", total, "Average:", avg)

    // 5️⃣ Named return values
    s, a := namedReturn(4, 5)
    fmt.Println("Named Return - Sum:", s, "Average:", a)

    // 6️⃣ Variadic function (takes any number of arguments)
    result := addAll(1, 2, 3, 4, 5)
    fmt.Println("Sum of all:", result)

    // 7️⃣ Anonymous function (inline function)
    func(x int) {
        fmt.Println("Anonymous function says:", x*x)
    }(5)

    // 8️⃣ Function as a variable
    square := func(n int) int {
        return n * n
    }
    fmt.Println("Square of 6:", square(6))

    // 9️⃣ Function returning another function (closure)
    nextNumber := counter()
    fmt.Println("Counter:", nextNumber())
    fmt.Println("Counter:", nextNumber())
    fmt.Println("Counter:", nextNumber())

    // 🔟 Recursive function
    fmt.Println("Factorial of 5:", factorial(5))

    fmt.Println("===== END OF FUNCTIONS =====")
}

/////////////////////////////////////////////////////////////
// Function Definitions Below
/////////////////////////////////////////////////////////////

// 1️⃣ Basic function (no parameters, no return)
func sayHello() {
    fmt.Println("Hello from Go Function!")
}

// 2️⃣ Function with parameters
func greet(firstName string, lastName string) {
    fmt.Println("Hello,", firstName, lastName)
}

// 3️⃣ Function with single return value
func add(a int, b int) int {
    return a + b
}

// 4️⃣ Function with multiple return values
func calculate(a, b, c int) (int, float64) {
    total := a + b + c
    avg := float64(total) / 3
    return total, avg
}

// 5️⃣ Named return values
func namedReturn(a, b int) (sum int, avg float64) {
    sum = a + b
    avg = float64(sum) / 2
    return // returns sum, avg automatically
}

// 6️⃣ Variadic function (accepts variable number of args)
func addAll(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// 9️⃣ Function returning another function (closure)
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

// 🔟 Recursive function
func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}
