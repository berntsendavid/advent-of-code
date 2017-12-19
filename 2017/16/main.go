package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	name string
	x, y int
	a, b string
}

func main() {
	programString := func(sa []string) string {
		res := ""
		for _, p := range sa {
			res += p
		}
		return res
	}

	programs := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p"}

	instr := strings.Split(getLinesFromFile("input.txt")[0], ",")
	instructions := make([]instruction, 0)
	for _, inst := range instr {
		if strings.HasPrefix(inst, "s") {
			inst = strings.TrimPrefix(inst, "s")
			steps, _ := strconv.Atoi(inst)
			instructions = append(instructions, instruction{name: "s", x: steps})
		} else if strings.HasPrefix(inst, "x") {
			inst = strings.TrimPrefix(inst, "x")
			vals := strings.Split(inst, "/")
			x, _ := strconv.Atoi(vals[0])
			y, _ := strconv.Atoi(vals[1])
			instructions = append(instructions, instruction{name: "x", x: x, y: y})
		} else if strings.HasPrefix(inst, "p") {
			inst = strings.TrimPrefix(inst, "p")
			vals := strings.Split(inst, "/")
			instructions = append(instructions, instruction{name: "p", a: vals[0], b: vals[1]})
		}
	}

	upper := 1000000000

	for i := 0; i < upper; i += 1 {

	}
	for i := 0; i < upper; i += 1 {
		for _, inst := range instructions {
			programs = getProg(inst, programs)
		}
		if programString(programs) == "abcdefghijklmnop" {
			fmt.Println(i)
			upper = upper % (i + 1)
		}
	}
	fmt.Println(programString(programs), upper)
	for i := 0; i < upper; i += 1 {
		for _, inst := range instructions {
			programs = getProg(inst, programs)
		}
	}

	fmt.Println(programString(programs))

}
func getProg(inst instruction, programs []string) []string {
	spin := func(sa []string, x int) []string { return append(sa[len(sa)-x:], sa[:len(sa)-x]...) }
	exchange := func(sa []string, x, y int) []string { sa[x], sa[y] = sa[y], sa[x]; return sa }
	partner := func(sa []string, a, b string) []string {
		aIndex := 0
		bIndex := 0
		for i, c := range sa {
			if c == a {
				aIndex = i
			} else if c == b {
				bIndex = i
			}
		}
		return exchange(sa, aIndex, bIndex)
	}
	if inst.name == "s" {
		return spin(programs, inst.x)
	} else if inst.name == "x" {
		return exchange(programs, inst.x, inst.y)
	} else {
		return partner(programs, inst.a, inst.b)
	}
}

func getLinesFromFile(fileName string) []string {
	file, err := os.Open(fileName)
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	check(err)
	return lines
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
