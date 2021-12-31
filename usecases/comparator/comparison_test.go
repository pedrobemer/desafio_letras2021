package comparator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsentive(t *testing.T) {

	type test struct {
		firstWord      string
		secondWord     string
		expectedResult bool
	}

	tests := []test{
		{
			firstWord:      "test",
			secondWord:     "testLower",
			expectedResult: false,
		},
		{
			firstWord:      "test",
			secondWord:     "teco",
			expectedResult: false,
		},
		{
			firstWord:      "test",
			secondWord:     "test",
			expectedResult: true,
		},
	}

	for _, testCase := range tests {
		result := Insensitive(testCase.firstWord, testCase.secondWord)
		assert.Equal(t, testCase.expectedResult, result)
	}

}
