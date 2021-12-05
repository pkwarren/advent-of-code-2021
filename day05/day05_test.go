package day05

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ParseLines(t *testing.T) {
	f, err := os.Open("testdata/input01")
	require.NoError(t, err)
	defer f.Close()
	lines, err := ParseLines(f)
	require.NoError(t, err)
	assert.Len(t, lines, 10)
}

func Test_CountOverlapping_Part1_Input1(t *testing.T) {
	f, err := os.Open("testdata/input01")
	require.NoError(t, err)
	defer f.Close()
	lines, err := ParseLines(f)
	require.NoError(t, err)
	assert.Len(t, lines, 10)
	var m Map
	for _, l := range lines {
		m.Mark(l, false)
	}
	assert.Equal(t, 5, m.OverlappingPoints())
}

func Test_CountOverlapping_Part1_Input2(t *testing.T) {
	f, err := os.Open("testdata/input02")
	require.NoError(t, err)
	defer f.Close()
	lines, err := ParseLines(f)
	require.NoError(t, err)
	assert.Len(t, lines, 500)
	var m Map
	for _, l := range lines {
		m.Mark(l, false)
	}
	assert.Equal(t, 5084, m.OverlappingPoints())
}

func Test_CountOverlapping_Part2_Input1(t *testing.T) {
	f, err := os.Open("testdata/input01")
	require.NoError(t, err)
	defer f.Close()
	lines, err := ParseLines(f)
	require.NoError(t, err)
	assert.Len(t, lines, 10)
	var m Map
	for _, l := range lines {
		m.Mark(l, true)
	}
	assert.Equal(t, 12, m.OverlappingPoints())
}

func Test_CountOverlapping_Part2_Input2(t *testing.T) {
	f, err := os.Open("testdata/input02")
	require.NoError(t, err)
	defer f.Close()
	lines, err := ParseLines(f)
	require.NoError(t, err)
	assert.Len(t, lines, 500)
	var m Map
	for _, l := range lines {
		m.Mark(l, true)
	}
	assert.Equal(t, 17882, m.OverlappingPoints())
}
