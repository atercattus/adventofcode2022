package solver

import (
	"bufio"
	"bytes"
	"fmt"
)

type pathItem struct {
	parent int
	name   string
	size   int
}

func Solver(raw []byte) int {
	content := parse(bufio.NewScanner(bytes.NewBuffer(raw)))

	atMostSum := 0
	for _, p := range content {
		if p.size <= 100_000 {
			atMostSum += p.size
		}
	}

	return atMostSum
}

func Solver2(raw []byte) int {
	content := parse(bufio.NewScanner(bytes.NewBuffer(raw)))

	free := 70_000_000 - content[1].size

	selectedSize := content[1].size
	for _, p := range content {
		if (free+p.size >= 30_000_000) && (p.size < selectedSize) {
			selectedSize = p.size
		}
	}

	return selectedSize
}

func parse(s *bufio.Scanner) []pathItem {
	var content []pathItem
	var contentPos int

	content = append(content, pathItem{
		parent: 0,
		name:   "",
		size:   0,
	})

	for s.Scan() {
		line := s.Bytes()
		if len(line) == 0 {
			continue
		}

		var size int

		if n, err := fmt.Sscanf(string(line), "%d ", &size); n > 0 && err == nil {
			content[contentPos].size += size
		}

		if bytes.HasPrefix(line, []byte("$ cd")) {
			dir := string(line[5:])
			switch dir {
			case "..":
				contentPos = content[contentPos].parent
			default:
				content = append(content, pathItem{
					parent: contentPos,
					name:   dir,
					size:   0,
				})
				contentPos = len(content) - 1
			}
		}
	}

	for i := len(content) - 1; i >= 0; i-- {
		if content[i].parent != 0 {
			content[content[i].parent].size += content[i].size
		}
	}

	return content
}
