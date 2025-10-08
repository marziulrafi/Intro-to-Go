package main

import "fmt"

func DataTypes() {
    fmt.Println("===== DATA TYPES IN GOLANG =====")

    // -------------------------------
    // 1️⃣ BASIC NUMERIC TYPES
    // -------------------------------

    // Integer types
    var a int = 42              // default int (platform-dependent: 32 or 64 bit)
    var b int8 = -128            // 8-bit signed integer (-128 to 127)
    var c uint16 = 65535         // 16-bit unsigned integer (0 to 65535)
    fmt.Println("Integers:", a, b, c)

    // Floating-point types
    var f1 float32 = 3.14
    var f2 float64 = 2.718281828459045
    fmt.Println("Floats:", f1, f2)

    // Complex numbers
    var complexNum complex64 = 1 + 2i
    fmt.Println("Complex:", complexNum)
    fmt.Println("Real part:", real(complexNum))
    fmt.Println("Imag part:", imag(complexNum))

    // -------------------------------
    // 2️⃣ BOOLEAN TYPE
    // -------------------------------
    var isTrue bool = true
    var isFalse bool = false
    fmt.Println("Booleans:", isTrue, isFalse)

    // Boolean expressions
    fmt.Println("Is 10 > 5?", 10 > 5)
    fmt.Println("Is 3 == 7?", 3 == 7)

    // -------------------------------
    // 3️⃣ STRING TYPE
    // -------------------------------
    var name string = "Marziul"
    message := "Learning Go is fun!"
    fmt.Println("Strings:", name, "-", message)

    // String concatenation
    fullMsg := name + " says: " + message
    fmt.Println("Concatenated:", fullMsg)

    // Multi-line string
    multiLine := `
This is a 
multi-line 
string in Go.`
    fmt.Println("Multiline:", multiLine)

    // Strings are immutable
    // name[0] = 'm' // ❌ Not allowed

    // -------------------------------
    // 4️⃣ BYTE & RUNE
    // -------------------------------
    var ch byte = 'A' // byte = uint8, used for ASCII characters
    var r rune = '🔥' // rune = int32, used for Unicode characters
    fmt.Println("Byte:", ch, "->", string(ch))
    fmt.Println("Rune:", r, "->", string(r))

    // -------------------------------
    // 5️⃣ ARRAYS
    // -------------------------------
    var arr [3]int = [3]int{1, 2, 3}
    fmt.Println("Array:", arr)

    // Access and modify
    arr[1] = 5
    fmt.Println("Modified array:", arr)
    fmt.Println("Array length:", len(arr))

    // -------------------------------
    // 6️⃣ SLICES (dynamic arrays)
    // -------------------------------
    nums := []int{10, 20, 30}
    fmt.Println("Slice:", nums)
    nums = append(nums, 40, 50)
    fmt.Println("Appended slice:", nums)

    // Slice operations
    fmt.Println("Slice[1:3]:", nums[1:3])

    // -------------------------------
    // 7️⃣ MAPS (key-value pairs)
    // -------------------------------
    person := map[string]string{
        "name": "Rafi",
        "city": "Dhaka",
    }
    fmt.Println("Map:", person)
    fmt.Println("Access by key:", person["name"])

    // Add and delete
    person["country"] = "Bangladesh"
    delete(person, "city")
    fmt.Println("Modified map:", person)

    // -------------------------------
    // 8️⃣ STRUCTS (custom data types)
    // -------------------------------
    type Student struct {
        name string
        age  int
    }

    s1 := Student{name: "John", age: 22}
    fmt.Println("Struct:", s1)
    fmt.Println("Name:", s1.name, "Age:", s1.age)

    // -------------------------------
    // 9️⃣ INTERFACE{} (any type)
    // -------------------------------
    var anyType interface{}
    anyType = 10
    fmt.Println("Interface value:", anyType)
    anyType = "GoLang"
    fmt.Println("Interface changed to:", anyType)

    fmt.Println("===== END OF DATA TYPES =====")
}
