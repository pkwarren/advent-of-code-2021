package day18

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

type Tree struct {
	Parent, Left, Right *Tree

	Value int
}

func (t Tree) Leaf() bool {
	return t.Left == nil && t.Right == nil
}

func (t *Tree) Clone() *Tree {
	if t == nil {
		return nil
	}
	return clone(nil, t)
}

func clone(parent, root *Tree) *Tree {
	node := &Tree{Value: root.Value, Parent: parent}
	if root.Left != nil {
		node.Left = clone(node, root.Left)
	}
	if root.Right != nil {
		node.Right = clone(node, root.Right)
	}
	return node
}

func (t *Tree) Split() {
	if !t.Leaf() {
		log.Printf("shouldn't attempt to split leaf node")
		return
	}
	if t.Value < 10 {
		log.Printf("shouldn't attempt to split a leaf node with value < 10")
		return
	}
	left := t.Value / 2
	right := left
	if t.Value%2 == 1 {
		right++
	}
	t.Value = 0
	t.Left = &Tree{Value: left, Parent: t}
	t.Right = &Tree{Value: right, Parent: t}
}

func (t *Tree) ReduceFully() *Tree {
	for t.Reduce() {
		// Perform all possible reductions
	}
	return t
}

func (t *Tree) Reduce() bool {
	// If nested within 4 pairs, explode
	exploded := false
	var right, prev, split *Tree
	t.TraversePreorder(func(node *Tree, depth int) bool {
		if node.Value >= 10 && split == nil {
			split = node
		}
		if right != nil {
			node.Value += right.Value
			return false
		}
		if depth > 4 {
			parent := node.Parent
			if parent.Left.Leaf() && parent.Right.Leaf() {
				exploded = true
				if prev != nil {
					prev.Value += node.Value
				}
				// Remove children from tree
				parent.Left.Parent, parent.Right.Parent = nil, nil
				right = parent.Right
				// Make parent a value node
				parent.Left, parent.Right = nil, nil
				parent.Value = 0
			}
		}
		prev = node
		return true
	})
	if exploded {
		return true
	}
	if split == nil {
		t.TraversePreorder(func(node *Tree, depth int) bool {
			if node.Value >= 10 {
				split = node
				return false
			}
			return true
		})
	}
	if split != nil {
		// We already found a split candidate
		split.Split()
		return true
	}
	return false
}

func (t *Tree) TraversePreorder(f func(node *Tree, depth int) bool) {
	preorderTraverse(t, 0, f)
}

func (t *Tree) Magnitude() int {
	if t.Leaf() {
		return t.Value
	}
	return 3*t.Left.Magnitude() + 2*t.Right.Magnitude()
}

func (t *Tree) Merge(rhs *Tree) *Tree {
	parent := &Tree{Left: t, Right: rhs}
	t.Parent = parent
	rhs.Parent = parent
	return parent
}

func preorderTraverse(root *Tree, depth int, f func(node *Tree, depth int) bool) bool {
	if root == nil {
		return true
	}
	if root.Leaf() {
		// Value node
		if !f(root, depth) {
			return false
		}
	}
	if root.Left != nil && !preorderTraverse(root.Left, depth+1, f) {
		return false
	}
	if root.Right != nil && !preorderTraverse(root.Right, depth+1, f) {
		return false
	}
	return true
}

func ParseInput(r io.Reader) ([]*Tree, error) {
	trees := make([]*Tree, 0)
	s := bufio.NewScanner(r)
	for s.Scan() {
		l := strings.TrimSpace(s.Text())
		if len(l) == 0 {
			continue
		}
		parents := make([]*Tree, 0)
		var root *Tree
		skipToNext := false
		for i, r := range l {
			switch r {
			case '[':
				node := new(Tree)
				if len(parents) == 0 {
					root = node
				} else {
					parent := parents[len(parents)-1]
					node.Parent = parent
					if parent.Left == nil {
						parent.Left = node
					} else if parent.Right == nil {
						parent.Right = node
					} else {
						return nil, fmt.Errorf("failed to parse: %s", l)
					}
				}
				parents = append(parents, node)
			case ']':
				parents = parents[:len(parents)-1]
				skipToNext = false
			case ',':
				skipToNext = false
			default:
				if skipToNext {
					continue
				}
				skipToNext = true
				end := strings.IndexAny(l[i:], ",]")
				if end == -1 {
					return nil, fmt.Errorf("failed to find closing character")
				}
				val, err := strconv.Atoi(l[i : i+end])
				if err != nil {
					return nil, err
				}
				parent := parents[len(parents)-1]
				node := &Tree{Value: val, Parent: parent}
				if parent.Left == nil {
					parent.Left = node
				} else {
					parent.Right = node
				}
			}
		}
		trees = append(trees, root)
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return trees, nil
}
