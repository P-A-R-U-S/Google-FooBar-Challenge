package Queue_To_Do

import (
	"testing"
	"fmt"
)

func Test_Queue_To_Do(t *testing.T) {
	testDatas := [] struct{
		start int
		lenght int
		result int
	} {
		{start:0, lenght:1, result:0},
		{start:1, lenght:1, result:1},
		{start:0, lenght:3, result:2},
		{start:17, lenght:4, result:14},
		{start:30, lenght:10, result:57},
		{start:100, lenght:10, result:119},
		{start:0, lenght:200000, result:40918466560},
		{start:0, lenght:200001, result:25372527936},
		{start:1, lenght:200000, result:65032293696},
		{start:1, lenght:200001, result:30913895745},
	}

	for _, testData := range testDatas {
		//result := answer(testData.start, testData.lenght)
		result := answer(testData.start, testData.lenght)

		t.Logf("result:%d", result)
		if result != testData.result {

			t.Errorf("expected:%d, actual:%d", testData.result, result)
		}
	}
}

func Test_XOR(t *testing.T) {
	testDatas := [] struct{
		n int
		result int
	}{
		{5, 1},
		{6, 7},
		{7, 0 },
		{8, 8 },

		{9, 1 },
		{10, 11 },
		{11, 0 },
		{12, 12 },

		{25, 1 },
		{26, 27 },
		{27, 0 },
		{28, 28 },

		{1453, 1 },
		{1454, 1455 },
		{1455, 0 },
		{1456, 1456 },
		//{29,29, 0},
	}

	for _, testData := range testDatas {

		result1 := 0
		for i:= 0; i <= testData.n; i++ {

			result1 ^= i
		}

		result2 := xor(testData.n)

		fmt.Printf("result1:%d, result2:%d\n", result1, result2)
		if result1 != result2 {
			t.Errorf("Exoected: %d, Actual:%d", result1, result2)
		}

		if result1 != testData.result {
			t.Errorf("Result1 Wrong: %d, Expected:%d", result1, testData.result)
		}

		if result2 != testData.result {
			t.Errorf("Result1 Wrong: %d, Expected:%d", result2, testData.result)
		}
	}
}

func Test_XOR_range(t *testing.T) {
	testDatas := [] struct{
		a int
		b int
		result int
	}{
		{0, 0, 0},
		{0, 6, 7},
		{0, 6, 7},
		{3, 7, 3},
		{8, 16, 16},
		{4, 12, 12},
		{8, 23, 0},
		{8, 24, 24},
		{8, 25, 1},
		{8, 26, 27},
		{8, 27, 0},
		{8, 28, 28},
		{8, 29, 1},
		{8, 30, 31},
		{8, 31, 0},
		{8, 32, 32},
		{8, 33, 1},
		{8, 34, 35},

		{17, 20, 4},
		{21, 23, 20},
		{25, 26, 3},
		{29, 29, 29},

		{54, 54, 54},
		{234890, 234890, 234890},
		{1, 1, 1},
		{66, 66, 66},
		{345, 75391, 344},
		{0, 99999999, 0},
		{2, 99999999, 1},
		{3, 99999999, 3},
		{4, 99999999, 0},
		{0, 2000000000, 0},

	}

	for _, testData := range testDatas {

		result1 := 0
		for i := testData.a; i <= testData.b; i++ {
			result1 ^= i
		}

		result2 := xor(testData.b)	 ^ xor(testData.a - 1)
		fmt.Printf("a:%d, b:%d = result1:%d (Expected:%d) = result2:%d\n",testData.a, testData.b, result1, testData.result, result2)

		if result1 !=  result2 {
			t.Errorf("a:%d, b:%d = result1:%d (Expected:%d) = result2:%d\n",testData.a, testData.b, result1, testData.result, result2)
		}
	}
}