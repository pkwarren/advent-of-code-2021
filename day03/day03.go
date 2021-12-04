package day03

import (
	"bufio"
	"io"
	"math/bits"
	"strconv"
	"strings"
)

func ParseInput(r io.Reader, f func(i uint, len int) error) error {
	s := bufio.NewScanner(r)
	for s.Scan() {
		l := strings.TrimSpace(s.Text())
		if len(l) > 0 {
			i, err := strconv.ParseUint(l, 2, 0)
			if err != nil {
				return err
			}
			if err := f(uint(i), len(l)); err != nil {
				return err
			}
		}
	}
	return s.Err()
}

func CalculatePowerConsumption(r io.Reader) (int, error) {
	numOnes := make([]int, bits.UintSize)
	numZeros := make([]int, bits.UintSize)
	err := ParseInput(r, func(val uint, len int) error {
		for i := 0; i < len; i++ {
			if (val & (1 << uint(i))) != 0 {
				numOnes[i] = numOnes[i] + 1
			} else {
				numZeros[i] = numZeros[i] + 1
			}
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	gamma, epsilon := 0, 0
	for i := 0; i < bits.UintSize; i++ {
		ones, zeros := numOnes[i], numZeros[i]
		if ones == 0 && zeros == 0 {
			// No bits counted at this precision - skip
			continue
		}
		if ones > zeros {
			gamma |= 1 << i
		} else {
			epsilon |= 1 << i
		}
	}
	return gamma * epsilon, nil
}
