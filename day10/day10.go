package day10

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

func ParseInput(r io.Reader, f func(string) error) error {
	s := bufio.NewScanner(r)
	for s.Scan() {
		l := strings.TrimSpace(s.Text())
		if len(l) == 0 {
			continue
		}
		if err := f(l); err != nil {
			return err
		}
	}
	if err := s.Err(); err != nil {
		return err
	}
	return nil
}

var openToClose = map[rune]rune{
	'{': '}',
	'(': ')',
	'[': ']',
	'<': '>',
}

func TotalSyntaxErrorScore(r io.Reader) (int, error) {
	total := 0
	scores := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	err := ParseInput(r, func(line string) error {
		stack := make([]rune, 0, len(line))
		for _, c := range line {
			switch c {
			case '{', '(', '[', '<':
				stack = append(stack, openToClose[c])
			case '}', ')', ']', '>':
				if len(stack) == 0 || stack[len(stack)-1] != c {
					total += scores[c]
					return nil
				}
				stack = stack[:len(stack)-1]
			default:
				return fmt.Errorf("invalid character: %v", c)
			}
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	return total, nil
}

func TotalCompletionScore(r io.Reader) (int, error) {
	var totals []int
	scores := map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}
	err := ParseInput(r, func(line string) error {
		stack := make([]rune, 0, len(line))
		for _, c := range line {
			switch c {
			case '{', '(', '[', '<':
				stack = append(stack, openToClose[c])
			case '}', ')', ']', '>':
				if len(stack) == 0 || stack[len(stack)-1] != c {
					return nil
				}
				stack = stack[:len(stack)-1]
			default:
				return fmt.Errorf("invalid character: %v", c)
			}
		}
		total := 0
		for i := len(stack) - 1; i >= 0; i-- {
			total *= 5
			total += scores[stack[i]]
		}
		totals = append(totals, total)
		return nil
	})
	if err != nil {
		return 0, err
	}
	sort.Ints(totals)
	return totals[len(totals)/2], nil
}
