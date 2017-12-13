package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	charcode := func(s string) string { return strconv.Itoa(int(s[0])) }
	xor := func(a, b int) int { return a ^ b }
	zeroPad := func(s string) string {
		if len(s) == 1 {
			return "0" + s
		}
		return s
	}
	toHex := func(s string) string {
		i, _ := strconv.Atoi(s)
		return strconv.FormatInt(int64(i), 16)
	}
	numbers := createRangeArr(256)
	input := "189,1,111,246,254,2,0,120,215,93,255,50,84,15,94,62"
	lengths := Map(charcode, strings.Split(input, ""))
	lengths = append(lengths, []string{"17", "31", "73", "47", "23"}...)
	pos, skip := 0, 0

	for i := 0; i < 64; i++ {
		for _, length := range lengths {
			l, _ := strconv.Atoi(length)
			numbers = append(numbers[pos:], numbers[:pos]...)
			numbers = append(reverse(numbers[:l]), numbers[l:]...)
			numbers = append(numbers[(len(numbers)-pos):], numbers[:(len(numbers)-pos)]...)
			pos = (pos + l + skip) % len(numbers)
			skip++
		}
	}

	denseHash := make([]string, 0)
	for i := 0; i < 16; i++ {
		o := Reduce(xor, numbers[i*16:i*16+16], 0)
		denseHash = append(denseHash, strconv.Itoa(o))
	}

	hex := Map(toHex, denseHash)
	result := Map(zeroPad, hex)
	resStr := ""
	for _, r := range result {
		resStr += r
	}

	fmt.Println(resStr)
}

func Reduce(f func(int, int) int, ia []int, acc int) int {
	for _, i := range ia {
		acc = f(i, acc)
	}
	return acc
}

func Map(f func(string) string, sa []string) []string {
	maps := make([]string, len(sa))
	for i, s := range sa {
		maps[i] = f(s)
	}
	return maps
}
func reverse(list []int) []int {
	for i := 0; i < len(list)/2; i++ {
		j := len(list) - i - 1
		list[i], list[j] = list[j], list[i]
	}
	return list

}

func createRangeArr(size int) []int {
	lst := make([]int, size)
	for i := 0; i < size; i++ {
		lst[i] = i
	}
	return lst
}
