package main

import (
	"bytes"
	"fmt"
	"log"
	"time"
	"github.com/yangkeusa/test/palindrome"
	"github.com/yangkeusa/test/exp"
	"github.com/yangkeusa/test/math"
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

func init() {
	fmt.Print("INITIALIZING....\n")
}

func MatMulBench(iter, m, k, n int) {
	a := math.NewMatrix(m, k)
	b := math.NewMatrix(k, n)
	for i := 0; i < m; i++ {
                for j := 0; j < k; j++ {
                        a.Set(i, j, float32(i + j))
                }
        }
        for i := 0; i < k; i++ {
                for j := 0; j < k; j++ {
                        b.Set(i, j, float32(i + j))
                }
        }
	begin := time.Now()
	for i := 0; i < iter; i++ {
		c, _ := math.Mult(*a, *b)
		_ = c
	}
	end := time.Now()
	dur := end.Sub(begin)

	num_ops := iter * m * k * n * 2
	tflops := float64(num_ops) / float64(dur.Nanoseconds())
	fmt.Printf("Took %v to finish %v operations = %v TFLOPS\n", dur, num_ops, tflops)
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


	const (
		_ = iota
		KB int = 1 << (10 * iota)
		MB
		GB
	)
	
	fmt.Printf("KB = %v, MB = %v, GB = %v\n", KB, MB, GB)

	MatMulBench(10, 128, 256, 512)
}
