package validate

import (

)

type Op int

const (
    GreaterThan Op = iota
    GreaterOrEqual
    LessThan
    LessOrEqual
)

func (op Op) Explain() (s string) {
    switch op {
    case GreaterThan:
        s = "greater than"
    case GreaterOrEqual:
        s = "greater or equal to"
    case LessThan:
        s = "less than"
    case LessOrEqual:
        s = "less or equal to"
    }
    return
}
