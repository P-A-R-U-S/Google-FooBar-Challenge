package Minion_Labor_Shifts

import (
	"testing"
	"fmt"
)

func answer(data []int, n int) []int {

	var result []int

	mlc := make(map[int]int)
	for i := 0; i < len(data); i++ {
		mlc[data[i]]++
	}
	for id, c := range mlc {
		if c <= n {
			result = append(result, id)
		}
	}
	return result
}

func answer2(data []int, n int) []int {
	i := 0

	for i < len(data) {

		count := 0
		for j := 0; j < len(data); j++ {
			if data[j] == data[i] {
				count++
			}
		}

		var newValues []int

		for j := 0; j < len(data); j++ {
			if data[j] != data[i] {
				newValues = append(newValues, data[j])
				//track++
			}

			if count <= n && j == i {
				newValues = append(newValues, data[j])
			}
		}

		data = []int(newValues)

		if count <= n {
			i++
		}

	}

	return data
}

func Test_Input(t *testing.T) {

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
			t.Errorf("length: result:%d not equal output:%d",len(result), len(testData.output) )
		}

		fmt.Printf("result:(%d) - output:(%d)\n", result, testData.output)
		for i := 0; i < len(result); i++ {
			if testData.output[i] != result[i] {
				t.Errorf("contains: value:%d are not part of output:%d", result[i], testData.output)
			}
		}
	}

}

func Test_Input2(t *testing.T) {

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

		result :=  answer2(testData.data, testData.n)

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