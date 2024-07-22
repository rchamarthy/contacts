package contacts

import (
	"fmt"
)

// Date represents a date.
type Date struct {
	Month int
	Day   int
	Year  int
}

func (d Date) MarshalText() ([]byte, error) {
	s := fmt.Sprintf("%d-%02d-%02d", d.Month, d.Day, d.Year)

	return []byte(s), nil
}

func (d *Date) UnmarshalText(data []byte) error {
	_, err := fmt.Sscanf(string(data), "%d-%d-%d", &d.Month, &d.Day, &d.Year)

	return err
}

func (d Date) String() string {
	b, e := d.MarshalText()
	if e != nil {
		return ""
	}

	return string(b)
}

func (d Date) Validate() error {
	if d.Day > 31 || d.Day < 1 {
		return fmt.Errorf("invalid day: %d", d.Day)
	}

	if d.Month > 12 || d.Month < 1 {
		return fmt.Errorf("invalid month: %d", d.Month)
	}

	if d.Year < 1900 {
		return fmt.Errorf("year must be greater than 1900: %d", d.Year)
	}

	return nil
}
