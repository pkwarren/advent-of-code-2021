package day09

import (
	"os"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_SumOfRiskLevels_Part1_Input1(t *testing.T) {
	f, err := os.Open("testdata/input01")
	require.NoError(t, err)
	defer f.Close()
	n, err := SumOfRiskLevels(f)
	require.NoError(t, err)
	assert.Equal(t, 15, n)
}

func Test_SumOfRiskLevels_Part1_Input2(t *testing.T) {
	f, err := os.Open("testdata/input02")
	require.NoError(t, err)
	defer f.Close()
	n, err := SumOfRiskLevels(f)
	require.NoError(t, err)
	assert.Equal(t, 572, n)
}

func Test_FindBasinSizes_Part2_Input1(t *testing.T) {
	f, err := os.Open("testdata/input01")
	require.NoError(t, err)
	defer f.Close()
	basinSizes, err := FindBasinSizes(f)
	require.NoError(t, err)
	sort.Ints(basinSizes)
	assert.Equal(t, []int{3, 9, 9, 14}, basinSizes)
}

func Test_FindBasinSizes_Part2_Input2(t *testing.T) {
	f, err := os.Open("testdata/input02")
	require.NoError(t, err)
	defer f.Close()
	basinSizes, err := FindBasinSizes(f)
	require.NoError(t, err)
	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))
	require.True(t, len(basinSizes) >= 3)
	assert.Equal(t, 847044, basinSizes[0]*basinSizes[1]*basinSizes[2])
}
