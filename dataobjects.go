package gonhl

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Type alias for yyyy-mm-dd formatting
type JsonBirthDate time.Time

type Height struct {
	Feet int
	Inches int
}

func (j *JsonBirthDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := CreateDateFromString(s)
	if err != nil {
		return err
	}
	*j = JsonBirthDate(t)
	return nil
}

func (j *JsonBirthDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(j)
}

// Format function for printing.
func (j *JsonBirthDate) Format(s string) string {
	t := time.Time(*j)
	return t.Format(s)
}

func (h *Height) UnmarshalJSON(b []byte) error {
	s := strings.Split(string(b), "' ")
	f, err := strconv.Atoi(strings.Trim(s[0], "\""))
	if err != nil {
		return err
	}
	i, err := strconv.Atoi(strings.Trim(s[1], "\\\""))
	if err != nil {
		return err
	}
	*h = Height {
		Feet: f,
		Inches: i,
	}
	return nil
}

func (h *Height) MarshalJSON() ([]byte, error) {
	return json.Marshal(h)
}

func (h *Height) Format() string {
	return fmt.Sprintf("%d' %d\"", h.Feet, h.Inches)
}