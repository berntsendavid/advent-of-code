package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := getLinesFromFile("input.txt")
	fw := make(map[int]int, len(lines))
	fwLength := 0
	for _, l := range lines {
		a := strings.Split(l, ": ")
		b, _ := strconv.Atoi(a[0])
		c, _ := strconv.Atoi(a[1])
		fw[b] = c
		if fwLength < b {
			fwLength = b
		}
	}

	delay := 0
	caught := true
	for caught {
		caught = false
		for i := 0; i < fwLength+1; i++ {
			if fw[i] != 0 {
				if (i+delay)%((fw[i]-1)*2) == 0 {
					caught = true
					break
				}
			}
		}
		if caught {
			delay++
		}
	}
	fmt.Println(delay)
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
