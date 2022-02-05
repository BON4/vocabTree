package trees

// v1 - should be longer than v2
func FindDelimeter(v1, v2 string) int {
	if len(v1) < len(v2) {
		v1, v2 = v2, v1
	}

	var i = 0
	for ; i < len(v2); i++ {
		if v1[i] != v2[i] {
			return i
		}
	}
	return i
}

func max(a []int) int {
	t := a[0]
	for _, i := range a {
		if t < i {
			t = i
		}
	}
	return t
}
