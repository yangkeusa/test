package main

import "fmt"

func fib(x int) int {
        y := 1
        z := 1
        for i := 0; i < x; i++ {
                t := z
                z = z + y
                y = t
        }
        return z
}

func main() {
        fmt.Printf("test test \n")
        for i := 0; i < 10; i++ {
                x := fib(i)
                fmt.Printf("Fib(%d) = %d\n", i, x)
        }
}
