package The_cake_is_not_a_lie

func answer1(x string) int  {

	ret := 1
	l := len(x)

	var pl []uint8
	var pr []uint8

	// Uncomment for Debug
	//var pls string
	//var prs string


	for i :=0; i < l; i++ {
		cl := x[i]
		cr := x[l-1-i]

		// Uncomment for Debug
		//pls = fmt.Sprintf("%s%c",pl, cl)
		//prs = fmt.Sprintf("%s%c",pr, cr)

		pl = append(pl, cl)
		pr = append(pr, cr)

		if i == 0 || pl[0] != cr || pr[0] != cl {
			continue
		}

		// Uncomment for Debug
		//fmt.Println(pls)
		//fmt.Println(prs)

		ret = l / len(pl)

		if len(pl) == 2 && pl[0] == pl[len(pl)-1] {
			ret = ret * 2
		}
		break

	}

	return ret
}