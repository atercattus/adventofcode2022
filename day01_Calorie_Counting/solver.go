package solver

import (
	"bufio"
	"bytes"
	"container/heap"
	"strconv"
)

// https://adventofcode.com/2022/day/1

func GetMax(raw []byte) int {
	var (
		cur int
		max int
	)

	s := bufio.NewScanner(bytes.NewReader(append(raw, "\n\n"...)))

	for s.Scan() {
		v, err := strconv.Atoi(s.Text())
		if err == nil {
			cur += v
			continue
		}

		if cur > max {
			max = cur
		}
		cur = 0
	}

	return max
}

func GetMaxN(raw []byte, n int) int {
	var (
		h   = &IntHeap{}
		cur int
	)

	heap.Init(h)

	s := bufio.NewScanner(bytes.NewReader(append(raw, "\n\n"...)))

	for s.Scan() {
		v, err := strconv.Atoi(s.Text())
		if err == nil {
			cur += v
			continue
		}

		heap.Push(h, cur)
		if h.Len() > n {
			heap.Pop(h)
		}

		cur = 0
	}

	var max int
	for h.Len() > 0 {
		max += h.Pop().(int)
	}

	return max
}
