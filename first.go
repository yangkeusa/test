package main

import (
"fmt"
"github.com/yangkeusa/test/palindrome"
)

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
	s := "今天aha天今"
	b := palindrome.IsPalinDrome(s)
	fmt.Printf("Palin (%s) = %b\n", s, b)

	s2 := "aanb"
	b2 := palindrome.IsPalinDrome(s2)
        fmt.Printf("Palin (%s) = %b\n", s2, b2)
}
