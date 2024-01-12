package rational

var NaN nan

const (
    Zero integer = 0
    One integer = 1
)

type Quotient interface {
    Denominator() int
    Numerator() int
}

type Number interface {
    Quotient
    String() string
    Nan() bool
}

type (
    Relation func(...Number) bool
    Operator func(...Number) Number
    Function func(Number) Number
    Mapping func(Function, []Number) []Number
)

func NewNumber(n, d int) (result Number) {
    return reduce(n, d)
}
