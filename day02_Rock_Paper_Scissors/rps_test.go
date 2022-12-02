package day01

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

var in = []byte(`A Y
B X
C Z`)

func Test(t *testing.T) {
	require.Equal(t, 15, GetScore(in))

	require.Equal(t, 12, GetStrategyScore(in))
}

func Test_Evaluate(t *testing.T) {
	raw, err := ioutil.ReadFile("input")
	require.NoError(t, err)

	fmt.Println("naive:", GetScore(raw))
	fmt.Println("strategy:", GetStrategyScore(raw))
}
