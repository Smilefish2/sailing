package user

import (
	"database/sql/driver"
	"errors"
)

// Type is the 'type' enum type from schema 'sailing_user'.
type Type uint16

const (
	// TypeDefault is the 'default' Type.
	TypeDefault = Type(1)

	// TypeOfficial is the 'official' Type.
	TypeOfficial = Type(2)

	// TypeMerchant is the 'merchant' Type.
	TypeMerchant = Type(3)
)

// String returns the string value of the Type.
func (t Type) String() string {
	var enumVal string

	switch t {
	case TypeDefault:
		enumVal = "default"

	case TypeOfficial:
		enumVal = "official"

	case TypeMerchant:
		enumVal = "merchant"
	}

	return enumVal
}

// MarshalText marshals Type into text.
func (t Type) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}

// UnmarshalText unmarshals Type from text.
func (t *Type) UnmarshalText(text []byte) error {
	switch string(text) {
	case "default":
		*t = TypeDefault

	case "official":
		*t = TypeOfficial

	case "merchant":
		*t = TypeMerchant

	default:
		return errors.New("invalid Type")
	}

	return nil
}

// Value satisfies the sql/driver.Valuer interface for Type.
func (t Type) Value() (driver.Value, error) {
	return t.String(), nil
}

// Scan satisfies the database/sql.Scanner interface for Type.
func (t *Type) Scan(src interface{}) error {
	buf, ok := src.([]byte)
	if !ok {
		return errors.New("invalid Type")
	}

	return t.UnmarshalText(buf)
}
