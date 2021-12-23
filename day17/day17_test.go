package day17

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
	target, err := ParseInput(f)
	require.NoError(t, err)
	assert.Equal(t, &Target{MinX: 20, MaxX: 30, MinY: -10, MaxY: -5}, target)
}

func Test_CalculateHighestYPosition_Part1(t *testing.T) {
	f, err := os.Open("testdata/input01")
	require.NoError(t, err)
	defer f.Close()
	target, err := ParseInput(f)
	require.NoError(t, err)
	assert.Equal(t, 45, CalculateHighestYPosition(target))
}

func Test_CalculateHighestYPosition_Part2(t *testing.T) {
	f, err := os.Open("testdata/input02")
	require.NoError(t, err)
	defer f.Close()
	target, err := ParseInput(f)
	require.NoError(t, err)
	assert.Equal(t, 5778, CalculateHighestYPosition(target))
}

func Test_CalculateDistinctInitialVelocities_Part1(t *testing.T) {
	f, err := os.Open("testdata/input01")
	require.NoError(t, err)
	defer f.Close()
	target, err := ParseInput(f)
	require.NoError(t, err)
	assert.Equal(t, 112, CalculateDistinctInitialVelocities(target))
}

func Test_CalculateDistinctInitialVelocities_Part2(t *testing.T) {
	f, err := os.Open("testdata/input02")
	require.NoError(t, err)
	defer f.Close()
	target, err := ParseInput(f)
	require.NoError(t, err)
	assert.Equal(t, 2576, CalculateDistinctInitialVelocities(target))
}
