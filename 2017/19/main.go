package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type pos struct {
	x, y int
	val  string
}

func (p pos) get(positions []pos) (pos, []pos) {
	for i, a := range positions {
		if p.x == a.x && p.y == a.y {
			return a, append(positions[:i], positions[i+1:]...)
		}
	}
	return p, positions
}

func (p pos) in(positions []pos) bool {
	for _, a := range positions {
		if p.x == a.x && p.y == a.y {
			return true
		}
	}
	return false
}
func (p pos) next(a pos) pos {
	p.x += a.x
	p.y += a.y
	return p
}

func main() {
	lines := getLinesFromFile("input.txt")
	regex, _ := regexp.Compile(`[A-Z]`)
	north, south, east, west := pos{0, -1, "|"}, pos{0, 1, "|"}, pos{1, 0, "-"}, pos{-1, 0, "-"}
	unseen := make([]pos, 0)
	res := make([]string, 0)
	for j, line := range lines {
		cells := strings.Split(line, "")
		for i, cell := range cells {
			if cell != " " {
				unseen = append(unseen, pos{i, j, cell})
			}
		}
	}
	var curr pos
	for i, p := range unseen {
		if p.y == 0 {
			curr = p
			unseen = append(unseen[:i], unseen[i+1:]...)
			break
		}
	}
	queued := make([]pos, 0)

	direction := south
	count := 0
	for len(unseen) > 0 {
		if curr.in(unseen) {
			curr, unseen = curr.get(unseen)
		}
		if regex.MatchString(curr.val) {
			res = append(res, curr.val)
			curr.val = direction.val
		} else if curr.val == "+" {
			queued = append(queued, curr.next(north), curr.next(south), curr.next(east), curr.next(west))
			for _, q := range queued {
				if q.in(unseen) {
					direction = getDir(curr, q)
					curr = q
				}
			}
			queued = make([]pos, 0)
			count++
		} else {
			curr = curr.next(direction)
			count++
		}
	}
	fmt.Println("Part 1: ", strings.Join(res, ""))
	fmt.Println("Part 2: ", count+1)

}
func getDir(a, b pos) pos {
	north, south, east, west := pos{0, -1, "|"}, pos{0, 1, "|"}, pos{1, 0, "-"}, pos{-1, 0, "-"}
	dx := b.x - a.x
	dy := b.y - a.y
	if dx == 1 {
		return east
	}
	if dx == -1 {
		return west
	}
	if dy == 1 {
		return south
	}
	return north
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
