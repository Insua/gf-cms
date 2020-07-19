package slice

func InInt(s []int, v int) (in bool) {
	in = false
	for _, value := range s {
		if value == v {
			in = true
			break
		}
	}
	return
}

func InString(s []string, v string) (in bool) {
	in = false
	for _, value := range s {
		if value == v {
			in = true
			break
		}
	}
	return
}
