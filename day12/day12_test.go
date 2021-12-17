package day12

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_CountDistinctPaths_Part1_Input1(t *testing.T) {
	f, err := os.Open("testdata/input01")
	require.NoError(t, err)
	defer f.Close()
	m, err := ParseInput(f)
	require.NoError(t, err)
	paths := m.CountDistinctPaths(false)
	assert.Equal(t, 10, paths)
}

func Test_CountDistinctPaths_Part1_Input2(t *testing.T) {
	f, err := os.Open("testdata/input02")
	require.NoError(t, err)
	defer f.Close()
	m, err := ParseInput(f)
	require.NoError(t, err)
	paths := m.CountDistinctPaths(false)
	assert.Equal(t, 19, paths)
}

func Test_CountDistinctPaths_Part1_Input3(t *testing.T) {
	f, err := os.Open("testdata/input03")
	require.NoError(t, err)
	defer f.Close()
	m, err := ParseInput(f)
	require.NoError(t, err)
	paths := m.CountDistinctPaths(false)
	assert.Equal(t, 226, paths)
}

func Test_CountDistinctPaths_Part1_Input4(t *testing.T) {
	f, err := os.Open("testdata/input04")
	require.NoError(t, err)
	defer f.Close()
	m, err := ParseInput(f)
	require.NoError(t, err)
	paths := m.CountDistinctPaths(false)
	assert.Equal(t, 3463, paths)
}

func Test_CountDistinctPaths_Part2_Input1(t *testing.T) {
	f, err := os.Open("testdata/input01")
	require.NoError(t, err)
	defer f.Close()
	m, err := ParseInput(f)
	require.NoError(t, err)
	paths := m.CountDistinctPaths(true)
	assert.Equal(t, 36, paths)
}

func Test_CountDistinctPaths_Part2_Input2(t *testing.T) {
	f, err := os.Open("testdata/input02")
	require.NoError(t, err)
	defer f.Close()
	m, err := ParseInput(f)
	require.NoError(t, err)
	paths := m.CountDistinctPaths(true)
	assert.Equal(t, 103, paths)
}

func Test_CountDistinctPaths_Part2_Input3(t *testing.T) {
	f, err := os.Open("testdata/input03")
	require.NoError(t, err)
	defer f.Close()
	m, err := ParseInput(f)
	require.NoError(t, err)
	paths := m.CountDistinctPaths(true)
	assert.Equal(t, 3509, paths)
}

func Test_CountDistinctPaths_Part2_Input4(t *testing.T) {
	f, err := os.Open("testdata/input04")
	require.NoError(t, err)
	defer f.Close()
	m, err := ParseInput(f)
	require.NoError(t, err)
	paths := m.CountDistinctPaths(true)
	assert.Equal(t, 91533, paths)
}
