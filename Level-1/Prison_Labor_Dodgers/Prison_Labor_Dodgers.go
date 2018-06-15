package Prison_Labor_Dodgers

func answer(x []int, y []int) int {

	data := make(map[int]int)

	for i:=0; i < len(x); i++ {
		data[x[i]] = 1
	}

	for i:=0; i < len(y); i++ {
		data[x[i]] = -1
	}

	for k, v := range data {
		if v == 1 || v == -1 {
			return k
		}
	}

	return 0
}
