package day04

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_BoardBasics(t *testing.T) {
	b, err := NewBoard([]int{
		22, 13, 17, 11, 0,
		8, 2, 23, 4, 24,
		21, 9, 14, 16, 7,
		6, 10, 3, 18, 5,
		1, 12, 20, 15, 19,
	}, 5)
	require.NoError(t, err)
	for _, v := range []int{11, 4, 16, 18, 15} {
		require.Truef(t, b.Mark(v), "expected value %d to be marked", v)
		if v == 15 {
			require.Truef(t, b.Won(), "expected board to win on 5th number")
		} else {
			require.False(t, b.Won())
		}
	}
	require.Equal(t, 236, b.UnmarkedTotal())
	b.Mark(15)
	require.Equal(t, 236, b.UnmarkedTotal())
	require.False(t, b.Mark(99))
}

func Test_Part1_Input1(t *testing.T) {
	f, err := os.Open("testdata/input01")
	require.NoError(t, err)
	defer f.Close()
	g, err := ParseGame(f)
	require.NoError(t, err)
	assert.Len(t, g.Boards, 3)
	assert.NotEmpty(t, g.Numbers)
won:
	for _, num := range g.Numbers {
		for i, b := range g.Boards {
			if marked := b.Mark(num); marked && b.Won() {
				t.Logf("won! on board %d with %d", i, num)
				require.Equal(t, 24, num)
				require.Equal(t, 188, b.UnmarkedTotal())
				require.Equal(t, 4512, num*b.UnmarkedTotal())
				break won
			}
		}
	}
}

func Test_Part1_Input2(t *testing.T) {
	f, err := os.Open("testdata/input02")
	require.NoError(t, err)
	defer f.Close()
	g, err := ParseGame(f)
	require.NoError(t, err)
	assert.Len(t, g.Boards, 100)
	assert.NotEmpty(t, g.Numbers)
won:
	for _, num := range g.Numbers {
		for i, b := range g.Boards {
			if marked := b.Mark(num); marked && b.Won() {
				t.Logf("won! on board %d with %d", i, num)
				require.Equal(t, 42, num)
				require.Equal(t, 794, b.UnmarkedTotal())
				require.Equal(t, 33348, num*b.UnmarkedTotal())
				break won
			}
		}
	}
}

func Test_Part2_Input1(t *testing.T) {
	f, err := os.Open("testdata/input01")
	require.NoError(t, err)
	defer f.Close()
	g, err := ParseGame(f)
	require.NoError(t, err)
	assert.Len(t, g.Boards, 3)
	assert.NotEmpty(t, g.Numbers)
	lastNum := 0
	var lastWinner *Board
	for _, num := range g.Numbers {
		for i, b := range g.Boards {
			if b.Won() {
				continue
			}
			if marked := b.Mark(num); marked && b.Won() {
				t.Logf("won! on board %d with %d", i, num)
				lastNum = num
				lastWinner = b
			}
		}
	}
	require.NotNil(t, lastWinner)
	require.Equal(t, 13, lastNum)
	require.Equal(t, 148, lastWinner.UnmarkedTotal())
	require.Equal(t, 1924, lastNum*lastWinner.UnmarkedTotal())
}

func Test_Part2_Input2(t *testing.T) {
	f, err := os.Open("testdata/input02")
	require.NoError(t, err)
	defer f.Close()
	g, err := ParseGame(f)
	require.NoError(t, err)
	assert.Len(t, g.Boards, 100)
	assert.NotEmpty(t, g.Numbers)
	lastNum := 0
	var lastWinner *Board
	for _, num := range g.Numbers {
		for i, b := range g.Boards {
			if b.Won() {
				continue
			}
			if marked := b.Mark(num); marked && b.Won() {
				t.Logf("won! on board %d with %d", i, num)
				lastNum = num
				lastWinner = b
			}
		}
	}
	require.NotNil(t, lastWinner)
	assert.Equal(t, 39, lastNum)
	assert.Equal(t, 208, lastWinner.UnmarkedTotal())
	assert.Equal(t, 8112, lastNum*lastWinner.UnmarkedTotal())
}
