package solver

import (
	"bufio"
	"bytes"
	"fmt"
)

func Solver(raw []byte) string {
	if raw[0] == '\n' {
		raw = raw[1:]
	}

	n := (bytes.IndexByte(raw, '\n') + 1) / 4
	crates := make([][]string, n)

	s := bufio.NewScanner(bytes.NewReader(raw))
	parse(s, crates)

	for s.Scan() {
		line := s.Text()
		var cnt, f, t int
		fmt.Sscanf(line, "move %d from %d to %d", &cnt, &f, &t)
		f--
		t--

		for cnt > 0 {
			l := len(crates[f])
			ch := crates[f][l-1]
			crates[f] = crates[f][:l-1]

			crates[t] = append(crates[t], ch)
			cnt--
		}
	}

	var tops string
	// get tops
	for _, top := range crates {
		tops += top[len(top)-1] // empty?
	}

	return tops
}

func Solver2(raw []byte) string {
	if raw[0] == '\n' {
		raw = raw[1:]
	}

	n := (bytes.IndexByte(raw, '\n') + 1) / 4
	crates := make([][]string, n)

	s := bufio.NewScanner(bytes.NewReader(raw))
	parse(s, crates)

	for s.Scan() {
		line := s.Text()
		var cnt, f, t int
		fmt.Sscanf(line, "move %d from %d to %d", &cnt, &f, &t)
		f--
		t--

		l := len(crates[f])
		chs := crates[f][l-cnt:]
		crates[f] = crates[f][:l-cnt]
		crates[t] = append(crates[t], chs...)
	}

	var tops string
	// get tops
	for _, top := range crates {
		tops += top[len(top)-1] // empty?
	}

	return tops
}

func parse(s *bufio.Scanner, dst [][]string) {
	for s.Scan() {
		line := s.Bytes()
		if bytes.IndexByte(line, '[') == -1 {
			break
		}

		for i := 0; i < len(dst); i++ {
			ch := string(rune(line[1+4*i]))
			if ch != " " {
				dst[i] = append([]string{ch}, dst[i]...)
			}
		}
	}

	s.Scan() // skip empty line
}
