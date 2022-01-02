package comparator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsentive(t *testing.T) {

	type test struct {
		firstWord             string
		secondWord            string
		expectedResult        bool
		expectedLenComparison int
	}

	tests := []test{
		{
			firstWord:             "test",
			secondWord:            "testLower",
			expectedResult:        false,
			expectedLenComparison: 2,
		},
		{
			firstWord:             "test",
			secondWord:            "ta",
			expectedResult:        false,
			expectedLenComparison: 1,
		},
		{
			firstWord:             "test",
			secondWord:            "teco",
			expectedResult:        false,
			expectedLenComparison: 1,
		},
		{
			firstWord:             "test",
			secondWord:            "test",
			expectedResult:        true,
			expectedLenComparison: 1,
		},
	}

	for _, testCase := range tests {
		result, lenComp := Insensitive(testCase.firstWord, testCase.secondWord)
		assert.Equal(t, testCase.expectedResult, result)
		assert.Equal(t, testCase.expectedLenComparison, lenComp)
	}

}
