package timef

import (
	"encoding/json"
	"fmt"
	"time"
)

// Time is a wrapper around time.Time.
type Time[F fmt.Stringer] struct {
	time.Time
}

// New returns a Time with the given time value and format.
func New[F fmt.Stringer](t time.Time) Time[F] {
	return Time[F]{Time: t}
}

// UnmarshalJSON unmarshal time from a string with its format.
func (t *Time[F]) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	var format F
	parsed, err := time.Parse(format.String(), s)
	if err != nil {
		return err
	}

	t.Time = parsed
	return nil
}

// MarshalJSON marshals the time to a string with its format.
func (t Time[F]) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.String() + `"`), nil
}

// String returns a textual representation of the time value formatted according
// to the layout defined by F.
func (t Time[F]) String() string {
	var format F
	return t.Format(format.String())
}

// Format returns a textual representation of the time value formatted according
// to the layout defined by F.
func Format[F fmt.Stringer](t time.Time) string {
	return Time[F]{t}.String()
}

// Parse parses a formatted string with format defined by F and returns the time
// value it represents.
func Parse[F fmt.Stringer](s string) (*Time[F], error) {
	var format F

	t, err := time.Parse(format.String(), s)
	if err != nil {
		return nil, err
	}

	return &Time[F]{t}, nil
}
