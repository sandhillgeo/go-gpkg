package gpkg

import (
	"database/sql/driver"
	"strconv"
)

// NullInt8 represents an int8 that may be null.
// NullInt8 implements the Scanner interface so
// it can be used as a scan destination, similar to NullString.
type NullInt8 struct {
	Int8  int8
	Valid bool // Valid is true if Int8 is not NULL
}

// Scan implements the Scanner interface.
func (n *NullInt8) Scan(value interface{}) error {
	if value == nil {
		n.Int8, n.Valid = 0, false
		return nil
	}
	n.Valid = true
	//return convertAssign(&n.Int8, value)
	switch value.(type) {
	case string:
		i, err := strconv.Atoi(value.(string))
		if err != nil {
			n.Int8 = 0
			n.Valid = false
			return err
		}
		n.Int8 = int8(i)
	case int8:
		n.Int8 = value.(int8)
	case int32:
		n.Int8 = int8(value.(int32))
	case int64:
		n.Int8 = int8(value.(int64))
	case uint8:
		n.Int8 = int8(value.(uint8))
	case uint32:
		n.Int8 = int8(value.(uint32))
	case uint64:
		n.Int8 = int8(value.(uint64))
	}
	return nil
}

// Value implements the driver Valuer interface.
func (n NullInt8) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Int8, nil
}
