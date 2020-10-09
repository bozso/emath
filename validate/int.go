package validate

import (
    "fmt"
    "encoding/json"
)

type IntValidator interface {
    Validate(int) error
}

func IntFromJson(iv IntValidator, b []byte) (ii int, err error) {
    if err = json.Unmarshal(b, &ii); err != nil {
        return
    }
    
    err = iv.Validate(ii)
    return
}

func (op Op) Int(limit int) (i IntLimit) {
    i.limit, i.op = limit, op
    return
}

type IntLimit struct {
    limit int
    op Op
}

var IntPositive = GreaterThan.Int(0)

type PositiveInt int

func (p *PositiveInt) UnmarshalJSON(b []byte) (err error) {
    val, err := IntFromJson(IntPositive, b)
    if err != nil {
        return
    }
    *p = PositiveInt(val)
    return
}

var IntNatural = GreaterOrEqual.Int(0)

type NaturalInt int

func (n *NaturalInt) UnmarshalJSON(b []byte) (err error) {
    val, err := IntFromJson(IntNatural, b)
    if err != nil {
        return
    }
    
    *n = NaturalInt(val)
    return
}

func (i IntLimit) Exceeded(val int) (ie IntLimitExceeded) {
    ie.IntLimit, ie.value = i, val
    return
}

func (i IntLimit) InLimit(val int) (b bool) {
    switch i.op {
    case GreaterThan:
        b = val > i.limit
    case GreaterOrEqual:
        b = val >= i.limit
    case LessThan:
        b = val < i.limit
    case LessOrEqual:
        b = val > i.limit
    }
    return    
}

func (i IntLimit) Validate(val int) (err error) {
    if !i.InLimit(val) {
        err = i.Exceeded(val)
    }
    return 
}

type IntLimitExceeded struct {
    IntLimit
    value int
}

func (f IntLimitExceeded) Error() (s string) {
    return fmt.Sprintf("float64 value '%f' is expected to be %s '%f'",
        f.value, f.op.Explain(), f.limit)
}
