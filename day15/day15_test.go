package day15

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ParseInput(t *testing.T) {
	f, err := os.Open("testdata/input01")
	require.NoError(t, err)
	defer f.Close()
	m, err := ParseInput(f)
	require.NoError(t, err)
	assert.Len(t, m, 10)
	assert.Len(t, m[0], 10)
	assert.Equal(t, []int{1, 1, 6, 3, 7, 5, 1, 7, 4, 2}, m[0])
}

func Test_Part1_Input1(t *testing.T) {
	f, err := os.Open("testdata/input01")
	require.NoError(t, err)
	defer f.Close()
	m, err := ParseInput(f)
	require.NoError(t, err)
	assert.Equal(t, 40, FindLowestTotalRisk(m))
}

func Test_Part1_Input2(t *testing.T) {
	f, err := os.Open("testdata/input02")
	require.NoError(t, err)
	defer f.Close()
	m, err := ParseInput(f)
	require.NoError(t, err)
	assert.Equal(t, 685, FindLowestTotalRisk(m))
}

func Test_Part2_Input1(t *testing.T) {
	f, err := os.Open("testdata/input01")
	require.NoError(t, err)
	defer f.Close()
	m, err := ParseInput(f)
	require.NoError(t, err)
	assert.Equal(t, 315, FindLowestTotalRisk(FullMap(m)))
}

func Test_Part2_Input2(t *testing.T) {
	f, err := os.Open("testdata/input02")
	require.NoError(t, err)
	defer f.Close()
	m, err := ParseInput(f)
	require.NoError(t, err)
	assert.Equal(t, 2995, FindLowestTotalRisk(FullMap(m)))
}
