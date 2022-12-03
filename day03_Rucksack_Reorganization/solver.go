package solver

import (
	"bufio"
	"bytes"
)

// https://adventofcode.com/2022/day/3

func Solver(raw []byte) int {
	var (
		score int
	)

	dup := func(s []byte) byte {
		var mem [256]byte // 'z'-'A'+1

		i := 0
		for ; i < len(s)/2; i++ {
			mem[s[i]] = 1
		}
		for ; i < len(s); i++ {
			if mem[s[i]] > 0 {
				return s[i]
			}
		}
		return 0
	}

	s := bufio.NewScanner(bytes.NewReader(raw))

	for s.Scan() {
		line := s.Bytes()
		score += scoring(dup(line))
	}

	return score
}

func Solver2(raw []byte) int {
	var (
		score int
		mem   [256]byte
	)

	s := bufio.NewScanner(bytes.NewReader(raw))

	i := -1
	for s.Scan() {
		i++
		line := s.Bytes()
		for _, b := range line {
			mem[b] |= 1 << i
			if mem[b] == 7 {
				score += scoring(b)
				copy(mem[:], make([]byte, len(mem)))
				i = -1
				break
			}
		}
	}

	return score
}

func scoring(b byte) int {
	if 'a' <= b && b <= 'z' {
		return int(b-'a') + 1
	}
	return int(b-'A') + 27
}
