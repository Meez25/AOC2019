package main

import (
	_ "embed"
	"fmt"
	"time"
)

//go:embed input.txt
var data []byte

func main() {
	start := time.Now()

	var total int
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
	}

	elapsed := time.Since(start)
	fmt.Printf("Part one: %d, in %.3fµs\n", total, float64(elapsed.Nanoseconds())/1000.0)
}
