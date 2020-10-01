package validate

import (
    "fmt"
    "reflect"
    "testing"
)

type Float64Case struct {
    value float64
    limiter Float64Limit
    expected *Float64LimitExceeded
}

func (f Float64Case) Test() (err error) {
    err = f.limiter.Validate(f.value)
    
    if f.expected == nil && err != nil {
        return fmt.Errorf("expected no error but got %w", err)
    }
    
    if f.expected != nil && !reflect.DeepEqual(err, *f.expected) {
        return fmt.Errorf(
            "expected identical errors but got different ones, %#v, %#v",
            err, f.expected)
    } else {
        return nil
    }
}

func testFloat64Limit(cases []Float64Case) (err error) {
    for ii, _ := range cases {
        if err = cases[ii].Test(); err != nil {
            break
        }
    }
    return
}

func Float64Exceeded(value float64, limiter Float64Limit) (f Float64Case) {
    exc := limiter.Exceeded(value)
    f.value, f.limiter, f.expected = value, limiter, &exc
    return
}

func TestFloat64Limit(t *testing.T) {
    cases := []Float64Case{
        Float64Case{
            value: 0.0,
            limiter: GreaterThan.Float64(-1.0),
            expected: nil,
        },
        Float64Case{
            value: -1.0,
            limiter: GreaterOrEqual.Float64(-1.0),
            expected: nil,
        },
        Float64Case{
            value: 100.0,
            limiter: GreaterOrEqual.Float64(-1.0),
            expected: nil,
        },
        Float64Exceeded(-1.0, GreaterThan.Float64(-1.0)),
        Float64Exceeded(-1.0, LessThan.Float64(-3.0)),
        Float64Exceeded(-1.0, LessOrEqual.Float64(-1.0)),
    }
    
    if err := testFloat64Limit(cases); err != nil {
        t.Fatalf("%s\n", err)
    }    
}
