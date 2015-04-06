package main

import (
	"bytes"
	"fmt"
	"log"
"github.com/yangkeusa/test/palindrome"
"github.com/yangkeusa/test/exp"
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
	fmt.Printf("Palin (%s) = %v\n", s, b)

	s2 := "aanb"
	b2 := palindrome.IsPalinDrome(s2)
        fmt.Printf("Palin (%s) = %v\n", s2, b2)

	var buf bytes.Buffer
	logger := log.New(&buf, "logger: ", log.Lshortfile)
	exp.C(logger, 10)
	fmt.Print(&buf)

}
