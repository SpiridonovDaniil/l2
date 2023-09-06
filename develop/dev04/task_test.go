package dev04

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchAnagram(t *testing.T) {
	testTable := []struct {
		name             string
		input            []string
		expectedResponse map[string][]string
	}{
		{
			name:             "кейс 1",
			input:            []string{"пятка", "тяпка", "пятак", "тряпка", "растяпа", "пот", "топ"},
			expectedResponse: map[string][]string{"пот": {"пот", "топ"}, "пятка": {"пятак", "пятка", "тяпка"}},
		},
		{
			name:             "кейс 2",
			input:            []string{"листок", "рп", "пр", "столик", "облако", "балкон", "слиток"},
			expectedResponse: map[string][]string{"листок": {"листок", "слиток", "столик"}, "рп": {"пр", "рп"}},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			result := SearchAnagram(testCase.input)
			assert.Equal(t, testCase.expectedResponse, result)
		})
	}
}
