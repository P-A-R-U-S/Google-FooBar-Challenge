package Prison_Labor_Dodgers

import (
	"testing"
)

/*
Test cases
==========

Inputs:
    (int list) x = [13, 5, 6, 2, 5]
    (int list) y = [5, 2, 5, 13]
Output:
    (int) 6

Inputs:
    (int list) x = [14, 27, 1, 4, 2, 50, 3, 1]
    (int list) y = [2, 4, -4, 3, 1, 1, 14, 27, 50]
Output:
    (int) -4
*/
func Test_Prison_Labor_Dodgers(t *testing.T) {

	testDatas := []struct {
		x      []int
		y      []int
		output int
	}{
		{[]int{13, 5, 6, 2, 5}, []int{5, 2, 5, 13}, 6},
		{[]int{14, 27, 1, 4, 2, 50, 3, 1}, []int{2, 4, -4, 3, 1, 1, 14, 27, 50}, -4},
	}

	for _, testData := range testDatas {

		result := answer(testData.x, testData.y)

		if result != testData.output {
			t.Errorf("Actual: result:%d not equal Expected: output:%d", result, testData.output)
		}

	}

}
