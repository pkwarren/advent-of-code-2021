package day11

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
	s, err := ParseInput(f)
	require.NoError(t, err)
	assert.Len(t, s.Grid, 10)
	assert.Len(t, s.Grid[0], 10)
}

func TestOctopusSimulation_Step(t *testing.T) {
	f, err := os.Open("testdata/input01")
	require.NoError(t, err)
	defer f.Close()
	s, err := ParseInput(f)
	require.NoError(t, err)
	s.Step()
	assert.Equal(t, [10][10]int{
		{6, 5, 9, 4, 2, 5, 4, 3, 3, 4},
		{3, 8, 5, 6, 9, 6, 5, 8, 2, 2},
		{6, 3, 7, 5, 6, 6, 7, 2, 8, 4},
		{7, 2, 5, 2, 4, 4, 7, 2, 5, 7},
		{7, 4, 6, 8, 4, 9, 6, 5, 8, 9},
		{5, 2, 7, 8, 6, 3, 5, 7, 5, 6},
		{3, 2, 8, 7, 9, 5, 2, 8, 3, 2},
		{7, 9, 9, 3, 9, 9, 2, 2, 4, 5},
		{5, 9, 5, 7, 9, 5, 9, 6, 6, 5},
		{6, 3, 9, 4, 8, 6, 2, 6, 3, 7},
	}, s.Grid)
	s.Step()
	assert.Equal(t, [10][10]int{
		{8, 8, 0, 7, 4, 7, 6, 5, 5, 5},
		{5, 0, 8, 9, 0, 8, 7, 0, 5, 4},
		{8, 5, 9, 7, 8, 8, 9, 6, 0, 8},
		{8, 4, 8, 5, 7, 6, 9, 6, 0, 0},
		{8, 7, 0, 0, 9, 0, 8, 8, 0, 0},
		{6, 6, 0, 0, 0, 8, 8, 9, 8, 9},
		{6, 8, 0, 0, 0, 0, 5, 9, 4, 3},
		{0, 0, 0, 0, 0, 0, 7, 4, 5, 6},
		{9, 0, 0, 0, 0, 0, 0, 8, 7, 6},
		{8, 7, 0, 0, 0, 0, 6, 8, 4, 8},
	}, s.Grid)
	assert.Equal(t, 35, s.Flashes)
	for i := 0; i < 98; i++ {
		s.Step()
	}
	assert.Equal(t, 1656, s.Flashes)
}

func TestOctopusSimulation_Step_Input2(t *testing.T) {
	f, err := os.Open("testdata/input02")
	require.NoError(t, err)
	defer f.Close()
	s, err := ParseInput(f)
	require.NoError(t, err)
	for i := 0; i < 100; i++ {
		s.Step()
	}
	assert.Equal(t, 1562, s.Flashes)
}

func TestOctopusSimulation_AllFlash_Input1(t *testing.T) {
	f, err := os.Open("testdata/input01")
	require.NoError(t, err)
	defer f.Close()
	s, err := ParseInput(f)
	require.NoError(t, err)
	allzeros := [10][10]int{}
	step := 0
	for {
		s.Step()
		step += 1
		if s.Grid == allzeros {
			assert.Equal(t, 195, step)
			break
		}
	}
}

func TestOctopusSimulation_AllFlash_Input2(t *testing.T) {
	f, err := os.Open("testdata/input02")
	require.NoError(t, err)
	defer f.Close()
	s, err := ParseInput(f)
	require.NoError(t, err)
	allzeros := [10][10]int{}
	step := 0
	for {
		s.Step()
		step += 1
		if s.Grid == allzeros {
			assert.Equal(t, 268, step)
			break
		}
	}
}
