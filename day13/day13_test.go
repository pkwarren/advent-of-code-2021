package day13

import (
	"fmt"
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
	fmt.Println(m)
}

func Test_Fold_Part1_Input1(t *testing.T) {
	f, err := os.Open("testdata/input01")
	require.NoError(t, err)
	defer f.Close()
	m, err := ParseInput(f)
	require.NoError(t, err)
	t.Log(m)
	assert.Equal(t, 18, m.NumDots())
	m.FoldNext()
	t.Log(m)
	assert.Equal(t, 17, m.NumDots())
	m.FoldNext()
	t.Log(m)
	assert.Equal(t, 16, m.NumDots())
}

func Test_Fold_Part1_Input2(t *testing.T) {
	f, err := os.Open("testdata/input02")
	require.NoError(t, err)
	defer f.Close()
	m, err := ParseInput(f)
	require.NoError(t, err)
	assert.Equal(t, 796, m.NumDots())
	m.FoldNext()
	assert.Equal(t, 666, m.NumDots())
}

func Test_Fold_Part2_Input2(t *testing.T) {
	f, err := os.Open("testdata/input02")
	require.NoError(t, err)
	defer f.Close()
	m, err := ParseInput(f)
	require.NoError(t, err)
	for i := 0; i < len(m.Folds); i++ {
		m.FoldNext()
	}
	assert.Equal(t, 97, m.NumDots())
	fmt.Println(m)
}
