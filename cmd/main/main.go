package main

import (
    "fmt"
    "rational/pkg/rational"
)

func CheckLCM(ints ...int) {
    lcm := rational.LCM(ints...)
    var fracs []rational.Number
    for _, v := range ints {
        fracs = append(fracs, rational.NewNumber(lcm, v))
    }
    fmt.Println(lcm, fracs)
}

func GuessLCM(guess int, ints ...int) {
    lcm := rational.LCM(ints...)
    var fracs []rational.Number
    for _, v := range ints {
        fracs = append(fracs, rational.NewNumber(guess, v))
    }
    fmt.Println(guess, fracs, lcm)
}

func CheckLCD(ints ...int) {
    nums := make([]rational.Number, len(ints))
    for i, v := range ints {
        nums[i] = rational.NewNumber(1, v)
    }
    fmt.Println(rational.Add(nums...), rational.LCD(nums...), rational.LCM(ints...))
}

func main() {
    vals := []int{
        1, 2,
        1, 3,
        4, 6,
        15, 13,
        5, 7,
        7, 5,
        //0, 11,
        //0, 4,
        2, 1,
        4, 2,
        6, 2,
        //3, 0,
    }

    fmt.Println(rational.GCD(vals...), rational.LCM(vals...))
    CheckLCM(vals...)
    CheckLCM(2, 3, 5, 7, 13, 15)
    GuessLCM(2730, vals...)

    CheckLCD(2, 3, 5, 7, 13)

    n1 := rational.NewNumber(1, 2)
    n2 := rational.NewNumber(9, 16)
    fmt.Println(rational.Add(n1, n2), rational.Sub(n1, n2), rational.Mult(n1, n2))
    fmt.Println(n1, rational.Inv(n1))
    fmt.Println(n2, rational.Inv(n2))

    nums := []rational.Number{n1, n2, rational.NewNumber(-4, 17)}
    fmt.Println(nums)
    f := func(n rational.Number) rational.Number {
        return rational.Add(n, rational.NewNumber(1, n.Denominator()))
    }
    fmt.Println(rational.Map(f, nums))
    fmt.Println(rational.Map(rational.Abs, nums))
}
