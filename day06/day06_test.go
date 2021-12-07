package day06

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
	require.Equal(t, []int{0, 1, 1, 2, 1, 0, 0, 0, 0}, s.FishDays)
}

func Test_Simulate_Part1_Input1(t *testing.T) {
	f, err := os.Open("testdata/input01")
	require.NoError(t, err)
	defer f.Close()
	s, err := ParseInput(f)
	require.NoError(t, err)
	require.Equal(t, []int{0, 1, 1, 2, 1, 0, 0, 0, 0}, s.FishDays)
	t.Logf("state[%d]: %v", 0, s.String())
	for i := 0; i < 18; i++ {
		s.Advance()
		t.Logf("state[%d]: %v", i+1, s.String())
	}
	assert.Equal(t, 26, s.NumFish())
	for i := 0; i < 80-18; i++ {
		s.Advance()
		t.Logf("state[%d]: %v", i+1, s.String())
	}
	assert.Equal(t, 5934, s.NumFish())
}

func Test_Simulate_Part1_Input2(t *testing.T) {
	f, err := os.Open("testdata/input02")
	require.NoError(t, err)
	defer f.Close()
	s, err := ParseInput(f)
	require.NoError(t, err)
	for i := 0; i < 80; i++ {
		s.Advance()
	}
	assert.Equal(t, 394994, s.NumFish())
}

func Test_Simulate_Part2_Input1(t *testing.T) {
	f, err := os.Open("testdata/input01")
	require.NoError(t, err)
	defer f.Close()
	s, err := ParseInput(f)
	require.NoError(t, err)
	for i := 0; i < 256; i++ {
		s.Advance()
	}
	assert.Equal(t, 26984457539, s.NumFish())
}

func Test_Simulate_Part2_Input2(t *testing.T) {
	f, err := os.Open("testdata/input02")
	require.NoError(t, err)
	defer f.Close()
	s, err := ParseInput(f)
	require.NoError(t, err)
	for i := 0; i < 256; i++ {
		s.Advance()
	}
	assert.Equal(t, 1765974267455, s.NumFish())
}
