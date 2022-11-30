package main

import (
    "fmt"
    "os"
    "strconv"
)

func print_usage() {
    fmt.Println("Usage: ", os.Args[0], "<integer> <operator> <integer>")
    fmt.Println(" where operator can be + for addition")
    fmt.Println("                       - for subtraction")
    fmt.Println("                       * for multiplication")
    fmt.Println("                       / for division")
}

func main() {
    if len(os.Args) != 4 {
        fmt.Println("Error: Unexpected number of parameters")
        print_usage()
        return
    }
    firstArg, err1 := strconv.Atoi(os.Args[1])
    secondArg, err2 := strconv.Atoi(os.Args[3])
    if ((err1 != nil) || (err2 != nil)) {
        fmt.Println("Error: Operands must be integer types")
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
        fmt.Println("Error: Unexpected operator", os.Args[2])
        print_usage()
        return
    }

    return
}
