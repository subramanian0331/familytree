package models

import (
	"encoding/json"
	"fmt"
)

type Sex int

const (
	Male Sex = iota
	Female
	Nonbinary
)

var (
	_SexNameToValue = map[string]Sex{
		"Male":      Male,
		"Female":    Female,
		"Nonbinary": Nonbinary,
	}

	_SexValueToName = map[Sex]string{
		Male:      "Male",
		Female:    "Female",
		Nonbinary: "Nonbinary",
	}
)

func init() {
	var v Sex
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_SexNameToValue = map[string]Sex{
			interface{}(Male).(fmt.Stringer).String():      Male,
			interface{}(Female).(fmt.Stringer).String():    Female,
			interface{}(Nonbinary).(fmt.Stringer).String(): Nonbinary,
		}
	}
}

// MarshalJSON is generated so Sex satisfies json.Marshaler.
func (r Sex) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _SexValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid Sex: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so Sex satisfies json.Unmarshaler.
func (r *Sex) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Sex should be a string, got %s", data)
	}
	v, ok := _SexNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid Sex %q", s)
	}
	*r = v
	return nil
}

func (r *Sex) String() string {
	return [...]string{"male", "female", "nonbinary"}[*r]
}
