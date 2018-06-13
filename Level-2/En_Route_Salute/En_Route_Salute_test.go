package En_Route_Salute

import (
	"testing"
)

func answer(s string) int {

	// From left to Right
	result := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '>' {
			for j := i; j < len(s); j++ {
				if s[j] == '<' {
					result++
				}
			}

		}
	}

	return result * 2
}

func Test_Input(t *testing.T) {

	testDatas := []struct{
		value string
		result int``
	} {
		{"--->-><-><-->-", 10},
		{">----<", 2},
		{"<<>><", 4},

	}

	for _,testData := range testDatas {

		result := answer(testData.value)

		// Compare results
		if result != testData.result {
			t.Errorf("Error: Expected:%d, Actual:%d", testData.result, result)
		}
	}
}
