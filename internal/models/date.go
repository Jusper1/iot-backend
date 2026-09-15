package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type Date string

func (d *Date) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*d = ""
		return nil
	}

	var value string

	if err := json.Unmarshal(data, &value); err != nil {
		return fmt.Errorf("tanggal harus berupa string dengan format YYYY-MM-DD")
	}

	if value == "" {
		*d = ""
		return nil
	}

	parsed, err := time.Parse("2006-01-02", value)
	if err != nil {
		return fmt.Errorf(
			"format tanggal tidak valid: %q, gunakan YYYY-MM-DD",
			value,
		)
	}

	formatted := parsed.Format("2006-01-02")

	if formatted != value {
		return fmt.Errorf(
			"format tanggal tidak valid: %q, gunakan YYYY-MM-DD",
			value,
		)
	}

	*d = Date(value)

	return nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	if d == "" {
		return []byte("null"), nil
	}

	return json.Marshal(string(d))
}

func (d Date) Value() (driver.Value, error) {
	if d == "" {
		return nil, nil
	}

	if _, err := time.Parse("2006-01-02", string(d)); err != nil {
		return nil, fmt.Errorf("tanggal tidak valid: %s", d)
	}

	return string(d), nil
}

func (d *Date) Scan(value interface{}) error {
	if value == nil {
		*d = ""
		return nil
	}

	if value == "" {
		*d = ""
		return nil
	}

	switch v := value.(type) {

	case time.Time:
		*d = Date(v.Format("2006-01-02"))
		return nil

	case []byte:
		date := string(v)

		if date == "" {
			*d = ""
			return nil
		}

		if _, err := time.Parse("2006-01-02", date); err != nil {
			return fmt.Errorf("tanggal database tidak valid: %s", date)
		}

		*d = Date(date)
		return nil

	case string:
		if v == "" {
			*d = ""
			return nil
		}

		if _, err := time.Parse("2006-01-02", v); err != nil {
			return fmt.Errorf("tanggal database tidak valid: %s", v)
		}

		*d = Date(v)
		return nil

	default:
		return fmt.Errorf("tidak dapat membaca tipe tanggal %T", value)
	}
}