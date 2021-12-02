package day02

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func Test_CalculateLocationPart1_Input1(t *testing.T) {
	f, err := os.Open("testdata/input01")
	require.NoError(t, err)
	defer f.Close()
	l, err := CalculateLocation(f)
	require.NoError(t, err)
	require.Equal(t, Location{Position: 15, Depth: 10}, l)
	require.Equal(t, 150, l.Position*l.Depth)
}

func Test_CalculateLocationPart1_Input2(t *testing.T) {
	f, err := os.Open("testdata/input02")
	require.NoError(t, err)
	defer f.Close()
	l, err := CalculateLocation(f)
	require.NoError(t, err)
	require.Equal(t, Location{Position: 1980, Depth: 866}, l)
	require.Equal(t, 1714680, l.Position*l.Depth)
}

func Test_CalculateLocationPart2_Input1(t *testing.T) {
	f, err := os.Open("testdata/input01")
	require.NoError(t, err)
	defer f.Close()
	l, err := CalculateLocationWithAim(f)
	require.NoError(t, err)
	require.Equal(t, 900, l.Position*l.Depth)
	require.Equal(t, LocationWithAim{Position: 15, Depth: 60, Aim: 10}, l)
}

func Test_CalculateLocationPart2_Input2(t *testing.T) {
	f, err := os.Open("testdata/input02")
	require.NoError(t, err)
	defer f.Close()
	l, err := CalculateLocationWithAim(f)
	require.NoError(t, err)
	require.Equal(t, 1963088820, l.Position*l.Depth)
	require.Equal(t, LocationWithAim{Position: 1980, Depth: 991459, Aim: 866}, l)
}
