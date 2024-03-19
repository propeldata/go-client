package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertStringToJSONValue(t *testing.T) {
	tests := []struct {
		name           string
		value          string
		propelType     PropelType
		expectedResult any
		expectedError  string
	}{
		{
			name:           "Array",
			value:          `[{"some_field":3}, "hey"]`,
			propelType:     JsonPropelType,
			expectedResult: []any{map[string]any{"some_field": float64(3)}, "hey"},
			expectedError:  "",
		},
		{
			name:           "Simple object",
			value:          `{"some_field":true, "value":"abc"}`,
			propelType:     JsonPropelType,
			expectedResult: map[string]any{"some_field": true, "value": "abc"},
			expectedError:  "",
		},
		{
			name:           "Simple string",
			value:          `daisies`,
			propelType:     JsonPropelType,
			expectedResult: "daisies",
			expectedError:  "",
		},
		{
			name:           "Simple boolean",
			value:          `false`,
			propelType:     JsonPropelType,
			expectedResult: false,
			expectedError:  "",
		},
		{
			name:           "Simple integer",
			value:          `30`,
			propelType:     JsonPropelType,
			expectedResult: float64(30),
			expectedError:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(st *testing.T) {
			a := assert.New(st)

			propelType, err := ConvertStringToJSONValue(tt.value, tt.propelType)
			a.Equal(tt.expectedResult, propelType)

			if tt.expectedError != "" {
				a.Equal(tt.expectedError, err.Error())
			} else {
				a.NoError(err)
			}
		})
	}
}
