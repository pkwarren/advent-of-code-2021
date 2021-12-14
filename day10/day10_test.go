package day10

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func Test_TotalSyntaxErrorScore_Input1(t *testing.T) {
	f, err := os.Open("testdata/input01")
	require.NoError(t, err)
	defer f.Close()
	n, err := TotalSyntaxErrorScore(f)
	require.NoError(t, err)
	assert.Equal(t, 26397, n)
}

func Test_TotalSyntaxErrorScore_Input2(t *testing.T) {
	f, err := os.Open("testdata/input02")
	require.NoError(t, err)
	defer f.Close()
	n, err := TotalSyntaxErrorScore(f)
	require.NoError(t, err)
	assert.Equal(t, 387363, n)
}

func Test_TotalCompletionScore_Input1(t *testing.T) {
	f, err := os.Open("testdata/input01")
	require.NoError(t, err)
	defer f.Close()
	n, err := TotalCompletionScore(f)
	require.NoError(t, err)
	assert.Equal(t, 288957, n)
}

func Test_TotalCompletionScore_Input2(t *testing.T) {
	f, err := os.Open("testdata/input02")
	require.NoError(t, err)
	defer f.Close()
	n, err := TotalCompletionScore(f)
	require.NoError(t, err)
	assert.Equal(t, 4330777059, n)
}
