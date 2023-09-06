package dev02

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnpacking(t *testing.T) {
	testTable := []struct {
		name             string
		input            string
		expectedResponse string
		expectedError    error
	}{
		{
			name:             "check a4bc2d5e",
			input:            "a4bc2d5e",
			expectedResponse: "aaaabccddddde",
			expectedError:    nil,
		},
		{
			name:             "check abcd",
			input:            "abcd",
			expectedResponse: "abcd",
			expectedError:    nil,
		},
		{
			name:             "check 45",
			input:            "45",
			expectedResponse: "",
			expectedError:    fmt.Errorf("некорректная строка"),
		},
		{
			name:             `check qwe\4\5`,
			input:            `qwe\4\5`,
			expectedResponse: `qwe45`,
			expectedError:    nil,
		},
		{
			name:             `check qwe\45`,
			input:            `qwe\45`,
			expectedResponse: `qwe44444`,
			expectedError:    nil,
		},
		{
			name:             `check qwe\\5`,
			input:            `qwe\\5`,
			expectedResponse: `qwe\\\\\`,
			expectedError:    nil,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := Unpacking(testCase.input)
			assert.Equal(t, testCase.expectedResponse, result)
			assert.Equal(t, testCase.expectedError, err)
		})
	}
}
