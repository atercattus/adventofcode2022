package solver

import (
	"bufio"
	"bytes"
	"fmt"
)

func Solver(raw []byte) int {
	var (
		count int
	)

	s := bufio.NewScanner(bytes.NewReader(raw))

	for s.Scan() {
		line := s.Text()

		var f1, t1, f2, t2 int

		fmt.Sscanf(line, "%d-%d,%d-%d", &f1, &t1, &f2, &t2)

		if t1-f1 < t2-f2 {
			f1, f2 = f2, f1
			t1, t2 = t2, t1
		}

		if (f1 <= f2) && (t2 <= t1) {
			count++
		}
	}

	return count
}

func Solver2(raw []byte) int {
	var (
		count int
	)

	s := bufio.NewScanner(bytes.NewReader(raw))

	for s.Scan() {
		line := s.Text()

		var f1, t1, f2, t2 int

		fmt.Sscanf(line, "%d-%d,%d-%d", &f1, &t1, &f2, &t2)

		if f1 > f2 {
			f1, f2 = f2, f1
			t1, t2 = t2, t1
		}

		if (f1 <= f2) && (t1 >= f2) {
			count++
		}
	}

	return count
}
