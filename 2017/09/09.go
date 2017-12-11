package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// lines := "<{oei!a,<{i<a>"
	lines := getLinesFromFile("input.txt")
	res := getGroup(lines[0], 0, 0)
	fmt.Println(res)
}

func getGroup(sa string, level int, res int) int {
	ignoreNext := false
	for i, s := range sa {
		if ignoreNext {
			ignoreNext = false
		} else {
			switch s {
			case '!':
				ignoreNext = true
			case '{':
				return getGroup(sa[i+1:], level+1, res)
			case '<':
				return getGarbage(sa[i+1:], level, res)
			case '}':
				return getGroup(sa[i+1:], level-1, res)
			}
		}
	}
	return res
}

func getGarbage(sa string, level int, res int) int {
	ignoreNext := false
	for i, s := range sa {
		if ignoreNext {
			ignoreNext = false
		} else if s == '!' {
			ignoreNext = true
		} else if s == '>' {
			return getGroup(sa[i+1:], level, res)
			res++
		} else {
			res++
		}
	}
	return res
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
