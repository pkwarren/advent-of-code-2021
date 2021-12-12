package day07

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ParseInput(t *testing.T) {
	f, err := os.Open("testdata/input01")
	require.NoError(t, err)
	nums, err := ParseInput(f)
	require.NoError(t, err)
	assert.Equal(t, []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}, nums)
}

func Test_CalculateMinFuelPosition_Part1_Input1(t *testing.T) {
	f, err := os.Open("testdata/input01")
	require.NoError(t, err)
	nums, err := ParseInput(f)
	require.NoError(t, err)
	pos, fuel := CalculateMinFuelPosition(nums)
	assert.Equal(t, 2, pos)
	assert.Equal(t, 37, fuel)
}

func Test_CalculateMinFuelPosition_Part1_Input2(t *testing.T) {
	f, err := os.Open("testdata/input02")
	require.NoError(t, err)
	nums, err := ParseInput(f)
	require.NoError(t, err)
	pos, fuel := CalculateMinFuelPosition(nums)
	assert.Equal(t, 339, pos)
	assert.Equal(t, 355764, fuel)
}

func Test_CalculateMinFuelPosition_Part2_Input1(t *testing.T) {
	f, err := os.Open("testdata/input01")
	require.NoError(t, err)
	nums, err := ParseInput(f)
	require.NoError(t, err)
	pos, fuel := CalculateMinFuelPosition_Part2(nums)
	assert.Equal(t, 5, pos)
	assert.Equal(t, 168, fuel)
}

func Test_CalculateMinFuelPosition_Part2_Input2(t *testing.T) {
	f, err := os.Open("testdata/input02")
	require.NoError(t, err)
	nums, err := ParseInput(f)
	require.NoError(t, err)
	pos, fuel := CalculateMinFuelPosition_Part2(nums)
	assert.Equal(t, 485, pos)
	assert.Equal(t, 99634572, fuel)
}
