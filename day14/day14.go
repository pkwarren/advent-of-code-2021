package day14

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func ParseInput(r io.Reader) (*Instructions, error) {
	var i Instructions
	i.Rules = make(map[[2]rune]rune)
	s := bufio.NewScanner(r)
	for s.Scan() {
		l := strings.TrimSpace(s.Text())
		if len(l) == 0 {
			continue
		}
		if len(i.Template) == 0 {
			i.Template = l
		} else {
			fields := strings.Split(l, " -> ")
			if len(fields) != 2 || len(fields[0]) != 2 || len(fields[1]) != 1 {
				return nil, fmt.Errorf("invalid rule: %s", l)
			}
			r := []rune(fields[0])
			i.Rules[[2]rune{r[0], r[1]}] = []rune(fields[1])[0]
		}
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return &i, nil
}

type Instructions struct {
	Template string
	Rules    map[[2]rune]rune
}

func (i *Instructions) StepMulti(steps int) map[rune]int {
	if steps < 0 || len(i.Template) <= 1 {
		return nil
	}
	pairs := make(map[[2]rune]int)
	r := []rune(i.Template)
	for x := 0; x < len(r)-1; x++ {
		pairs[[2]rune{r[x], r[x+1]}]++
	}
	for x := 0; x < steps; x++ {
		next := make(map[[2]rune]int)
		for k, v := range pairs {
			c := i.Rules[k]
			next[[2]rune{k[0], c}] += v
			next[[2]rune{c, k[1]}] += v
		}
		pairs = next
	}
	m := make(map[rune]int)
	for k, v := range pairs {
		m[k[0]] += v
	}
	m[r[len(r)-1]] += 1 // We need to count the last character from the initial template
	return m
}
