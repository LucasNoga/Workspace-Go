package main

import "fmt"
import "time"

var s [61]int

func fib(n int) int {
        if n < 3 {
                return 1
        }
        if s[n] != 0 {
                return s[n]
        }
        s[n] = fib(n-1) + fib(n-2)
        return s[n]

}

func main() {
        var i int
        t0 := time.Now()
        for i = 1; i <= 60; i++ {
                fmt.Printf("fib(%d) = %-15d\t", i, fib(i))
        }
        println()
        println("Durée :", time.Since(t0).Seconds())
}