package models

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type Dob time.Time

func (d *Dob) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`) //get rid of "
	if value == "" || value == "null" {
		return nil
	}

	t, err := time.Parse("2006-01-02", value) //parse time
	if err != nil {
		return err
	}
	*d = Dob(t) //set result using the pointer
	return nil
}

type Member struct {
	Id         uuid.UUID `json:"id"`
	Firstname  string    `json:"firstname"`
	Lastname   string    `json:"lastname"`
	Sex        Sex       `json:"sex"`
	Dob        Dob       `json:"dob"`
	Occupation string    `json:"occupation"`
	Family     string    `json:"family"`
}

func (m *Member) Age() int {
	AgeDuration := time.Since(time.Time(m.Dob))
	return int(AgeDuration.Hours()) / 8760
}
