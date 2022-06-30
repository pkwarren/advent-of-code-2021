package day16

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math"
	"math/bits"
	"strings"
)

type Operator uint

const (
	Sum Operator = iota
	Product
	Minimum
	Maximum
	Literal
	GreaterThan
	LessThan
	EqualTo
)

type Header struct {
	Version uint8    // first three bits
	TypeID  Operator // second three bits
}

type Packet interface {
	Header() Header
	Value() int
	SubPackets() []Packet
}

type LiteralPacket struct {
	h     Header
	value int
}

func (l *LiteralPacket) Header() Header {
	return l.h
}

func (l *LiteralPacket) Value() int {
	return l.value
}

func (l *LiteralPacket) SubPackets() []Packet {
	return nil
}

type OperatorPacket struct {
	h          Header
	subPackets []Packet
}

func (o *OperatorPacket) Header() Header {
	return o.h
}

func (o *OperatorPacket) SubPackets() []Packet {
	return o.subPackets
}

func (o *OperatorPacket) Value() int {
	switch o.h.TypeID {
	case Sum:
		sum := 0
		for _, p := range o.subPackets {
			sum += p.Value()
		}
		return sum
	case Product:
		product := 0
		for i, p := range o.subPackets {
			if i == 0 {
				product = p.Value()
			} else {
				product *= p.Value()
			}
		}
		return product
	case Minimum:
		min := math.MaxInt
		for _, p := range o.subPackets {
			val := p.Value()
			if val < min {
				min = val
			}
		}
		return min
	case Maximum:
		max := math.MinInt
		for _, p := range o.subPackets {
			val := p.Value()
			if val > max {
				max = val
			}
		}
		return max
	case GreaterThan:
		if o.subPackets[0].Value() > o.subPackets[1].Value() {
			return 1
		}
		return 0
	case LessThan:
		if o.subPackets[0].Value() < o.subPackets[1].Value() {
			return 1
		}
		return 0
	case EqualTo:
		if o.subPackets[0].Value() == o.subPackets[1].Value() {
			return 1
		}
		return 0
	default:
		return 0
	}
}

type BitReader struct {
	r        io.ByteScanner
	offset   uint8
	position uint
}

func NewBitReader(r io.ByteScanner) *BitReader {
	return &BitReader{r: r}
}

func (b *BitReader) ReadBit() (uint8, error) {
	var current byte
	var err error
	if b.offset == 0 {
		current, err = b.r.ReadByte()
		if err != nil {
			return 0, err
		}
	} else {
		err := b.r.UnreadByte()
		if err != nil {
			return 0, err
		}
		current, err = b.r.ReadByte()
		if err != nil {
			return 0, err
		}
	}
	bit := (current >> (3 - b.offset)) & 0x1
	b.offset++
	if b.offset > 3 {
		b.offset = 0
	}
	b.position++
	return bit, nil
}

func (b *BitReader) ReadBits(num int) (uint, error) {
	if num == 0 || num > bits.UintSize {
		return 0, fmt.Errorf("bits must be > 0 and less than uint size")
	}
	var val uint
	for i := 0; i < num; i++ {
		val <<= 1
		bit, err := b.ReadBit()
		if err != nil {
			return 0, err
		}
		val |= uint(bit)
	}
	return val, nil
}

func (b BitReader) Position() uint {
	return b.position
}

func ParseInput(r io.Reader) ([]byte, error) {
	var b bytes.Buffer
	s := bufio.NewScanner(r)
	for s.Scan() {
		l := strings.TrimSpace(s.Text())
		if len(l) == 0 {
			continue
		}
		for _, r := range l {
			switch r {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				b.WriteByte(byte(r - '0'))
			case 'A', 'B', 'C', 'D', 'E', 'F':
				b.WriteByte(byte(r-'A') + 10)
			default:
				return nil, fmt.Errorf("invalid character: %v", r)
			}
		}
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func ParsePacket(b []byte) (Packet, error) {
	br := NewBitReader(bytes.NewReader(b))
	return readPacket(br)
}

func readPacket(br *BitReader) (Packet, error) {
	version, err := br.ReadBits(3)
	if err != nil {
		return nil, err
	}
	typeID, err := br.ReadBits(3)
	if err != nil {
		return nil, err
	}
	h := Header{Version: uint8(version), TypeID: Operator(typeID)}
	var packet Packet
	if Operator(typeID) == Literal {
		var literal int
		for {
			// Next 5 bits - if start with 1, read more
			moreBit, err := br.ReadBit()
			if err != nil {
				return nil, err
			}
			literal <<= 4
			nextBits, err := br.ReadBits(4)
			if err != nil {
				return nil, err
			}
			literal |= int(nextBits)
			if moreBit == 0 {
				break
			}
		}
		packet = &LiteralPacket{h: h, value: literal}
	} else {
		// Operator packet
		lengthTypeID, err := br.ReadBit()
		if err != nil {
			return nil, err
		}
		opPacket := &OperatorPacket{h: h}
		if lengthTypeID == 0 {
			totalLength, err := br.ReadBits(15)
			if err != nil {
				return nil, err
			}
			startPos := br.Position()
			for br.Position() < startPos+totalLength {
				subPacket, err := readPacket(br)
				if err != nil {
					return nil, err
				}
				opPacket.subPackets = append(opPacket.subPackets, subPacket)
			}
		} else {
			numSubPackets, err := br.ReadBits(11)
			if err != nil {
				return nil, err
			}
			for i := 0; i < int(numSubPackets); i++ {
				subPacket, err := readPacket(br)
				if err != nil {
					return nil, err
				}
				opPacket.subPackets = append(opPacket.subPackets, subPacket)
			}
		}
		packet = opPacket
	}
	return packet, nil
}
