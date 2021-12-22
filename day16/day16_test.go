package day16

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ParseInput(t *testing.T) {
	f, err := os.Open("testdata/input01")
	require.NoError(t, err)
	b, err := ParseInput(f)
	require.NoError(t, err)
	assert.Equal(t, []byte{0xD, 0x2, 0xF, 0xE, 0x2, 0x8}, b)
}

func Test_Part1_Input1(t *testing.T) {
	f, err := os.Open("testdata/input01")
	require.NoError(t, err)
	b, err := ParseInput(f)
	require.NoError(t, err)

	p, err := ParsePacket(b)
	require.NoError(t, err)
	assert.IsType(t, &LiteralPacket{}, p)
	assert.Equal(t, uint8(6), p.Header().Version)
	assert.Equal(t, Literal, p.Header().TypeID)
	assert.Equal(t, 2021, p.Value())
}

func Test_Part1_Input2(t *testing.T) {
	b, err := ParseInput(strings.NewReader("38006F45291200"))
	require.NoError(t, err)

	p, err := ParsePacket(b)
	require.NoError(t, err)
	assert.IsType(t, &OperatorPacket{}, p)
	opPacket := p.(*OperatorPacket)
	require.Len(t, opPacket.SubPackets(), 2)
	assert.IsType(t, &LiteralPacket{}, opPacket.SubPackets()[0])
	assert.IsType(t, &LiteralPacket{}, opPacket.SubPackets()[1])
}

func Test_Part1_Input3(t *testing.T) {
	b, err := ParseInput(strings.NewReader("EE00D40C823060"))
	require.NoError(t, err)

	p, err := ParsePacket(b)
	require.NoError(t, err)
	assert.IsType(t, &OperatorPacket{}, p)
	opPacket := p.(*OperatorPacket)
	subPackets := opPacket.SubPackets()
	require.Len(t, subPackets, 3)
	assert.IsType(t, &LiteralPacket{}, subPackets[0])
	assert.IsType(t, &LiteralPacket{}, subPackets[1])
	assert.IsType(t, &LiteralPacket{}, subPackets[2])
	assert.Equal(t, 3, p.Value())
}

func Test_Part1_Input4(t *testing.T) {
	b, err := ParseInput(strings.NewReader("8A004A801A8002F478"))
	require.NoError(t, err)

	p, err := ParsePacket(b)
	require.NoError(t, err)
	subPackets := p.SubPackets()
	require.Len(t, subPackets, 1)
	subPackets = subPackets[0].SubPackets()
	require.Len(t, subPackets, 1)
	subPackets = subPackets[0].SubPackets()
	require.Len(t, subPackets, 1)
	assert.IsType(t, &LiteralPacket{}, subPackets[0])
}

func Test_Part1_Input_Sum_Part1(t *testing.T) {
	expected := []struct {
		hex string
		sum int
	}{
		{"8A004A801A8002F478", 16},
		{"620080001611562C8802118E34", 12},
		{"C0015000016115A2E0802F182340", 23},
		{"A0016C880162017C3686B18A3D4780", 31},
	}
	for _, test := range expected {
		assert.Equal(t, test.sum, versionSum(t, strings.NewReader(test.hex)))
	}
}

func Test_Part1_Input_Sum_Part2(t *testing.T) {
	f, err := os.Open("testdata/input02")
	require.NoError(t, err)
	defer f.Close()
	assert.Equal(t, 949, versionSum(t, f))
}

func Test_Part2_Value_Part1(t *testing.T) {
	expected := []struct {
		hex   string
		value int
	}{
		{"C200B40A82", 3},
		{"04005AC33890", 54},
		{"880086C3E88112", 7},
		{"CE00C43D881120", 9},
		{"D8005AC2A8F0", 1},
		{"F600BC2D8F", 0},
		{"9C005AC2F8F0", 0},
		{"9C0141080250320F1802104A08", 1},
	}
	for _, test := range expected {
		assert.Equal(t, test.value, value(t, strings.NewReader(test.hex)))
	}
}

func Test_Part2_Value_Part2(t *testing.T) {
	f, err := os.Open("testdata/input02")
	require.NoError(t, err)
	defer f.Close()
	assert.Equal(t, 1114600142730, value(t, f))
}

func value(t *testing.T, hex io.Reader) int {
	t.Helper()
	b, err := ParseInput(hex)
	require.NoError(t, err)

	p, err := ParsePacket(b)
	require.NoError(t, err)
	return p.Value()
}

func versionSum(t *testing.T, hex io.Reader) int {
	t.Helper()
	sum := 0
	b, err := ParseInput(hex)
	require.NoError(t, err)

	p, err := ParsePacket(b)
	require.NoError(t, err)
	sum = int(p.Header().Version)
	remaining := p.SubPackets()
	for len(remaining) > 0 {
		next := make([]Packet, 0)
		for _, p := range remaining {
			sum += int(p.Header().Version)
			next = append(next, p.SubPackets()...)
		}
		remaining = next
	}
	return sum
}
