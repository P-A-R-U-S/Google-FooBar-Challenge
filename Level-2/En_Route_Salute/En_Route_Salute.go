package En_Route_Salute

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

