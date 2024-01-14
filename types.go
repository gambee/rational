package rational

import (
    "strconv"
)

type nan struct {}

func (n nan) Denominator() int {return 0}
func (n nan) Numerator() int {return 0}
func (n nan) String() string {return "NaN"}
func (n nan) Nan() bool {return true}


type integer int

func (i integer) Denominator() int {return 1}
func (i integer) Numerator() int {return int(i)}
func (i integer) String() string {return strconv.Itoa(int(i))}
func (i integer) Nan() bool {return false}


type quotient struct {
    num, den int
}

func (q *quotient) Denominator() int {
    return q.den
}

func (q *quotient) Numerator() int {
    return q.num
}

func (q *quotient) Nan() bool {
    return q.den == 0
}

func (q *quotient) String() string {
    switch {
        case q.Nan():
            return "NaN"
        case q.Numerator() == 0:
            return "0"
        case q.Denominator() == 1:
            return strconv.Itoa(q.Numerator())
        case q.Denominator() == -1:
            return strconv.Itoa(-1 * q.Numerator())
        case q.Denominator() < 0:
            return strconv.Itoa(-1 * q.Numerator()) + "/" + strconv.Itoa(-1 * q.Denominator())
        default:
            return strconv.Itoa(q.Numerator()) + "/" + strconv.Itoa(q.Denominator())
    }
}
