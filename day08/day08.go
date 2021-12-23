package day08

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"sort"
	"strings"
)

func ParseInput(r io.Reader, f func(signals, outputs []string) error) error {
	s := bufio.NewScanner(r)
	for s.Scan() {
		l := strings.TrimSpace(s.Text())
		if len(l) == 0 {
			continue
		}
		fields := strings.Split(l, "|")
		if len(fields) != 2 {
			return fmt.Errorf("expected signals and outputs on %s", l)
		}
		signals, outputs := strings.Fields(strings.TrimSpace(fields[0])), strings.Fields(strings.TrimSpace(fields[1]))
		if len(signals) != 10 {
			return fmt.Errorf("expected 10 signals per line, found: %v", len(signals))
		}
		if len(outputs) != 4 {
			return fmt.Errorf("expected 4 output digits per line, found: %v", len(outputs))
		}
		if err := f(signals, outputs); err != nil {
			return err
		}
	}
	if err := s.Err(); err != nil {
		return err
	}
	return nil
}

var digitToWires = map[int]string{
	0: "abcefg",
	1: "cf",
	2: "acdeg",
	3: "acdfg",
	4: "bcdf",
	5: "abdfg",
	6: "abdefg",
	7: "acf",
	8: "abcdefg",
	9: "abcdfg",
}

func FindUniqueSegments(r io.Reader) (int, error) {
	lenToCount := make(map[int]int)
	for _, wire := range digitToWires {
		lenToCount[len(wire)] += 1
	}
	unique := 0
	err := ParseInput(r, func(signals, outputs []string) error {
		for _, o := range outputs {
			if lenToCount[len(o)] == 1 {
				unique += 1
			}
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	return unique, nil
}

func DecodeLines(r io.Reader) ([]int, error) {
	decoded := make([]int, 0)
	err := ParseInput(r, func(signals, outputs []string) error {
		decoded = append(decoded, decode(signals, outputs))
		return nil
	})
	if err != nil {
		return nil, err
	}
	return decoded, nil
}

func decode(signals, outputs []string) int {
	num := 0
	vals := make([]string, 10)
	remaining := make([]string, 0, len(signals))
	for _, s := range signals {
		s = sortChars(s)
		switch len(s) {
		case 2:
			// 1 = 'cf' (only one with len 2)
			vals[1] = s
		case 3:
			// 7 = 'acf' (only one with len 3)
			vals[7] = s
		case 4:
			// 4 = 'bcdf' (only one with len 4)
			vals[4] = s
		case 7:
			// 8 = 'abcdefg' (only one with len 7)
			vals[8] = s
		default:
			remaining = append(remaining, s)
		}
	}
	for _, s := range remaining {
		switch len(s) {
		case 5:
			if charsInCommon(s, vals[1]) == 2 {
				// 3 ('acdfg') contains all letters from 1 ('cf')
				vals[3] = s
			} else if charsInCommon(s, vals[4]) == 2 {
				// 2 ('acdeg') is missing two characters from 4 ('bcdf')
				vals[2] = s
			} else {
				// 5 ('abdfg') is last remaining one with len 5
				vals[5] = s
			}
		case 6:
			if charsInCommon(s, vals[4]) == 4 {
				// 9 ('abcdfg') contains all letters from 4 ('bcdf')
				vals[9] = s
			} else if charsInCommon(s, vals[1]) == 2 {
				// 0 ('abcefg') contains all letters from 1 ('cf')
				vals[0] = s
			} else {
				// 6 ('abdefg') is last remaining one with len 6
				vals[6] = s
			}
		default:
			log.Printf("unexpected length: %v", len(s))
		}
	}
	wordToValue := make(map[string]int, len(vals))
	for i, v := range vals {
		wordToValue[v] = i
	}
	for _, o := range outputs {
		digit := wordToValue[sortChars(o)]
		num = num*10 + digit
	}
	return num
}

func charsInCommon(s1, s2 string) int {
	common := 0
	shorter, longer := s1, s2
	if len(s2) < len(s1) {
		shorter, longer = s2, s1
	}
	for _, c := range shorter {
		if strings.ContainsRune(longer, c) {
			common += 1
		}
	}
	return common
}

func sortChars(s string) string {
	chars := make([]rune, 0, len(s))
	for _, c := range s {
		chars = append(chars, c)
	}
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})
	return string(chars)
}
