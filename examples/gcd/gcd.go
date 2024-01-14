package main

import (
    "flag"
    "fmt"
    "github.com/gambee/rational"
)

var cnt int

func main() {
    flag.IntVar(&cnt, "count", 100, "repeat gcd this many times")
    flag.Parse()
    for i := 0; i < cnt; i++ {
        for j := 0; j < cnt; j++ {
            fmt.Printf("gcd(%d,%d)=%d\n", i, j, rational.GCD(i, j))
        }
    }
}
