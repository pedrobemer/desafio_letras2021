package treatment

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTreatString(t *testing.T) {

	type test struct {
		text                  string
		expectedTreatedString string
		expectedError         error
	}

	tests := []test{
		{
			text:                  "TeST",
			expectedTreatedString: "test",
			expectedError:         nil,
		},
		{
			text:                  "TrÉm bÃlà",
			expectedTreatedString: "trem bala",
			expectedError:         nil,
		},
		{
			text:                  "",
			expectedTreatedString: "",
			expectedError:         nil,
		},
	}

	for _, testCase := range tests {
		treatedString, err := TreatString(testCase.text)
		assert.Equal(t, testCase.expectedTreatedString, treatedString)
		assert.Equal(t, testCase.expectedError, err)

	}

}
