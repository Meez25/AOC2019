package main

import (
	_ "embed"
	"fmt"
	"time"
)

//go:embed input.txt
var data []byte

func computeFuel(fuel, total int) int {
	if fuel <= 0 {
		return total
	}
	fuelNeeded := fuel/3 - 2
	if fuelNeeded <= 0 {
		return total
	}
	return computeFuel(fuelNeeded, total+fuelNeeded)
}

func main() {
	start := time.Now()
	var total, totalp2 int
	i := 0

	for i < len(data) {
		if data[i] == '\n' {
			i++
			continue
		}
		num := 0
		for i < len(data) && data[i] != '\n' {
			num = num*10 + int(data[i]-'0')
			i++
		}
		total += num/3 - 2
		totalp2 += computeFuel(num, 0)
	}

	elapsed := time.Since(start)
	fmt.Printf("Part one: %d, Part two: %d in %.3fµs\n",
		total, totalp2, float64(elapsed.Nanoseconds())/1000.0)
}
