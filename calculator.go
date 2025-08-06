package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run main.go [add|sub|mul|div] num1 num2")
		return
	}

	op := os.Args[1]
	a, _ := strconv.Atoi(os.Args[2])
	b, _ := strconv.Atoi(os.Args[3])

	switch op {
	case "add":
		fmt.Println("Result:", a+b)
	case "sub":
		fmt.Println("Result:", a-b)
	case "mul":
		fmt.Println("Result:", a*b)
	case "div":
		if b == 0 {
			fmt.Println("Cannot divide by zero")
			return
		}
		fmt.Println("Result:", a/b)
	default:
		fmt.Println("Unknown operation:", op)
	}
}
