package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type pipe struct {
	id            int
	connectionIds []int
}

func main() {
	graph := make(map[int][]int)
	lines := getLinesFromFile("input.txt")
	for _, line := range lines {
		words := strings.Split(line, " <-> ")
		a, _ := strconv.Atoi(words[0])
		bs := strings.Split(words[1], ", ")
		for _, b := range bs {
			ib, _ := strconv.Atoi(b)
			graph[a] = append(graph[a], ib)
			graph[ib] = append(graph[ib], a)
		}
	}

	q := make([]int, 1)
	vis := make([]bool, len(lines))
	var a int
	groups := 0
	for i, _ := range lines {
		q = append(q, i)
		if !vis[i] {
			groups++
			for len(q) > 0 {
				a, q = q[0], q[1:]
				for _, b := range graph[a] {
					if !vis[b] {
						vis[b] = true
						q = append(q, b)
					}
				}
			}
		}
	}

	fmt.Println(groups)

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
