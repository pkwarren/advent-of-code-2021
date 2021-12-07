package day06

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

func ParseInput(r io.Reader) (*Simulation, error) {
	s := bufio.NewScanner(r)
	sim := Simulation{}
	for s.Scan() {
		l := strings.TrimSpace(s.Text())
		if len(l) == 0 {
			continue
		}
		if len(sim.FishDays) > 0 {
			return nil, fmt.Errorf("input already read")
		}
		fields := strings.Split(l, ",")
		for _, f := range fields {
			i, err := strconv.ParseInt(f, 10, 0)
			if err != nil {
				return nil, err
			}
			sim.AddFish(int(i))
		}
	}
	return &sim, nil
}

type Simulation struct {
	FishDays []int
	index    int
}

func (s *Simulation) AddFish(days int) {
	if days < 0 || days > 8 {
		log.Printf("invalid fish days: %v", days)
	}
	if len(s.FishDays) == 0 {
		s.FishDays = make([]int, 9)
	}
	s.FishDays[days] += 1
}

func (s *Simulation) Advance() {
	toCreate := s.FishDays[s.index]
	s.FishDays[(s.index+7)%len(s.FishDays)] += toCreate
	s.index += 1
	if s.index >= len(s.FishDays) {
		s.index = 0
	}
}

func (s *Simulation) NumFish() int {
	sum := 0
	for _, v := range s.FishDays {
		sum += v
	}
	return sum
}

func (s *Simulation) String() string {
	ordered := append(make([]int, 0, len(s.FishDays)), s.FishDays[s.index:]...)
	ordered = append(ordered, s.FishDays[:s.index]...)
	return fmt.Sprintf("%v", ordered)
}
