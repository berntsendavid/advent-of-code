package main

import (
	"fmt"
	"strconv"
	"strings"
)

type pos struct {
	x, y int
}

func (p pos) equal(a pos) bool { return p.x == a.x && p.y == a.y }
func (p pos) in(pa []pos) bool {
	for _, a := range pa {
		if p.equal(a) {
			return true
		}
	}
	return false
}

func (p pos) removeFrom(pa []pos) []pos {
	for i, a := range pa {
		if a.equal(p) {
			return append(pa[:i], pa[i+1:]...)
		}
	}
	return pa
}

func main() {
	input := "uugsqrei"
	unseen := make([]pos, 0)
	rows := make([]string, 128)
	for i, _ := range rows {
		hash := Knothash(input + "-" + strconv.Itoa(i))
		binHash := bin(strings.Split(hash, ""))
		for j, d := range binHash {
			if d == "1" {
				unseen = append(unseen, pos{i, j})
			}
		}
	}
	fmt.Printf("Part 1: %d\n", len(unseen))
	queued := make([]pos, 0)
	var curr pos
	count := 0
	for len(unseen) > 0 {
		queued = append(queued, unseen[0])
		for len(queued) > 0 {
			curr, queued = queued[0], queued[1:]
			if curr.in(unseen) {
				unseen = curr.removeFrom(unseen)
				queued = append(queued, pos{curr.x - 1, curr.y},
					pos{curr.x + 1, curr.y},
					pos{curr.x, curr.y - 1},
					pos{curr.x, curr.y + 1})
			}
		}
		count += 1
	}
	fmt.Printf("Part 2: %d\n", count)
}

func bin(charArray []string) []string {
	res := ""
	for _, char := range charArray {
		val := charVal(char)
		res += toBinaryString(val)
	}
	return strings.Split(res, "")
}

func toBinaryString(i int) string {
	binaryInt := i / 8 * 1000
	i = i % 8
	binaryInt += i / 4 * 100
	i = i % 4
	binaryInt += i / 2 * 10
	i = i % 2
	binaryInt += i
	return fmt.Sprintf("%04d", binaryInt)
}
func charVal(char string) int {
	var val int
	if char == "a" {
		val = 10
	} else if char == "b" {
		val = 11
	} else if char == "c" {
		val = 12
	} else if char == "d" {
		val = 13
	} else if char == "e" {
		val = 14
	} else if char == "f" {
		val = 15
	} else {
		val, _ = strconv.Atoi(char)
	}
	return val
}
