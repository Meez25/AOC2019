package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"time"
)

//go:embed input.txt
var inputFile []byte

func main() {
	start := time.Now()
	table := bytes.Split(bytes.TrimSpace(inputFile), []byte(","))
	asInt := make([]int, len(table))
	for i := range table {
		asInt[i] = convertByteToInt(table[i])
	}

	i := 0
	asInt[1] = 12
	asInt[2] = 2
	for {
		if i >= len(asInt) {
			break
		}
		if asInt[i] == 99 {
			break
		}
		if asInt[i] == 1 {
			asInt[asInt[i+3]] = asInt[asInt[i+1]] + asInt[asInt[i+2]]
		}
		if asInt[i] == 2 {
			asInt[asInt[i+3]] = asInt[asInt[i+1]] * asInt[asInt[i+2]]
		}
		i += 4
	}
	elapsed := time.Since(start)
	fmt.Println("Part 1 : ", asInt[0], "Done in ", elapsed)
}

func convertByteToInt(input []byte) int {
	num := 0
	for _, b := range input {
		num = num*10 + int(b-'0')
	}
	return num
}
