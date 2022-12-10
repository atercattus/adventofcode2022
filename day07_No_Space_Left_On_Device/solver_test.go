package solver

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var in = []byte(`
$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k
`)

func Test(t *testing.T) {
	require.Equal(t, 95437, Solver(in))

	require.Equal(t, 24933642, Solver2(in))
}

func Test_Evaluate(t *testing.T) {
	raw, err := os.ReadFile("input")
	require.NoError(t, err)

	fmt.Println("#1:", Solver(raw))
	fmt.Println("#2:", Solver2(raw))
}
