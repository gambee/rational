package rational

import (
    "strconv"
    "strings"
)

func GCD(is ...int) (ret int) {
    if len(is) == 0 {
        ret = 0
    } else {
        ret = is[0]
        for _, x := range is[1:] {
            ret = gcd(ret, x)
        }
    }
    return
}

func gcd(a, b int) int {
    if a > b {
        return gcdRecursive(a, b)
    } else {
        return gcdRecursive(b, a)
    }
}

func gcdRecursive(g, l int) int {
    if l == 0 {
        return g
    } else {
        return gcdRecursive(l, g % l)
    }
}

func abs(a int) int {
    if a < 0 {
        return -a
    } else {
        return a
    }
}

func LCM(is ...int) (ret int) {
    if len(is) == 0 {
        ret = -1
    } else {
        ret = is[0]
        for _, x := range is[1:] {
            ret = lcm(ret, x)
        }
    }
    return
}

func lcm(a, b int) int {
    if a == 0 && b == 0 {
        return 0
    } else {
        return abs(a * b) / gcd(a, b)
    }
}

func anyNans(numbers ...Number) (nan bool) {
    for _, n := range numbers {
        if nan = n.Nan(); nan {
            break
        }
    }
    return
}

func anyZeroes(numbers ...Number) (zero bool) {
    for _, n := range numbers {
        if zero = !n.Nan() && n.Numerator() == 0; zero {
            break
        }
    }
    return
}

func anyNansOrZeroes(numbers ...Number) (nan, zero bool) {
    for i, n := range numbers {
        if nan = n.Nan(); nan {
            zero = anyZeroes(numbers[i+1:]...)
            break
        }
        if zero = n.Numerator() == 0; zero {
            nan = anyNans(numbers[i+1:]...)
            break
        }
    }
    return
}

func reduce(num, den int) (result Number) {
    if den == 0 {
        return nan{}
    }

    if den < 0 {
        num *= -1
        den *= -1
    }

    if num == 0 {
        return integer(0)
    }

    g := GCD(num, den)
    num /= g
    den /= g

    if den == 1 {
        return integer(num)
    }

    return &quotient{num, den}
}


type commonized struct {
    den int
    nums []int
}

func (c *commonized) String() string {
    if c.den == -1 || c.den == 0 {
        return "[NaNs]"
    } else {
        ss := make([]string, len(c.nums))
        ds := "/" + strconv.Itoa(c.den)
        for i, n := range c.nums {
            ss[i] = strconv.Itoa(n) + ds
        }
        return "[" + strings.Join(ss, " ") + "]"
    }
}

func commonize(numbers ...Number) (c *commonized) {
    c = &commonized{}
    if c.den = LCD(numbers...); c.den != -1 {
        c.nums = make([]int, len(numbers))
        for i, n := range numbers {
            c.nums[i] = (c.den / n.Denominator()) * n.Numerator()
        }
    }
    return
}
