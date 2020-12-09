package solution

func Solution(A []int) int {
	set := make(map[int]int)
	for _, v := range A {
		set[v] = 1
	}

	s := 0
	e := 1
	delete(set, A[0])

	for ; e < len(A); e++ {
		if A[e] == A[s] {
			s++
		}

		for s <= e && A[s] == A[s+1] {
			s++
		}

		delete(set, A[e])
		if len(set) == 0 {
			break
		}
	}
	return e - s + 1
}
