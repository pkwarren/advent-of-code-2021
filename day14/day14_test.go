package day14

import (
	"math"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ParseInput(t *testing.T) {
	f, err := os.Open("testdata/input01")
	require.NoError(t, err)
	i, err := ParseInput(f)
	require.NoError(t, err)
	assert.Equal(t, "NNCB", i.Template)
}

func Test_Part1_Input1(t *testing.T) {
	f, err := os.Open("testdata/input01")
	require.NoError(t, err)
	i, err := ParseInput(f)
	require.NoError(t, err)
	assert.Equal(t, map[rune]int{'N': 2, 'C': 1, 'B': 1}, i.StepMulti(0))
	res := i.StepMulti(10)
	min, max := minMax(res)
	assert.Equal(t, 1588, max-min)
}

func Test_Part1_Input2(t *testing.T) {
	f, err := os.Open("testdata/input02")
	require.NoError(t, err)
	i, err := ParseInput(f)
	require.NoError(t, err)
	res := i.StepMulti(10)
	min, max := minMax(res)
	assert.Equal(t, 3906, max-min)
}

func Test_Part2_Input1(t *testing.T) {
	f, err := os.Open("testdata/input01")
	require.NoError(t, err)
	i, err := ParseInput(f)
	require.NoError(t, err)
	res := i.StepMulti(40)
	min, max := minMax(res)
	assert.Equal(t, 2192039569602, res['B'])
	assert.Equal(t, 3849876073, res['H'])
	assert.Equal(t, 2188189693529, max-min)
}

func Test_Part2_Input2(t *testing.T) {
	f, err := os.Open("testdata/input02")
	require.NoError(t, err)
	i, err := ParseInput(f)
	require.NoError(t, err)
	res := i.StepMulti(40)
	min, max := minMax(res)
	assert.Equal(t, 4441317262452, max-min)
}

func minMax(m map[rune]int) (int, int) {
	maxCount, minCount := 0, math.MaxInt
	for _, v := range m {
		if v > maxCount {
			maxCount = v
		}
		if v < minCount {
			minCount = v
		}
	}
	return minCount, maxCount
}
