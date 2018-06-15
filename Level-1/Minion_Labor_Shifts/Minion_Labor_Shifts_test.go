package Minion_Labor_Shifts

import (
	"testing"
	"fmt"
)

func Test_Minion_Labor_Shifts(t *testing.T) {

	testDatas := []struct {
		data []int
		n int
		output []int
	} {
		{[]int {1, 2, 3 }, 0, []int{}},
		{[]int {1, 2, 2, 3, 3, 3, 4, 5, 5},1, []int {1,4}},
		{[]int {1, 2, 3}, 6, []int {1, 2, 3}},
		{[]int {1, 2, 2, 3, 3, 3, 4, 5, 5, 6, 6, 6, 7, 7}, 2, []int {1, 2, 4, 5, 7}},
		{[]int {2, 2, 1, 6, 3, 3, 4, 7, 5, 6, 3, 6, 7, 5}, 2, []int {2, 1, 4, 7, 5}},
		{[]int {2, 2, 1, 6, 3, 3, 4, 7, 5, 6, 3, 6, 7, 5}, 3, []int {2, 1, 6, 3, 4, 7, 5}},
	}

	for _,testData := range testDatas {

		result :=  answer(testData.data, testData.n)

		if len(result) != len(testData.output) {
			t.Errorf("length: result:%d(%d) not equal output:%d (%d)",len(result), result,  len(testData.output), testData.output )
		}

		fmt.Printf("result:(%d) - output:(%d)\n", result, testData.output)
		for i := 0; i < len(result); i++ {
			if testData.output[i] != result[i] {
				t.Errorf("contains: value:%d are not part of output:%d", result[i], testData.output)
			}
		}
	}

}
