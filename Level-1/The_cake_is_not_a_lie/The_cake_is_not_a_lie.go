package The_cake_is_not_a_lie

import "strings"

func answer1(x string) int  {

	ret := 1
	l := len(x)

	var pl []uint8
	var pr []uint8

	for i :=0; i < l; i++ {
		cl := x[i]
		cr := x[l-1-i]

		pl = append(pl, cl)
		pr = append(pr, cr)

		if pl[0] != cr || pr[0] != cl {
			continue
		}

		ret = l / len(pl)

		//Verify Pattern
		if strings.Compare(x, strings.Repeat(string([]byte(pl[:])),ret)) != 0 {
			continue
		}

		ret = l / len(pl)

		break

	}

	return ret
}