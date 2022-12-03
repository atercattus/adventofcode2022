package solver

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var in = []byte(`vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`)

func Test(t *testing.T) {
	require.Equal(t, 157, Solver(in))

	require.Equal(t, 70, Solver2(in))
}

func Test_Evaluate(t *testing.T) {
	raw, err := os.ReadFile("input")
	require.NoError(t, err)

	fmt.Println("#1:", Solver(raw))
	fmt.Println("#2:", Solver2(raw))
}
