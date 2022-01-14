package main

import (
	"fmt"
	"log"
	"runtime"
	"strconv"
	"strings"

	"github.com/monicagangwar/advent-of-code-2021/input"
)

type operation int

const (
	sum operation = iota
	product
	minimum
	maximum
	value
	greaterThan
	lessThan
	equal
)

func main() {
	_, currentFilePath, _, _ := runtime.Caller(0)
	inputLine := string(input.ReadInput(currentFilePath))

	var sb strings.Builder
	for _, ch := range inputLine {
		binaryValue, _ := strconv.ParseInt(string(ch), 16, 64)
		sb.WriteString(fmt.Sprintf("%04b", binaryValue))
	}

	encoded := sb.String()
	ptr := 0
	parsedPacket := parsePacket(strings.Split(encoded, ""), &ptr)
	fmt.Printf("versionSum: %d evaluatedValue: %d", parsedPacket.versionSum(), parsedPacket.evaluate())
}

func parsePacket(encoded []string, ptr *int) packet {
	p := packet{}
	version := parseVersion(encoded, ptr)
	typeId := parseTypeId(encoded, ptr)
	if typeId == value {
		p = parseLiteralPacket(encoded, ptr)
	} else {
		p = parseOperatorPacket(encoded, ptr)
	}
	p.version = version
	p.typeId = typeId

	return p
}

func parseVersion(encoded []string, ptr *int) int {
	version, err := strconv.ParseInt(strings.Join(encoded[*ptr:*ptr+3], ""), 2, 64)
	if err != nil {
		log.Fatalf("unable to parse version due to err : %s", err)
	}
	*ptr += 3
	return int(version)
}

func parseTypeId(encoded []string, ptr *int) operation {
	typeId, err := strconv.ParseInt(strings.Join(encoded[*ptr:*ptr+3], ""), 2, 64)
	if err != nil {
		log.Fatalf("unable to parse typeId due to err : %s", err)
	}
	*ptr += 3
	return operation(int(typeId))
}

func parseLiteralPacket(encoded []string, ptr *int) packet {
	var sb strings.Builder
	end := false
	for !end {
		if encoded[*ptr] == "0" {
			end = true
		}
		sb.WriteString(strings.Join(encoded[*ptr+1:*ptr+5], ""))
		*ptr += 5
	}
	literalValue, err := strconv.ParseInt(sb.String(), 2, 64)
	if err != nil {
		log.Fatalf("unable to convert value %s to int due to err: %s", sb.String(), err)
	}
	return packet{literalValue: int(literalValue)}
}

func parseOperatorPacket(encoded []string, ptr *int) packet {
	lenId, lenValue := parseLenIdAndValue(encoded, ptr)
	subPackets := make([]packet, 0)
	if lenId == 0 {
		subEnd := *ptr + lenValue
		for *ptr < subEnd {
			subPackets = append(subPackets, parsePacket(encoded, ptr))
		}
	} else {
		for idx := 1; idx <= lenValue; idx++ {
			subPackets = append(subPackets, parsePacket(encoded, ptr))
		}
	}
	return packet{subPackets: subPackets}
}

func parseLenIdAndValue(encoded []string, ptr *int) (int, int) {
	bitsNeeded := 0
	lenId := 0
	if encoded[*ptr] == "0" {
		lenId = 0
		bitsNeeded = 15
	} else {
		lenId = 1
		bitsNeeded = 11
	}
	*ptr++
	lenVal, err := strconv.ParseInt(strings.Join(encoded[*ptr:*ptr+bitsNeeded], ""), 2, 64)
	if err != nil {
		log.Fatalf("unable to parse length due to err:%s", err)
	}
	*ptr += bitsNeeded
	return lenId, int(lenVal)
}

type packet struct {
	version      int
	typeId       operation
	literalValue int
	subPackets   []packet
}

func (p packet) versionSum() int {
	if len(p.subPackets) == 0 {
		return p.version
	}
	versionSum := 0
	for _, sp := range p.subPackets {
		versionSum += sp.versionSum()
	}
	return p.version + versionSum
}

func (p packet) evaluate() int {
	switch p.typeId {
	case sum:
		v := 0
		for _, sp := range p.subPackets {
			v += sp.evaluate()
		}
		return v

	case product:
		v := 1
		for _, sp := range p.subPackets {
			v *= sp.evaluate()
		}
		return v
	case minimum:
		v := -1
		for _, sp := range p.subPackets {
			val := sp.evaluate()
			if v == -1 || val < v {
				v = val
			}
		}
		return v
	case maximum:
		v := 0
		for _, sp := range p.subPackets {
			val := sp.evaluate()
			if val > v {
				v = val
			}
		}
		return v
	case value:
		return p.literalValue
	case greaterThan:
		if p.subPackets[0].evaluate() > p.subPackets[1].evaluate() {
			return 1
		}
		return 0
	case lessThan:
		if p.subPackets[0].evaluate() < p.subPackets[1].evaluate() {
			return 1
		}
		return 0
	case equal:
		if p.subPackets[0].evaluate() == p.subPackets[1].evaluate() {
			return 1
		}
		return 0
	default:
		log.Fatalf("Packet type ID is %d is not recognized.", p.typeId)
		return 0
	}
}
