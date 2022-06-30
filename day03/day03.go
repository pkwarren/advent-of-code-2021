package day03

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"math/bits"
	"sort"
	"strconv"
	"strings"
)

func ParseInput(r io.Reader, f func(i int, len int) error) error {
	s := bufio.NewScanner(r)
	for s.Scan() {
		l := strings.TrimSpace(s.Text())
		if len(l) > 0 {
			i, err := strconv.ParseUint(l, 2, 0)
			if err != nil {
				return err
			}
			if i > math.MaxInt {
				return fmt.Errorf("value out of range: %v", i)
			}
			if err := f(int(i), len(l)); err != nil {
				return err
			}
		}
	}
	return s.Err()
}

func CalculatePowerConsumption(r io.Reader) (int, error) {
	numOnes := make([]int, bits.UintSize)
	numZeros := make([]int, bits.UintSize)
	maxPrecision := 0
	err := ParseInput(r, func(val int, len int) error {
		if len > maxPrecision {
			maxPrecision = len
		}
		for i := 0; i < len; i++ {
			if (val & (1 << i)) != 0 {
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
	for i := 0; i < maxPrecision; i++ {
		ones, zeros := numOnes[i], numZeros[i]
		if ones > zeros {
			gamma |= 1 << i
		} else {
			epsilon |= 1 << i
		}
	}
	return gamma * epsilon, nil
}

func CalculateLifeSupportRating(r io.Reader) (int, error) {
	values := make([]int, 0)
	maxPrecision := 0
	err := ParseInput(r, func(val int, len int) error {
		if len > maxPrecision {
			maxPrecision = len
		}
		values = append(values, val)
		return nil
	})
	if err != nil {
		return 0, err
	}
	sort.Ints(values)
	o2 := calculateO2Rating(values, maxPrecision)
	co2 := calculateCO2Rating(values, maxPrecision)
	return o2 * co2, nil
}

func calculateO2Rating(values []int, maxPrecision int) int {
	o2min, o2max := 0, int(math.Pow(2, float64(maxPrecision))-1)
	for i := maxPrecision - 1; i >= 0; i -= 1 {
		ones, zeros := countOnesZeros(values, i)
		if ones >= zeros {
			o2min |= 1 << i
			lowerBound := sort.Search(len(values), func(i int) bool {
				return values[i] >= o2min
			})
			values = values[lowerBound:]
		} else {
			o2max &= ^(1 << i)
			upperBound := sort.Search(len(values), func(i int) bool {
				return values[i] > o2max
			})
			values = values[:upperBound]
		}
		fmt.Printf("%b (o2min) %b (o2max) values %v\n", o2min, o2max, valuesBinary(values, maxPrecision))
		if len(values) == 1 {
			return values[0]
		}
	}
	return 0
}

func calculateCO2Rating(values []int, maxPrecision int) int {
	co2min, co2max := 0, int(math.Pow(2, float64(maxPrecision))-1)
	fmt.Printf("%b (co2min) %b (co2max) values %v\n", co2min, co2max, valuesBinary(values, maxPrecision))
	for i := maxPrecision - 1; i >= 0; i -= 1 {
		ones, zeros := countOnesZeros(values, i)
		if zeros <= ones {
			co2max &= ^(1 << i)
			upperBound := sort.Search(len(values), func(i int) bool {
				return values[i] > co2max
			})
			values = values[:upperBound]
		} else {
			co2min |= 1 << i
			lowerBound := sort.Search(len(values), func(i int) bool {
				return values[i] >= co2min
			})
			values = values[lowerBound:]
		}
		fmt.Printf("%b (co2min) %b (co2max) values %v\n", co2min, co2max, valuesBinary(values, maxPrecision))
		if len(values) == 1 {
			return values[0]
		}
	}
	return 0
}

func countOnesZeros(values []int, place int) (int, int) {
	ones, zeros := 0, 0
	for _, val := range values {
		if val&(1<<place) != 0 {
			ones += 1
		} else {
			zeros += 1
		}
	}
	return ones, zeros
}

func valuesBinary(values []int, maxPrecision int) string {
	var sb strings.Builder
	sb.WriteByte('[')
	for i, v := range values {
		b := strconv.FormatInt(int64(v), 2)
		if len(b) < maxPrecision {
			sb.WriteString(strings.Repeat("0", maxPrecision-len(b)))
		}
		sb.WriteString(b)
		if i < len(values)-1 {
			sb.WriteByte(' ')
		}
	}
	sb.WriteByte(']')
	return sb.String()
}
