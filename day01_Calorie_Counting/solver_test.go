package solver

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

var in = []byte(`1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`)

func Test_Day01(t *testing.T) {
	require.Equal(t, 24000, GetMax(in))

	require.Equal(t, 24000, GetMaxN(in, 1))
	require.Equal(t, 45000, GetMaxN(in, 3))
}

func Test_Day01_Evaluate(t *testing.T) {
	raw, err := ioutil.ReadFile("input")
	require.NoError(t, err)

	t.Run("max", func(t *testing.T) {
		t.Parallel()
		fmt.Printf("%s: %d\n", t.Name(), GetMax(raw))
	})

	t.Run("top3 sum", func(t *testing.T) {
		t.Parallel()
		fmt.Printf("%s: %d\n", t.Name(), GetMaxN(raw, 3))
	})
}
