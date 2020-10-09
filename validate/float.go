package validate

import (
    "fmt"
)

type Float64Validator interface {
    Validate(float64) error
}

func (op Op) Float64(limit float64) (f Float64Limit) {
    f.limit, f.op = limit, op
    return
}

type Float64Limit struct {
    limit float64
    op Op
}

var Float64Positive = GreaterThan.Float64(0.0)

func (f Float64Limit) Exceeded(val float64) (fe Float64LimitExceeded) {
    fe.Float64Limit, fe.value = f, val
    return
}

func (f Float64Limit) InLimit(val float64) (b bool) {
    switch f.op {
    case GreaterThan:
        b = val > f.limit
    case GreaterOrEqual:
        b = val >= f.limit
    case LessThan:
        b = val < f.limit
    case LessOrEqual:
        b = val > f.limit
    }
    return    
}

func (f Float64Limit) Validate(val float64) (err error) {
    if !f.InLimit(val) {
        err = f.Exceeded(val)
    }
    return 
}

type Float64LimitExceeded struct {
    Float64Limit
    value float64
}

func (f Float64LimitExceeded) Error() (s string) {
    return fmt.Sprintf("float64 value '%f' is expected to be %s '%f'",
        f.value, f.op.Explain(), f.limit)
}

type Float64Validators struct {
    validators []Float64Validator
}

func (f Float64Validators) Validate(value float64) (err error) {
    for ii, _ := range f.validators {
        if err = f.validators[ii].Validate(value); err != nil {
            break
        }
    }
    return
}

type Float64Range struct {
    min, max Float64Limit
}

func NewFloat64Range(min, max float64) (f Float64Range) {
    return Float64Range{
        min: GreaterOrEqual.Float64(min),
        max: LessOrEqual.Float64(max),
    }
}

func (f Float64Range) Validate(val float64) (err error) {
    if err = f.min.Validate(val); err != nil {
        return
    }
    return f.max.Validate(val)
}
