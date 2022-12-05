package solver

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var in = []byte(`
    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`)

func Test(t *testing.T) {
	require.Equal(t, "CMZ", Solver(in))

	require.Equal(t, "MCD", Solver2(in))
}

func Test_Evaluate(t *testing.T) {
	raw, err := os.ReadFile("input")
	require.NoError(t, err)

	fmt.Println("#1:", Solver(raw))
	fmt.Println("#2:", Solver2(raw))
}
