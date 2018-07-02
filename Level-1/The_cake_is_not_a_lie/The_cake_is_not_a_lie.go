package The_cake_is_not_a_lie

import "strings"

func answer1(x string) int {

	ret := 1
	l := len(x)

	pl := x[0]
	pr := x[l-1]

	for i := 0; i < l; i++ {
		cl := x[i]
		cr := x[l-1-i]

		if pl != cr || pr != cl {
			continue
		}

		ret = l / (i + 1)

		//Verify Pattern
		if strings.Compare(x, strings.Repeat(string([]byte(x[:i+1])), ret)) != 0 {
			continue
		}

		break

	}

	return ret
}
