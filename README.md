# Go-Codes Repository

This repository contains example codes, exercises, and explanations for learning the Go programming language (Golang).

## What is Go?
Go is an open-source programming language developed by Google. It is designed for simplicity, reliability, and efficiency, making it ideal for building scalable and high-performance applications.

## Installation

### 1. Download and Install Go
- Visit the official Go website: [https://go.dev/dl/](https://go.dev/dl/)
- Download the installer for your operating system (Windows, macOS, Linux).
- Run the installer and follow the instructions.

### 2. Verify Installation
Open your terminal and run:
```sh
go version
```
You should see the installed Go version.

### 3. Set Up Workspace
- Create a directory for your Go projects (e.g., `Go-Codes`).
- Set the `GOPATH` environment variable if needed (modern Go uses modules, so this is optional).

## Go Syntax Overview

### Hello World
```go
package main
import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

### Variables
```go
var x int = 10
name := "GoLang"
```

### Functions
```go
func add(a int, b int) int {
    return a + b
}
```

### Arrays and Slices
```go
arr := [3]int{1, 2, 3}
slice := []int{4, 5, 6}
```

### Loops
```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

### Conditionals
```go
if x > 5 {
    fmt.Println("x is greater than 5")
} else {
    fmt.Println("x is 5 or less")
}
```

### Structs
```go
type Person struct {
    Name string
    Age  int
}
```

### Maps
```go
m := map[string]int{"a": 1, "b": 2}
```

### Packages
```go
import "fmt"
```

## How to Run Go Code
1. Save your code in a `.go` file (e.g., `main.go`).
2. Open terminal in the file's directory.
3. Run:
```sh
go run main.go
```

## Repository Structure
- `arrays/` - Examples related to arrays
- `functions/` - Function examples
- `hello-world/` - Basic Hello World program
- `slices/` - Slice examples
- `variables/` - Variable usage examples

## Contribution
Feel free to contribute by adding new examples, fixing issues, or improving documentation.

---
Happy Coding in Go!
