package main

import (
    "flag"
    "fmt"
    "rational/pkg/rational"
)

var cnt int

func main() {
    flag.IntVar(&cnt, "count", 100, "repeat lcm this many times")
    flag.Parse()
    for i := 0; i < cnt; i++ {
        for j := 0; j < cnt; j++ {
            fmt.Printf("lcm(%d,%d)=%d\n", i, j, rational.LCM(i, j))
        }
    }
}
