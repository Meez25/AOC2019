package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"strconv"
	"time"
)

//go:embed input.txt
var inputFile []byte

type Horizontal struct {
	startX, startY, endX, endY int
}

type Vertical struct {
	startX, startY, endX, endY int
}

type Wire struct {
	vertical   []Vertical
	horizontal []Horizontal
}

func main() {
	start := time.Now()
	firstWire := Wire{vertical: make([]Vertical, 0), horizontal: make([]Horizontal, 0)}
	secondWire := Wire{vertical: make([]Vertical, 0), horizontal: make([]Horizontal, 0)}
	for i, line := range bytes.Split(bytes.TrimSpace(inputFile), []byte("\n")) {
		currentX, currentY := 0, 0
		for _, value := range bytes.Split(line, []byte(",")) {
			switch value[0] {
			case 'U':
				num, _ := strconv.Atoi(string(value[1:]))
				if i == 0 {
					firstWire.vertical = append(firstWire.vertical, Vertical{currentX, currentY, currentX, currentY - num})
				} else {
					secondWire.vertical = append(secondWire.vertical, Vertical{currentX, currentY, currentX, currentY - num})
				}
				currentY = currentY - num
			case 'D':
				num, _ := strconv.Atoi(string(value[1:]))
				if i == 0 {
					firstWire.vertical = append(firstWire.vertical, Vertical{currentX, currentY, currentX, currentY + num})
				} else {
					secondWire.vertical = append(secondWire.vertical, Vertical{currentX, currentY, currentX, currentY + num})
				}
				currentY = currentY + num
			case 'R':
				num, _ := strconv.Atoi(string(value[1:]))
				if i == 0 {
					firstWire.horizontal = append(firstWire.horizontal, Horizontal{currentX, currentY, currentX + num, currentY})
				} else {
					secondWire.horizontal = append(secondWire.horizontal, Horizontal{currentX, currentY, currentX + num, currentY})
				}
				currentX = currentX + num
			case 'L':
				num, _ := strconv.Atoi(string(value[1:]))
				if i == 0 {
					firstWire.horizontal = append(firstWire.horizontal, Horizontal{currentX, currentY, currentX - num, currentY})
				} else {
					secondWire.horizontal = append(secondWire.horizontal, Horizontal{currentX, currentY, currentX - num, currentY})
				}
				currentX = currentX - num
			}
		}
	}

	closest := -1

	// Check cross position
	for _, vertical := range firstWire.vertical {
		for _, horizontal := range secondWire.horizontal {
			manhattan, hasIntersection := checkIntersection(vertical, horizontal)
			if hasIntersection && (closest == -1 || closest > manhattan) {
				closest = manhattan
			}
		}
	}

	for _, vertical := range secondWire.vertical {
		for _, horizontal := range firstWire.horizontal {
			manhattan, hasIntersection := checkIntersection(vertical, horizontal)
			if hasIntersection && (closest == -1 || closest > manhattan) {
				closest = manhattan
			}
		}
	}

	fmt.Println("Part 1 :", closest, "in", time.Since(start))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func checkIntersection(v Vertical, h Horizontal) (int, bool) {
	if min(h.startX, h.endX) < v.startX &&
		max(h.startX, h.endX) > v.startX &&
		min(v.startY, v.endY) < h.startY &&
		max(v.startY, v.endY) > h.startY {
		return abs(v.startX) + abs(h.startY), true
	}
	return 0, false
}
