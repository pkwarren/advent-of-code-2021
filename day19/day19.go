package day19

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func ParseInput(r io.Reader) ([]Scanner, error) {
	var scanners []Scanner
	s := bufio.NewScanner(r)
	for s.Scan() {
		l := strings.TrimSpace(s.Text())
		if len(l) == 0 {
			continue
		}
		if strings.HasPrefix(l, "--- ") {
			scanners = append(scanners, Scanner{Name: strings.TrimSuffix(strings.TrimPrefix(l, "--- "), " ---")})
		} else {
			fields := strings.SplitN(l, ",", 3)
			if len(fields) != 3 {
				return nil, fmt.Errorf("expected 3 fields, found: %d", len(fields))
			}
			coords := make([]int, 0, 3)
			for _, f := range fields {
				v, err := strconv.Atoi(f)
				if err != nil {
					return nil, err
				}
				coords = append(coords, v)
			}
			scanners[len(scanners)-1].Beacons = append(scanners[len(scanners)-1].Beacons, Point3D{X: coords[0], Y: coords[1], Z: coords[2]})
		}
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return scanners, nil
}

type Scanner struct {
	Name    string
	Beacons []Point3D
}

type transformFunc func(x, y, z int) (int, int, int)

// https://www.euclideanspace.com/maths/algebra/matrix/transforms/examples/index.htm
// This was something I've never used before and had to read up on how to build these.
var transforms = []transformFunc{
	func(x, y, z int) (int, int, int) { return x, y, z },
	func(x, y, z int) (int, int, int) { return x, z, -y },
	func(x, y, z int) (int, int, int) { return x, -y, -z },
	func(x, y, z int) (int, int, int) { return x, -z, y },
	func(x, y, z int) (int, int, int) { return -x, -y, z },
	func(x, y, z int) (int, int, int) { return -x, z, y },
	func(x, y, z int) (int, int, int) { return -x, y, -z },
	func(x, y, z int) (int, int, int) { return -x, -z, -y },
	func(x, y, z int) (int, int, int) { return y, z, x },
	func(x, y, z int) (int, int, int) { return y, x, -z },
	func(x, y, z int) (int, int, int) { return y, -z, -x },
	func(x, y, z int) (int, int, int) { return y, -x, z },
	func(x, y, z int) (int, int, int) { return -y, -z, x },
	func(x, y, z int) (int, int, int) { return -y, x, z },
	func(x, y, z int) (int, int, int) { return -y, z, -x },
	func(x, y, z int) (int, int, int) { return -y, -x, -z },
	func(x, y, z int) (int, int, int) { return z, x, y },
	func(x, y, z int) (int, int, int) { return z, y, -x },
	func(x, y, z int) (int, int, int) { return z, -x, -y },
	func(x, y, z int) (int, int, int) { return z, -y, x },
	func(x, y, z int) (int, int, int) { return -z, -x, y },
	func(x, y, z int) (int, int, int) { return -z, y, x },
	func(x, y, z int) (int, int, int) { return -z, x, -y },
	func(x, y, z int) (int, int, int) { return -z, -y, -x },
}

func (s *Scanner) Orientations() []Scanner {
	scanners := make([]Scanner, len(transforms))
	for i, t := range transforms {
		for _, b := range s.Beacons {
			x, y, z := t(b.X, b.Y, b.Z)
			scanners[i].Beacons = append(scanners[i].Beacons, Point3D{X: x, Y: y, Z: z})
		}
	}
	return scanners
}

func BuildTrenchMap(scanners []Scanner) (map[Point3D]struct{}, []Point3D) {
	// We'll use scanner[0] as the reference for all the others
	trenchMap := make(map[Point3D]struct{})
	for _, beacon := range scanners[0].Beacons {
		trenchMap[beacon] = struct{}{}
	}
	scannerLocations := make([]Point3D, 0)
	scannerLocations = append(scannerLocations, Point3D{X: 0, Y: 0, Z: 0})
	remaining := scanners[1:]
	for len(remaining) > 0 {
		var next []Scanner
		for _, scanner := range remaining {
			found := false
		outer:
			for _, orientation := range scanner.Orientations() {
				offsets := make(map[Point3D]int)
				for _, beacon := range orientation.Beacons {
					for existing := range trenchMap {
						offset := existing.Subtract(beacon)
						offsets[offset]++
						if offsets[offset] >= 12 {
							scannerLocations = append(scannerLocations, offset)
							for _, beacon := range orientation.Beacons {
								trenchMap[beacon.Add(offset)] = struct{}{}
							}
							found = true
							break outer
						}
					}
				}
			}
			if !found {
				next = append(next, scanner)
			}
		}
		remaining = next
	}
	return trenchMap, scannerLocations
}

type Point3D struct {
	X, Y, Z int
}

func (p Point3D) Add(o Point3D) Point3D {
	return Point3D{X: p.X + o.X, Y: p.Y + o.Y, Z: p.Z + o.Z}
}

func (p Point3D) Subtract(o Point3D) Point3D {
	return Point3D{X: p.X - o.X, Y: p.Y - o.Y, Z: p.Z - o.Z}
}

func (p Point3D) Distance(o Point3D) int {
	return abs(p.X-o.X) + abs(p.Y-o.Y) + abs(p.Z-o.Z)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
