package rational

func Mult(numbers ...Number) Number {
    switch n, z := anyNansOrZeroes(numbers...); {
        case n, len(numbers) == 0:
            return NaN
        case z:
            return Zero
        default:
            num, den := numbers[0].Numerator(), numbers[0].Denominator()
            for _, nxt := range numbers[1:] {
                num *= nxt.Numerator()
                den *= nxt.Denominator()
            }
            return reduce(num, den)
    }
}

func Div(numbers ...Number) Number {
    if len(numbers) == 0 {
        return NaN
    } else {
        args := make([]Number, 1, len(numbers))
        args[0] = numbers[0]
        return Mult(append(args, Map(Inv, numbers[1:])...)...)
    }
}

func Equal(numbers ...Number) bool {
    return false
}

func LCD(numbers ...Number) int {
    switch {
        case len(numbers) == 0, anyNans(numbers...):
            return -1
        case len(numbers) == 1:
            return Reduce(numbers[0]).Denominator()
        default:
            dens := make([]int, len(numbers))
            for i, n := range numbers {
                dens[i] = Reduce(n).Denominator()
            }
            return LCM(dens...)
    }
}

func Add(numbers ...Number) Number {
    c := commonize(numbers...)
    if c.den == -1 {
        return NaN
    } else {
        var sum int
        for _, n := range c.nums {
            sum += n
        }
        return NewNumber(sum, c.den)
    }
}

func Sub(numbers ...Number) Number {
    c := commonize(numbers...)
    if c.den == -1 {
        return NaN
    } else {
        dif := c.nums[0]
        for _, n := range c.nums[1:] {
            dif -= n
        }
        return NewNumber(dif, c.den)
    }
}
