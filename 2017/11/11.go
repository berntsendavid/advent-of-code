package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type pos struct {
	x, y, z int
}

func (p *pos) N() {
	p.z -= 1
	p.y += 1
}
func (p *pos) S() {
	p.z += 1
	p.y -= 1
}
func (p *pos) NW() {
	p.x += 1
	p.z -= 1
}
func (p *pos) SE() {
	p.x -= 1
	p.z += 1
}
func (p *pos) NE() {
	p.x -= 1
	p.y += 1
}
func (p *pos) SW() {
	p.y -= 1
	p.x += 1
}

func main() {
	lines := getLinesFromFile("input.txt")
	instructions := strings.Split(lines[0], ",")
	position, maxDist := goTo(instructions)
	dist := distanceFromOrigin(position)
	fmt.Println(dist, maxDist)
}
func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
func distanceFromOrigin(p pos) int {
	return (abs(p.x) + abs(p.y) + abs(p.z)) / 2
}

func goTo(sa []string) (pos, int) {
	pos := pos{0, 0, 0}
	maxDist := 0
	fmt.Println(len(sa))
	for _, s := range sa {
		switch s {
		case "n":
			pos.N()
		case "s":
			pos.S()
		case "se":
			pos.SE()
		case "sw":
			pos.SW()
		case "ne":
			pos.NE()
		case "nw":
			pos.NW()
		default:
		}
		dist := distanceFromOrigin(pos)
		if dist > maxDist {
			maxDist = dist
		}
	}
	return pos, maxDist
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
