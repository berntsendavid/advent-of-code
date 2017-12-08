package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	lines := getLinesFromFile("input.txt")
	registers := make(map[string]int)
	var highest, val int
	for _, line := range lines {
		registers, val = updateRegisters(registers, line)
		if val > highest {
			highest = val
		}
	}
	var max int
	for _, v := range registers {
		if v > max {
			max = v
		}
	}
	fmt.Println(max, highest)
}

func updateRegisters(reg map[string]int, str string) (map[string]int, int) {
	res := getIntstructions(str)
	return calculateNewValue(res, reg)
}

func calculateNewValue(instr []string, reg map[string]int) (map[string]int, int) {
	targetRegister := instr[1]
	operation := instr[2]
	value, err := strconv.Atoi(instr[3])
	check(err)
	conditionalRegister := instr[4]
	conditional := instr[5]
	conditionalValue, err := strconv.Atoi(instr[6])
	check(err)
	reg[targetRegister] = initMapKeyIfNeeded(reg, targetRegister)
	reg[conditionalRegister] = initMapKeyIfNeeded(reg, conditionalRegister)
	if valid(conditional, reg[conditionalRegister], conditionalValue) {
		reg[targetRegister] = incOrDec(operation, reg[targetRegister], value)
	}
	return reg, reg[targetRegister]
}

func incOrDec(op string, value int, newVal int) int {
	if op == "inc" {
		return value + newVal
	}
	return value - newVal
}

func valid(cond string, lhv, rhv int) bool {
	switch cond {
	case "<":
		return lhv < rhv
	case ">":
		return lhv > rhv
	case "<=":
		return lhv <= rhv
	case ">=":
		return lhv >= rhv
	case "==":
		return lhv == rhv
	case "!=":
		return lhv != rhv
	default:
		return false
	}
}
func initMapKeyIfNeeded(m map[string]int, k string) int {
	if val, ok := m[k]; ok {
		return val
	}
	return 0
}

func getIntstructions(str string) []string {
	regex, err := regexp.Compile(`([a-z]+) (inc|dec) (-?\d*) if ([a-z]+) ([<>!=]?=?) (-?\d*)`)
	check(err)
	return regex.FindStringSubmatch(str)
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
