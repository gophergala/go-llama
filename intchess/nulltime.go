// this is used to give the ability to NULL times in the database
// source: https://github.com/jinzhu/gorm/issues/10

package intchess

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

const (
	HTMLFormTime = "15:04"
	HTMLFormDate = "2006-01-02"
)

type NullTime struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL
}

// Scan implements the Scanner interface.
// Can scan from both time.Time and NullTime interfaces
func (nt *NullTime) Scan(value interface{}) error {
	if value == nil {
		nt.Valid = false
		return nil
	}

	switch iface := value.(type) {
	case time.Time:
		nt.Time, nt.Valid = iface, true
		return nil
	case NullTime:
		nt.Time, nt.Valid = iface.Time, iface.Valid
		return nil
	default:
		return nil
	}
	return nil
}

// Value implements the driver Valuer interface.
func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}

//this function is used when JSON tries to marshal this struct.
//Without it it will be marshalled into a JSON object with fields Time and Valid.
//This is far more elegant.
func (nt NullTime) MarshalJSON() ([]byte, error) {
	if nt.Valid {
		return nt.Time.MarshalJSON()
	} else {
		return json.Marshal(nil)
	}
}

func (nt *NullTime) UnmarshalJSON(b []byte) error {
	if string(b) == "null" || string(b) == "" {
		nt.Valid = false
		return nil
	}
	v, err := time.Parse(time.RFC3339, string(b[1:len(b)-1]))
	if err != nil {
		return err
	}
	nt.Time = v
	nt.Valid = true
	return nil
}

//This method is used to return the value for an HTML form default value field.
//The provided parameter specifies if it should be in Date or Time format.
func (nt NullTime) DefaultValue(param string) string {
	if nt.Valid == false {
		return ""
	}
	if param == "time" {
		return nt.Time.Format(HTMLFormTime)
	}
	if param == "date" {
		return nt.Time.Format(HTMLFormDate)
	}
	return nt.Time.Format(time.RFC3339)
}

//This is useful for getting the date from a NullTime in string format. It is intended for use in the HTML templates.
func (nt NullTime) GetDate() string {
	return nt.DefaultValue("date")
}

//This function is used to convert an HTML form value (returned from a create or edit, for instance) to a NullTime.
//It will first try parse it as a Time, if that does not work it will parse it as a Date.
func NullTimeConverter(b string) reflect.Value {
	decodedTime := NullTime{}
	//first check if we have been given a time
	v, err := time.Parse(HTMLFormTime, b)
	if err == nil {
		v := v.AddDate(1, 0, 0) //Nasty hack so that the year is not 0000, which is valid in Go but not MsSQL
		decodedTime.Time = v
		decodedTime.Valid = true
		return reflect.ValueOf(decodedTime)
	}
	//now check if it was a date
	v, err = time.Parse(HTMLFormDate, b)
	if err == nil {
		decodedTime.Time = v
		decodedTime.Valid = true
		return reflect.ValueOf(decodedTime)
	}
	return reflect.ValueOf(decodedTime)
}

type NullString struct {
	String string
	Valid  bool // Valid is true if Time is not NULL
}

// Scan implements the Scanner interface.
func (nt *NullString) Scan(value interface{}) error {
	if value == nil {
		nt.Valid = false
		return nil
	}
	nt.String, nt.Valid = value.(string), true
	return nil
}

// Value implements the driver Valuer interface.
func (nt NullString) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.String, nil
}

//this function is used when JSON tries to marshal this struct.
//Without it it will be marshalled into a JSON object with fields Time and Valid.
//This is far more elegant.
func (nt NullString) MarshalJSON() ([]byte, error) {
	if nt.Valid {
		return []byte(fmt.Sprintf("\"%v\"", nt.String)), nil
	} else {
		return json.Marshal(nil)
	}
}

func (nt *NullString) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		nt.Valid = false
		return nil
	}
	nt.String = string(b[1 : len(b)-1])
	nt.Valid = true
	return nil
}

func ConvertBool(value string) reflect.Value {
	if value == "on" {
		return reflect.ValueOf(true)
	} else if v, err := strconv.ParseBool(value); err == nil {
		return reflect.ValueOf(v)
	}

	return reflect.ValueOf(false)
}

func (nt NullString) DefaultValue(param string) string {
	if nt.Valid == false {
		return ""
	}
	return nt.String
}

//This function is used to convert an HTML form value (returned from a create or edit, for instance) to a NullTime.
//It will first try parse it as a Time, if that does not work it will parse it as a Date.
func NullStringConverter(s string) reflect.Value {
	decodedString := NullString{}
	//If string is empty we'll make it null
	if s == "" {
		decodedString.Valid = false
	} else {
		decodedString.String = s
		decodedString.Valid = true
	}
	return reflect.ValueOf(decodedString)
}
