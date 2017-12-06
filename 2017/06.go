package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func configurationToString(conf []int) string {
	var buffer bytes.Buffer

	for c := range conf {
		buffer.WriteString(strconv.Itoa(conf[c]))
	}
	return buffer.String()
}

func findLargestStack(stacks []int) int {
	max := 0
	maxIndex := 0
	for index, stack := range stacks {
		if stack > max {
			max = stack
			maxIndex = index
		}
	}
	return maxIndex
}

func hasBeenBefore(configuration string, configurations []string) bool {
	for _, conf := range configurations {
		if configuration == conf {
			return true
		}
	}
	return false
}
func findConfIndex(conf string, confStrings []string) int {
	for i, str := range confStrings {
		if conf == str {
			return i
		}
	}
	return 0
}
func main() {
	input := []int{11, 11, 13, 7, 0, 15, 5, 5, 4, 4, 1, 1, 7, 1, 15, 11}
	var configurationStrings []string
	configurationString := configurationToString(input)
	configurationStrings = append(configurationStrings, configurationString)

	redistributionCycles := 0
	repeatedConfiguration := false
	var boxes int

	for !repeatedConfiguration {
		index := findLargestStack(input)
		boxes = input[index]
		input[index] = 0
		for boxIndex := 0; boxIndex < boxes; boxIndex++ {
			index = (index + 1) % len(input)
			input[index]++
		}
		redistributionCycles++

		configurationString = configurationToString(input)
		if hasBeenBefore(configurationString, configurationStrings) {
			repeatedConfiguration = true
		}
		configurationStrings = append(configurationStrings, configurationString)
	}

	firstOccurance := findConfIndex(configurationString, configurationStrings)

	fmt.Print(redistributionCycles - firstOccurance - 1)
}
