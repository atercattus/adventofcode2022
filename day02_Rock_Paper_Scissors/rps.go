package day01

import (
	"bufio"
	"bytes"
	"fmt"
)

// https://adventofcode.com/2022/day/2

type T byte

const (
	_        = T(iota)
	Rock     // A, X
	Paper    // B, Y
	Scissors // C, Z

	LoseScore = 0
	DrawScore = 3
	WinScore  = 6
)

var (
	Loses = map[T]T{Rock: Paper, Scissors: Rock, Paper: Scissors}
	Wins  = map[T]T{Rock: Scissors, Scissors: Paper, Paper: Rock}
)

func GetScore(raw []byte) int {
	var (
		score int
	)

	s := bufio.NewScanner(bytes.NewReader(raw))

	for s.Scan() {
		var yr, or rune
		fmt.Sscanf(s.Text(), "%c %c", &or, &yr)
		y := t(yr)
		o := t(or)

		ys := yourScore(y, o) + int(y)

		score += ys
	}

	return score
}

func GetStrategyScore(raw []byte) int {
	var (
		score int
	)

	s := bufio.NewScanner(bytes.NewReader(raw))

	for s.Scan() {
		var yr, or rune
		fmt.Sscanf(s.Text(), "%c %c", &or, &yr)
		o := t(or)

		ys := yourStrategyScore(yr, o)

		score += ys
	}

	return score
}

func t(r rune) T {
	switch r {
	case 'A', 'X':
		return Rock
	case 'B', 'Y':
		return Paper
	case 'C', 'Z':
		return Scissors
	}
	panic(fmt.Sprintf("wrong rune %v)", r))
}

func yourScore(you, other T) int {
	if you == other {
		return DrawScore
	}

	if Wins[you] == other {
		return WinScore
	}

	return LoseScore
}

func yourStrategyScore(youRune rune, other T) int {
	switch youRune {
	case 'X': // lose
		score := Wins[other]
		return LoseScore + int(score)
	case 'Y': // draw
		return DrawScore + int(other)
	case 'Z': // win
		score := Loses[other]
		return WinScore + int(score)
	}

	panic(fmt.Sprintf("wrong rune %v)", youRune))
}
