package solver

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var in = []byte(`2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`)

func Test(t *testing.T) {
	require.Equal(t, 2, Solver(in))

	require.Equal(t, 4, Solver2(in))
}

func Test_Evaluate(t *testing.T) {
	raw, err := os.ReadFile("input")
	require.NoError(t, err)

	fmt.Println("#1:", Solver(raw))
	fmt.Println("#2:", Solver2(raw))
}
