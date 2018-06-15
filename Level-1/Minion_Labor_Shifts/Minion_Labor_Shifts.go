package Minion_Labor_Shifts

func answer(data []int, n int) []int {
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
