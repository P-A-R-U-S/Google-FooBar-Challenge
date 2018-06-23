package The_cake_is_not_a_lie

import "fmt"

func answer1(x string) int  {

	l := len(x)

	var pl []uint8
	var pr []uint8

	var pls string //fmt.Sprintf("%c",x[0])
	var prs string //fmt.Sprintf("%c",x[l-1])


	for i :=0; i < l; i++ {
		cl := x[i]
		cr := x[l-1-i]

		pls = fmt.Sprintf("%s%c",pl, cl)
		prs = fmt.Sprintf("%s%c",pr, cr)

		pl = append(pl, cl)
		pr = append(pr, cr)

		if i == 0 || pl[0] != cr || pr[0] != cl {

			//pls = fmt.Sprintf("%s%c",pl, cl)
			//prs = fmt.Sprintf("%c%s",cr, pr)

			continue
		} else {

			//pls = fmt.Sprintf("%s%c",pl, cl)
			//prs = fmt.Sprintf("%c%s",cr, pr)

			//pl = append(pl, cl)
			//pr = append(pr, cr)

			fmt.Println(pls)
			fmt.Println(prs)

			return l / len(pl)
		}
	}

	return 1
}