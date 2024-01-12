package rational

func Neg(n Number) Number {
    if n.Nan() {
        return NaN
    } else {
        return NewNumber(-1 * n.Numerator(), n.Denominator())
    }
}

func Inv(n Number) Number {
    if n.Nan() || n.Numerator() == 0 {
        return NaN
    } else {
        return NewNumber(n.Denominator(), n.Numerator())
    }
}

func Reduce(n Number) Number {
    return reduce(n.Numerator(), n.Denominator())
}

func Abs(n Number) Number {
    return nil
}

