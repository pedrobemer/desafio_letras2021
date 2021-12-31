package score

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMeasureWordScore(t *testing.T) {

	type test struct {
		firstWord     string
		secondWord    string
		expectedScore int
		expectedError error
	}

	tests := []test{
		{
			firstWord:     "testGreater",
			secondWord:    "test",
			expectedScore: 0,
			expectedError: errors.New("firstWord is greater than secondWord"),
		},
		{
			firstWord:     "Havana",
			secondWord:    "Cavunol",
			expectedScore: 3,
			expectedError: nil,
		},
	}

	for _, testCase := range tests {
		score, err := measureWordScore(testCase.firstWord, testCase.secondWord)
		assert.Equal(t, testCase.expectedScore, score)
		assert.Equal(t, testCase.expectedError, err)
	}
}

func TestMeasureTotalScore(t *testing.T) {
	type test struct {
		searchTitle   string
		musicTitle    string
		expectedScore int
		expectedError error
	}

	tests := []test{
		{
			searchTitle:   "Havana",
			musicTitle:    "Havana feat Young Thug",
			expectedScore: 11,
			expectedError: nil,
		},
		{
			searchTitle:   "Havana feat Young Thug",
			musicTitle:    "Havana feat Young Thug",
			expectedScore: 61,
			expectedError: nil,
		},
		{
			searchTitle:   "Coração Test",
			musicTitle:    "Falta Coração",
			expectedScore: 18,
			expectedError: nil,
		},
	}

	for _, testCase := range tests {
		score, err := MeasureTotalScore(testCase.searchTitle, testCase.musicTitle)

		assert.Equal(t, testCase.expectedScore, score)
		assert.Equal(t, testCase.expectedError, err)
	}
}
