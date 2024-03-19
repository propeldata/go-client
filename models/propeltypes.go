package models

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type PropelType struct {
	id string
}

var (
	BooleanPropelType   = PropelType{id: "BOOLEAN"}
	StringPropelType    = PropelType{id: "STRING"}
	FloatPropelType     = PropelType{id: "FLOAT"}
	DoublePropelType    = PropelType{id: "DOUBLE"}
	Int8PropelType      = PropelType{id: "INT8"}
	Int16PropelType     = PropelType{id: "INT16"}
	Int32PropelType     = PropelType{id: "INT32"}
	Int64PropelType     = PropelType{id: "INT64"}
	DatePropelType      = PropelType{id: "DATE"}
	TimestampPropelType = PropelType{id: "TIMESTAMP"}
	JsonPropelType      = PropelType{id: "JSON"}

	Types = []PropelType{
		BooleanPropelType,
		StringPropelType,
		FloatPropelType,
		DoublePropelType,
		Int8PropelType,
		Int16PropelType,
		Int32PropelType,
		Int64PropelType,
		DatePropelType,
		TimestampPropelType,
		JsonPropelType,
	}
)

func (t *PropelType) String() string {
	return t.id
}

func (t *PropelType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *PropelType) UnmarshalJSON(b []byte) error {
	var name string
	if err := json.Unmarshal(b, &name); err != nil {
		return err
	}

	for _, typ := range Types {
		if name == typ.String() {
			*t = typ
			return nil
		}
	}
	return fmt.Errorf("invalid Propel type %q", name)
}

func ConvertStringToJSONValue(value string, propelType PropelType) (any, error) {
	switch propelType {
	case BooleanPropelType:
		switch value {
		case "0", "f", "F", "false", "FALSE", "off", "OFF", "n", "N", "no", "NO":
			return false, nil
		case "1", "t", "T", "true", "TRUE", "on", "ON", "y", "Y", "yes", "YES":
			return true, nil
		}
	case StringPropelType:
		return value, nil
	case FloatPropelType:
		return strconv.ParseFloat(value, 32)
	case DoublePropelType:
		return strconv.ParseFloat(value, 64)
	case Int8PropelType:
		return strconv.ParseInt(value, 10, 8)
	case Int16PropelType:
		return strconv.ParseInt(value, 10, 16)
	case Int32PropelType:
		return strconv.ParseInt(value, 10, 32)
	case Int64PropelType:
		return strconv.ParseInt(value, 10, 64)
	case DatePropelType:
		return value, nil
	case TimestampPropelType:
		return value, nil
	case JsonPropelType:
		var jsonData any
		if err := json.Unmarshal([]byte(value), &jsonData); err == nil {
			return jsonData, nil
		}

		return value, nil // JSON propel type accepts any kind of data
	}

	return nil, fmt.Errorf("Propel type %q not supported", propelType)
}
