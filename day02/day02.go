package day02

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

type Location struct {
	Position int
	Depth    int
}

func (l *Location) Apply(instruction Instruction) {
	switch instruction.Direction {
	case Forward:
		l.Position += instruction.Amount
	case Down:
		l.Depth += instruction.Amount
	case Up:
		l.Depth -= instruction.Amount
	default:
		log.Printf("invalid direction: %v", instruction.Direction)
	}
}

type Direction int

const (
	Forward Direction = iota
	Down
	Up
)

type Instruction struct {
	Direction Direction
	Amount    int
}

func ParseInput(r io.Reader, f func(instruction Instruction) error) error {
	s := bufio.NewScanner(r)
	var inst Instruction
	for s.Scan() {
		l := strings.TrimSpace(s.Text())
		if len(l) == 0 {
			continue
		}
		fields := strings.Split(l, " ")
		if len(fields) != 2 {
			return fmt.Errorf("invalid line: %s", l)
		}
		switch fields[0] {
		case "forward":
			inst.Direction = Forward
		case "down":
			inst.Direction = Down
		case "up":
			inst.Direction = Up
		default:
			return fmt.Errorf("invalid line: %s", l)
		}
		amt, err := strconv.ParseInt(fields[1], 10, 0)
		if err != nil {
			return err
		}
		inst.Amount = int(amt)
		if err := f(inst); err != nil {
			return err
		}
	}
	return s.Err()
}

func CalculateLocation(r io.Reader) (Location, error) {
	var l Location
	err := ParseInput(r, func(inst Instruction) error {
		l.Apply(inst)
		return nil
	})
	if err != nil {
		return Location{}, err
	}
	return l, err
}

type LocationWithAim struct {
	Position int
	Depth    int
	Aim      int
}

func (l *LocationWithAim) Apply(instruction Instruction) {
	switch instruction.Direction {
	case Forward:
		l.Position += instruction.Amount
		l.Depth += l.Aim * instruction.Amount
	case Down:
		l.Aim += instruction.Amount
	case Up:
		l.Aim -= instruction.Amount
	default:
		log.Printf("invalid direction: %v", instruction.Direction)
	}
}

func CalculateLocationWithAim(r io.Reader) (LocationWithAim, error) {
	var l LocationWithAim
	err := ParseInput(r, func(inst Instruction) error {
		l.Apply(inst)
		return nil
	})
	if err != nil {
		return LocationWithAim{}, err
	}
	return l, err
}
