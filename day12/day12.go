package day12

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func ParseInput(r io.Reader) (*CaveMap, error) {
	var m CaveMap
	s := bufio.NewScanner(r)
	for s.Scan() {
		l := strings.TrimSpace(s.Text())
		if len(l) == 0 {
			continue
		}
		fields := strings.Split(l, "-")
		if len(fields) != 2 {
			return nil, fmt.Errorf("expected 2 fields, found: %v", len(fields))
		}
		m.AddEdge(Cave(fields[0]), Cave(fields[1]))
		m.AddEdge(Cave(fields[1]), Cave(fields[0]))
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return &m, nil
}

type CaveMap struct {
	Nodes map[Cave][]Cave
}

type Cave string

const CaveStart = Cave("start")
const CaveEnd = Cave("end")

type Path []Cave

func (c Cave) IsBig() bool {
	return strings.ToUpper(string(c)) == string(c)
}

func (c *CaveMap) AddEdge(from, to Cave) {
	if from == to {
		return
	}
	exists := false
	for _, n := range c.Nodes[from] {
		if n == to {
			exists = true
			break
		}
	}
	if !exists {
		if c.Nodes == nil {
			c.Nodes = make(map[Cave][]Cave)
		}
		c.Nodes[from] = append(c.Nodes[from], to)
	}
}

func (c *CaveMap) CountDistinctPaths(exploreSmallCavesTwice bool) int {
	numPaths := 0
	c.explore(CaveStart, nil, &numPaths, exploreSmallCavesTwice)
	return numPaths
}

func (c *CaveMap) explore(cave Cave, visited map[Cave]struct{}, numPaths *int, exploreSmallCavesTwice bool) {
	if cave == CaveEnd {
		// Finished exploring
		*numPaths += 1
		return
	}

	nextVisited := visited
	if !cave.IsBig() {
		if _, ok := nextVisited[cave]; ok {
			exploreSmallCavesTwice = false
		} else {
			nextVisited = copyMap(visited)
			nextVisited[cave] = struct{}{}
		}
	}
	nextCaves := c.Nodes[cave]
	for _, next := range nextCaves {
		if next == CaveStart {
			// Can't visit start multiple times
			continue
		}
		if _, ok := nextVisited[next]; ok && !exploreSmallCavesTwice {
			// We've already visited this small cave
			continue
		}
		c.explore(next, nextVisited, numPaths, exploreSmallCavesTwice)
	}
}

func copyMap(original map[Cave]struct{}) map[Cave]struct{} {
	m := make(map[Cave]struct{}, len(original)+1)
	for k := range original {
		m[k] = struct{}{}
	}
	return m
}
