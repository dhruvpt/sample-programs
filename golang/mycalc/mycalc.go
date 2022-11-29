package main

import (
    "fmt"
    "os"
    "strconv"
)

func print_usage() {
    fmt.Println("Usage: ",os.Args[0]," <integer> <operator> <integer>")
    fmt.Println("       operator can be + for addition")
    fmt.Println("                       - for subtraction")
    fmt.Println("                       * for multiplication")
    fmt.Println("                       / for division")
}

func main() {
    if len(os.Args) != 4 {
        print_usage()
        return
    }
    firstArg, err1 := strconv.Atoi(os.Args[1])
    secondArg, err2 := strconv.Atoi(os.Args[3])
    if ((err1 != nil) || (err2 != nil)) {
        print_usage()
        return
    }
    switch (os.Args[2]) {
    case "+":
        fmt.Println("Result = ", firstArg + secondArg)
    case "-":
        fmt.Println("Result = ", firstArg - secondArg)
    case "/":
        fmt.Println("Result = ", float64(firstArg) / float64(secondArg))
    case "*":
        fmt.Println("Result = ", firstArg * secondArg)
    default:
        print_usage()
        return
    }

    return;
}
