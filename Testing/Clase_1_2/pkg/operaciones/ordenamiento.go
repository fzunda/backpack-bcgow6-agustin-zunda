package operaciones

func ListOrdenadaMayorMenor(listNum []int, Cant int) []int {
	tmp := 0
	for x := 0; x < Cant; x++ {
		for y := 0; y < Cant; y++ {
			if listNum[x] < listNum[y] {
				tmp = listNum[y]
				listNum[y] = listNum[x]
				listNum[x] = tmp
			}
		}
	}

	return listNum
}
