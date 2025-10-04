package main

import "fmt"

// global variable
var g int // default 0

// This is no longer 'main', it's a function you can call
func Variables() {
    // basic
    var a int = 10
    b := "GoLang"
    fmt.Println(a, b)

    // multiple
    x, y := 5, 6
    fmt.Println(x, y)

    // zero values
    var i int
    var s string
    fmt.Println(i, s) // 0 ""

    // float
    pi := 3.1415
    fmt.Println(pi)

    // boolean
    isCool := true
    fmt.Println(isCool)

    // complex
    c := 1 + 2i
    fmt.Println(real(c), imag(c))

    // array
    arr := [3]int{1, 2, 3}
    fmt.Println(arr)

    // slice
    nums := []int{10, 20, 30}
    fmt.Println(nums)

    // map
    person := map[string]string{"name": "Bob", "city": "Dhaka"}
    fmt.Println(person)

    // pointer
    p := &a
    fmt.Println(*p)

    // constant
    const Pi = 3.14
    fmt.Println(Pi)
}
