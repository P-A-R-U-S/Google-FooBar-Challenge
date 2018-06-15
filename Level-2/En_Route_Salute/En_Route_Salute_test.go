package En_Route_Salute

import "testing"

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
