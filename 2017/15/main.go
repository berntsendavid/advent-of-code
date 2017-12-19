package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	genA, genB := 116, 299
	factorA, factorB := 16807, 48271
	remainder := 2147483647
	count := 0
	validA, validB := false, false

	for i := 0; i < 5000000; i++ {
		validA, validB = false, false
		for !validA {
			genA = genA * factorA % remainder
			validA = genA%4 == 0
		}
		for !validB {
			genB = genB * factorB % remainder
			validB = genB%8 == 0
		}
		if sameLeastSig16(genA, genB) {
			count++
		}
	}
	fmt.Printf("Part 2: %d\n", count)
}

func sameLeastSig16(a, b int) bool {
	toString := func(sa []string) string {
		s := ""
		for _, c := range sa {
			s += c
		}
		return s
	}
	leastSig16Str := func(i int) string {
		sa := strings.Split(strconv.FormatInt(int64(i), 2), "")
		if len(sa) >= 16 {
			return toString(sa[len(sa)-16:])
		}
		return toString(sa)
	}
	return leastSig16Str(a) == leastSig16Str(b)
}
