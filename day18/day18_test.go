package day18

import (
	"math"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ParseInput(t *testing.T) {
	trees, err := ParseInput(strings.NewReader("[[[9,[3,8]],[[0,9],6]],[[[3,7],[4,9]],3]]"))
	require.NoError(t, err)
	require.Len(t, trees, 1)
	root := trees[0]
	preorderTraverse(root, 0, func(node *Tree, depth int) bool {
		return true
	})
}

type ValueAndDepth struct {
	Value, Depth int
}

func getValueAndDepths(t *testing.T, root *Tree) []ValueAndDepth {
	t.Helper()
	values := make([]ValueAndDepth, 0)
	root.TraversePreorder(func(t *Tree, depth int) bool {
		values = append(values, ValueAndDepth{Value: t.Value, Depth: depth})
		return true
	})
	return values
}

func Test_Reduce_Explode_Example1(t *testing.T) {
	trees, err := ParseInput(strings.NewReader("[[[[[9,8],1],2],3],4]"))
	require.NoError(t, err)
	require.Len(t, trees, 1)
	root := trees[0]
	before := getValueAndDepths(t, root)
	assert.Equal(t, []ValueAndDepth{
		{Value: 9, Depth: 5},
		{Value: 8, Depth: 5},
		{Value: 1, Depth: 4},
		{Value: 2, Depth: 3},
		{Value: 3, Depth: 2},
		{Value: 4, Depth: 1},
	}, before)

	assert.True(t, root.Reduce())
	after := getValueAndDepths(t, root)
	assert.Equal(t, []ValueAndDepth{
		{Value: 0, Depth: 4},
		{Value: 9, Depth: 4},
		{Value: 2, Depth: 3},
		{Value: 3, Depth: 2},
		{Value: 4, Depth: 1},
	}, after)
}

func Test_Reduce_Explode_Example2(t *testing.T) {
	trees, err := ParseInput(strings.NewReader("[7,[6,[5,[4,[3,2]]]]]"))
	require.NoError(t, err)
	require.Len(t, trees, 1)
	root := trees[0]
	assert.True(t, root.Reduce())
	after := getValueAndDepths(t, root)
	assert.Equal(t, []ValueAndDepth{
		{Value: 7, Depth: 1},
		{Value: 6, Depth: 2},
		{Value: 5, Depth: 3},
		{Value: 7, Depth: 4},
		{Value: 0, Depth: 4},
	}, after)
}

func Test_Reduce_Explode_Example3(t *testing.T) {
	trees, err := ParseInput(strings.NewReader("[[6,[5,[4,[3,2]]]],1]"))
	require.NoError(t, err)
	require.Len(t, trees, 1)
	root := trees[0]
	assert.True(t, root.Reduce())
	after := getValueAndDepths(t, root)
	assert.Equal(t, []ValueAndDepth{
		{Value: 6, Depth: 2},
		{Value: 5, Depth: 3},
		{Value: 7, Depth: 4},
		{Value: 0, Depth: 4},
		{Value: 3, Depth: 1},
	}, after)
}

func Test_Reduce_Complete_Example1(t *testing.T) {
	input := `[[[[4,3],4],4],[7,[[8,4],9]]]
[1,1]`
	trees, err := ParseInput(strings.NewReader(input))
	require.NoError(t, err)
	require.Len(t, trees, 2)
	root := trees[0].Merge(trees[1])
	t.Log("before:", getValueAndDepths(t, root))
	for root.Reduce() {
		t.Log("reduced:", getValueAndDepths(t, root))
	}
	assert.Equal(t, []ValueAndDepth{
		{Value: 0, Depth: 4},
		{Value: 7, Depth: 4},
		{Value: 4, Depth: 3},
		{Value: 7, Depth: 4},
		{Value: 8, Depth: 4},
		{Value: 6, Depth: 4},
		{Value: 0, Depth: 4},
		{Value: 8, Depth: 2},
		{Value: 1, Depth: 2},
	}, getValueAndDepths(t, root))
}

func Test_TreeMagnitude(t *testing.T) {
	tests := []struct {
		tree      string
		magnitude int
	}{
		{"[[1,2],[[3,4],5]]", 143},
		{"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", 1384},
		{"[[[[1,1],[2,2]],[3,3]],[4,4]]", 445},
		{"[[[[3,0],[5,3]],[4,4]],[5,5]]", 791},
		{"[[[[5,0],[7,4]],[5,5]],[6,6]]", 1137},
		{"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", 3488},
	}
	for _, test := range tests {
		trees, err := ParseInput(strings.NewReader(test.tree))
		require.NoError(t, err)
		require.Len(t, trees, 1)
		root := trees[0]
		assert.Equal(t, test.magnitude, root.Magnitude())
	}
}

func Test_TreeMagnitude_Part1_Input01(t *testing.T) {
	f, err := os.Open("testdata/input01")
	require.NoError(t, err)
	defer f.Close()
	trees, err := ParseInput(f)
	require.NoError(t, err)
	require.Len(t, trees, 10)
	var root *Tree
	for _, tree := range trees {
		if root == nil {
			root = tree
		} else {
			root = root.Merge(tree)
		}
		for root.Reduce() {
			// reducing
		}
	}
	assert.Equal(t, 4140, root.Magnitude())
	assert.Equal(t, []ValueAndDepth{
		{Value: 6, Depth: 4},
		{Value: 6, Depth: 4},
		{Value: 7, Depth: 4},
		{Value: 6, Depth: 4},
		{Value: 7, Depth: 4},
		{Value: 7, Depth: 4},
		{Value: 7, Depth: 4},
		{Value: 0, Depth: 4},
		{Value: 7, Depth: 4},
		{Value: 7, Depth: 4},
		{Value: 7, Depth: 4},
		{Value: 7, Depth: 4},
		{Value: 7, Depth: 4},
		{Value: 8, Depth: 4},
		{Value: 9, Depth: 4},
		{Value: 9, Depth: 4},
	}, getValueAndDepths(t, root))
}

func Test_TreeMagnitude_Part1_Input02(t *testing.T) {
	f, err := os.Open("testdata/input02")
	require.NoError(t, err)
	defer f.Close()
	trees, err := ParseInput(f)
	require.NoError(t, err)
	var root *Tree
	for _, tree := range trees {
		if root == nil {
			root = tree
		} else {
			root = root.Merge(tree)
		}
		for root.Reduce() {
			// reducing
		}
	}
	assert.Equal(t, 3869, root.Magnitude())
}

func Test_TreeMagnitude_Part2_Input01(t *testing.T) {
	f, err := os.Open("testdata/input01")
	require.NoError(t, err)
	defer f.Close()
	trees, err := ParseInput(f)
	require.NoError(t, err)
	require.Len(t, trees, 10)
	maxMagnitude := math.MinInt
	for i, left := range trees {
		for j, right := range trees {
			if i != j {
				l2r := left.Clone().Merge(right.Clone()).ReduceFully()
				l2rMagnitude := l2r.Magnitude()
				if l2rMagnitude > maxMagnitude {
					maxMagnitude = l2rMagnitude
				}
				r2l := right.Clone().Merge(left.Clone()).ReduceFully()
				r2lMagnitude := r2l.Magnitude()
				if r2lMagnitude > maxMagnitude {
					maxMagnitude = r2lMagnitude
				}
			}
		}
	}
	assert.Equal(t, 3993, maxMagnitude)
}

func Test_TreeMagnitude_Part2_Input02(t *testing.T) {
	f, err := os.Open("testdata/input02")
	require.NoError(t, err)
	defer f.Close()
	trees, err := ParseInput(f)
	require.NoError(t, err)
	maxMagnitude := math.MinInt
	for i, left := range trees {
		for j, right := range trees {
			if i != j {
				l2r := left.Clone().Merge(right.Clone()).ReduceFully()
				l2rMagnitude := l2r.Magnitude()
				if l2rMagnitude > maxMagnitude {
					maxMagnitude = l2rMagnitude
				}
				r2l := right.Clone().Merge(left.Clone()).ReduceFully()
				r2lMagnitude := r2l.Magnitude()
				if r2lMagnitude > maxMagnitude {
					maxMagnitude = r2lMagnitude
				}
			}
		}
	}
	assert.Equal(t, 4671, maxMagnitude)
}
