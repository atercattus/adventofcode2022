package solver

func Solver(raw []byte) int {
	return solver(raw, 4)
}

func Solver2(raw []byte) int {
	return solver(raw, 14)
}

func solver(raw []byte, cnt int) int {
	cnt--
	for p := cnt; p < len(raw); p++ {
		if diffs(raw[p-cnt : p+1]) {
			return p + 1
		}
	}

	return 0
}

func diffs(quad []byte) (diffs bool) {
	var m [256]byte
	for _, ch := range quad {
		if m[ch]++; m[ch] > 1 {
			return false
		}
	}
	return true
}
