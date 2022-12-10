package solver

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	require.Equal(t, 7, Solver([]byte(`mjqjpqmgbljsphdztnvjfqwrcgsmlb`)))
	require.Equal(t, 5, Solver([]byte(`bvwbjplbgvbhsrlpgdmjqwftvncz`)))
	require.Equal(t, 6, Solver([]byte(`nppdvjthqldpwncqszvftbrmjlhg`)))
	require.Equal(t, 10, Solver([]byte(`nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg`)))
	require.Equal(t, 11, Solver([]byte(`zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw`)))

	require.Equal(t, 19, Solver2([]byte(`mjqjpqmgbljsphdztnvjfqwrcgsmlb`)))
	require.Equal(t, 23, Solver2([]byte(`bvwbjplbgvbhsrlpgdmjqwftvncz`)))
	require.Equal(t, 23, Solver2([]byte(`nppdvjthqldpwncqszvftbrmjlhg`)))
	require.Equal(t, 29, Solver2([]byte(`nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg`)))
	require.Equal(t, 26, Solver2([]byte(`zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw`)))
}

func Test_Evaluate(t *testing.T) {
	raw, err := os.ReadFile("input")
	require.NoError(t, err)

	fmt.Println("#1:", Solver(raw))
	fmt.Println("#2:", Solver2(raw))
}
