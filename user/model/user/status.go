package user

import (
	"database/sql/driver"
	"errors"
)

// Status is the 'status' enum type from schema 'sailing_user'.
type Status uint16

const (
	// StatusNormal is the 'normal' Status.
	StatusNormal = Status(1)

	// StatusDisable is the 'disable' Status.
	StatusDisable = Status(2)

	// StatusDeleted is the 'deleted' Status.
	StatusDeleted = Status(3)
)

// String returns the string value of the Status.
func (s Status) String() string {
	var enumVal string

	switch s {
	case StatusNormal:
		enumVal = "normal"

	case StatusDisable:
		enumVal = "disable"

	case StatusDeleted:
		enumVal = "deleted"
	}

	return enumVal
}

// MarshalText marshals Status into text.
func (s Status) MarshalText() ([]byte, error) {
	return []byte(s.String()), nil
}

// UnmarshalText unmarshals Status from text.
func (s *Status) UnmarshalText(text []byte) error {
	switch string(text) {
	case "normal":
		*s = StatusNormal

	case "disable":
		*s = StatusDisable

	case "deleted":
		*s = StatusDeleted

	default:
		return errors.New("invalid Status")
	}

	return nil
}

// Value satisfies the sql/driver.Valuer interface for Status.
func (s Status) Value() (driver.Value, error) {
	return s.String(), nil
}

// Scan satisfies the database/sql.Scanner interface for Status.
func (s *Status) Scan(src interface{}) error {
	buf, ok := src.([]byte)
	if !ok {
		return errors.New("invalid Status")
	}

	return s.UnmarshalText(buf)
}
