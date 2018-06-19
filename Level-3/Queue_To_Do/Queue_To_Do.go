package Queue_To_Do

import (
	"fmt"
)

// f(a,b) = f(1,b) ^ f(a - 1)
func answer(start int, length int) int {

	v1 := start
	v2 := start

	result := 0

	for i:=0; i < length; i++ {
		v1 = v2 + i
		v2 = v1 + length - i -1

		fmt.Printf("v1:%d v2:%d\n", v1, v2)
		fmt.Printf("xor(v2):%d ^ xor(v1-1):%d\n", xor(v2), xor(v1-1))
		result ^=  xor(v2) ^ xor(v1-1)
	}

	return result
}

//Calculate XOR for 0....N

//Number Binary-Repr  XOR-from-1-to-n
//1         1           [0001]
//2        10           [0011]
//3        11           [0000]  <----- We get a 0
//4       100           [0100]  <----- Equals to n
//5       101           [0001]
//6       110           [0111]
//7       111           [0000]  <----- We get 0
//8      1000           [1000]  <----- Equals to n
//9      1001           [0001]
//10     1010           [1011]
//11     1011           [0000] <------ We get 0
//12     1100           [1100] <------ Equals to n
func xor(n int) int {

	if n == 0 {
		return 0
	}

	switch n & 3 {
	case 0:
		return n
	case 1:
		return 1
	case 2:
		return n + 1
	case 3:
		return 0
	}

	return 0
}