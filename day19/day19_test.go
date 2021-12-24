package day19

import (
	"math"
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
	scanners, err := ParseInput(f)
	require.NoError(t, err)
	require.Len(t, scanners, 5)
}

func Test_Orientations(t *testing.T) {
	scanners, err := ParseInput(strings.NewReader(`
	--- scanner 0 ---
	-1,-1,1
	-2,-2,2
	-3,-3,3
	-2,-3,1
	5,6,-4
	8,0,7
	`))
	require.NoError(t, err)
	require.Len(t, scanners, 1)
	beaconCombos := make([][]Point3D, 0)
	for _, r := range scanners[0].Orientations() {
		beaconCombos = append(beaconCombos, r.Beacons)
	}
	assert.Len(t, beaconCombos, 24)
	assert.Contains(t, beaconCombos, []Point3D{
		{-1, -1, 1},
		{-2, -2, 2},
		{-3, -3, 3},
		{-2, -3, 1},
		{5, 6, -4},
		{8, 0, 7},
	})
	assert.Contains(t, beaconCombos, []Point3D{
		{1, -1, 1},
		{2, -2, 2},
		{3, -3, 3},
		{2, -1, 3},
		{-5, 4, -6},
		{-8, -7, 0},
	})
	assert.Contains(t, beaconCombos, []Point3D{
		{-1, -1, -1},
		{-2, -2, -2},
		{-3, -3, -3},
		{-1, -3, -2},
		{4, 6, 5},
		{-7, 0, 8},
	})
	assert.Contains(t, beaconCombos, []Point3D{
		{1, 1, -1},
		{2, 2, -2},
		{3, 3, -3},
		{1, 3, -2},
		{-4, -6, 5},
		{7, 0, 8},
	})
	assert.Contains(t, beaconCombos, []Point3D{
		{1, 1, 1},
		{2, 2, 2},
		{3, 3, 3},
		{3, 1, 2},
		{-6, -4, -5},
		{0, 7, -8},
	})
}

func Test_BuildTrenchMap_Part1_Input01(t *testing.T) {
	f, err := os.Open("testdata/input01")
	require.NoError(t, err)
	defer f.Close()
	scanners, err := ParseInput(f)
	require.NoError(t, err)
	m, _ := BuildTrenchMap(scanners)
	assert.Len(t, m, 79)
}

func Test_BuildTrenchMap_Part1_Input02(t *testing.T) {
	f, err := os.Open("testdata/input02")
	require.NoError(t, err)
	defer f.Close()
	scanners, err := ParseInput(f)
	require.NoError(t, err)
	m, _ := BuildTrenchMap(scanners)
	assert.Len(t, m, 306)
}

func Test_BuildTrenchMap_Part2_Input01(t *testing.T) {
	f, err := os.Open("testdata/input01")
	require.NoError(t, err)
	defer f.Close()
	scanners, err := ParseInput(f)
	require.NoError(t, err)
	_, locations := BuildTrenchMap(scanners)
	maxDistance := math.MinInt
	for i := range locations {
		for j := range locations {
			if i != j {
				distance := locations[i].Distance(locations[j])
				if distance > maxDistance {
					maxDistance = distance
				}
			}
		}
	}
	assert.Equal(t, 3621, maxDistance)
}

func Test_BuildTrenchMap_Part2_Input02(t *testing.T) {
	f, err := os.Open("testdata/input02")
	require.NoError(t, err)
	defer f.Close()
	scanners, err := ParseInput(f)
	require.NoError(t, err)
	_, locations := BuildTrenchMap(scanners)
	maxDistance := math.MinInt
	for i := range locations {
		for j := range locations {
			if i != j {
				distance := locations[i].Distance(locations[j])
				if distance > maxDistance {
					maxDistance = distance
				}
			}
		}
	}
	assert.Equal(t, 9764, maxDistance)
}
