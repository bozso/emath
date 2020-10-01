package validate

import (
    "encoding/json"
)

type Validator interface {
    Validate() error
}

func FromJson(v Validator, b []byte) (err error) {
    if err = json.Unmarshal(b, v); err != nil {
        return
    }
    
    return v.Validate()
}
