// Code generated by github.com/infraboard/mcube
// DO NOT EDIT

package mongodb

import (
	"bytes"
	"fmt"
	"strings"
)

// ParseSTATUSFromString Parse STATUS from string
func ParseSTATUSFromString(str string) (STATUS, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := STATUS_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown STATUS: %s", str)
	}

	return STATUS(v), nil
}

// Equal type compare
func (t STATUS) Equal(target STATUS) bool {
	return t == target
}

// IsIn todo
func (t STATUS) IsIn(targets ...STATUS) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t STATUS) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *STATUS) UnmarshalJSON(b []byte) error {
	ins, err := ParseSTATUSFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}
