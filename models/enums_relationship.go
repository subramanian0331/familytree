package models

import (
	"encoding/json"
	"fmt"
)

type Relationship int

const (
	ChildOf Relationship = iota
	Spouse
	Sibling
	Cousin
)

var (
	_RelationshipNameToValue = map[string]Relationship{
		"ChildOf": ChildOf,
		"Spouse":  Spouse,
		"Sibling": Sibling,
		"Cousin":  Cousin,
	}

	_RelationshipValueToName = map[Relationship]string{
		ChildOf: "ChildOf",
		Spouse:  "Spouse",
		Sibling: "Sibling",
		Cousin:  "Cousin",
	}
)

func init() {
	var v Relationship
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_RelationshipNameToValue = map[string]Relationship{
			interface{}(ChildOf).(fmt.Stringer).String(): ChildOf,
			interface{}(Spouse).(fmt.Stringer).String():  Spouse,
			interface{}(Sibling).(fmt.Stringer).String(): Sibling,
			interface{}(Cousin).(fmt.Stringer).String():  Cousin,
		}
	}
}

// MarshalJSON is generated so Relationship satisfies json.Marshaler.
func (r Relationship) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _RelationshipValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid Relationship: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so Relationship satisfies json.Unmarshaler.
func (r *Relationship) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Relationship should be a string, got %s", data)
	}
	v, ok := _RelationshipNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid Relationship %q", s)
	}
	*r = v
	return nil
}
