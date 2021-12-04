package day03

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func Test_CalculatePowerConsumption_Part1_Input1(t *testing.T) {
	f, err := os.Open("testdata/input01")
	require.NoError(t, err)
	consumption, err := CalculatePowerConsumption(f)
	require.NoError(t, err)
	require.Equal(t, 198, consumption)
}

func Test_CalculatePowerConsumption_Part1_Input2(t *testing.T) {
	f, err := os.Open("testdata/input02")
	require.NoError(t, err)
	consumption, err := CalculatePowerConsumption(f)
	require.NoError(t, err)
	require.Equal(t, 3959450, consumption)
}

func Test_CalculateLifeSupportRating_Part2_Input1(t *testing.T) {
	f, err := os.Open("testdata/input01")
	require.NoError(t, err)
	consumption, err := CalculateLifeSupportRating(f)
	require.NoError(t, err)
	require.Equal(t, 230, consumption)
}

func Test_CalculateLifeSupportRating_Part2_Input2(t *testing.T) {
	f, err := os.Open("testdata/input02")
	require.NoError(t, err)
	consumption, err := CalculateLifeSupportRating(f)
	require.NoError(t, err)
	require.Equal(t, 7440311, consumption)
}
