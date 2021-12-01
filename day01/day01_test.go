package day01

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"testing"
)

func loadFile(t *testing.T, file string) io.Reader {
	t.Helper()
	f, err := os.Open("testdata/" + file)
	if err != nil {
		t.Fatalf("failed to open file: %v", err)
	}
	return f
}

func Test_NumLargerMeasurements_Input01(t *testing.T) {
	n, err := NumLargerMeasurements(loadFile(t, "input01"))
	require.NoError(t, err)
	assert.Equal(t, 7, n)
}

func Test_NumLargerMeasurements_Input02(t *testing.T) {
	n, err := NumLargerMeasurements(loadFile(t, "input02"))
	require.NoError(t, err)
	assert.Equal(t, 1766, n)
}

func Test_NumLargerMeasurementsSlidingWindow_Input01(t *testing.T) {
	n, err := NumLargerMeasurementsSlidingWindow(loadFile(t, "input01"), 3)
	require.NoError(t, err)
	assert.Equal(t, 5, n)
}

func Test_NumLargerMeasurementsSlidingWindow_Input02(t *testing.T) {
	n, err := NumLargerMeasurementsSlidingWindow(loadFile(t, "input02"), 3)
	require.NoError(t, err)
	assert.Equal(t, 1797, n)
}
