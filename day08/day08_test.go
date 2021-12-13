package day08

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ParseInput(t *testing.T) {
	f, err := os.Open("testdata/input01")
	require.NoError(t, err)
	defer f.Close()
	lines := 0
	err = ParseInput(f, func(signals, outputs []string) error {
		lines += 1
		return nil
	})
	require.NoError(t, err)
	assert.Equal(t, 10, lines)
}

func Test_FindUniqueSegments_Part1_Input1(t *testing.T) {
	f, err := os.Open("testdata/input01")
	require.NoError(t, err)
	defer f.Close()
	n, err := FindUniqueSegments(f)
	require.NoError(t, err)
	assert.Equal(t, 26, n)
}

func Test_FindUniqueSegments_Part1_Input2(t *testing.T) {
	f, err := os.Open("testdata/input02")
	require.NoError(t, err)
	defer f.Close()
	n, err := FindUniqueSegments(f)
	require.NoError(t, err)
	assert.Equal(t, 365, n)
}

func Test_DecodeLines_Part2_Input0(t *testing.T) {
	r := strings.NewReader(`acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf`)
	lines, err := DecodeLines(r)
	require.NoError(t, err)
	assert.Equal(t, []int{5353}, lines)
}

func Test_DecodeLines_Part2_Input1(t *testing.T) {
	f, err := os.Open("testdata/input01")
	require.NoError(t, err)
	defer f.Close()
	lines, err := DecodeLines(f)
	require.NoError(t, err)
	assert.Equal(t, []int{8394, 9781, 1197, 9361, 4873, 8418, 4548, 1625, 8717, 4315}, lines)
	sum := 0
	for _, l := range lines {
		sum += l
	}
	assert.Equal(t, 61229, sum)
}

func Test_DecodeLines_Part2_Input2(t *testing.T) {
	f, err := os.Open("testdata/input02")
	require.NoError(t, err)
	defer f.Close()
	lines, err := DecodeLines(f)
	require.NoError(t, err)
	sum := 0
	for _, l := range lines {
		sum += l
	}
	assert.Equal(t, 975706, sum)
}
