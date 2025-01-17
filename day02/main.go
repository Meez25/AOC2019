package main

import (
	"bytes"
	_ "embed"
	"errors"
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

	initialMem := make([]int, len(asInt))
	copy(initialMem, asInt)

	runProgram(asInt)
	copy(asInt, initialMem)

	elapsed := time.Since(start)
	fmt.Println("Part 1 : ", asInt[0], "Done in ", elapsed)
	fmt.Println("Part 2 : ", solveP2(initialMem), "Done in ", time.Since(start))

}

func convertByteToInt(input []byte) int {
	num := 0
	for _, b := range input {
		num = num*10 + int(b-'0')
	}
	return num
}

func solve1(memory []int) (int, error) {
	i := 0
	memory[1] = 12
	memory[2] = 2
	for {
		if i >= len(memory) || memory[i] == 99 {
			break
		}

		if memory[i+1] >= len(memory) || memory[i+2] >= len(memory) || memory[i+3] >= len(memory) {
			return 0, errors.New("Memory error")
		}

		if memory[i] == 1 {
			memory[memory[i+3]] = memory[memory[i+1]] + memory[memory[i+2]]
		}

		if memory[i] == 2 {
			memory[memory[i+3]] = memory[memory[i+1]] * memory[memory[i+2]]
		}
		i += 4
	}
	return memory[0], nil
}

func runProgram(memory []int) error {
	i := 0
	for {
		if i >= len(memory) || memory[i] == 99 {
			break
		}

		if memory[i+1] >= len(memory) || memory[i+2] >= len(memory) || memory[i+3] >= len(memory) {
			return errors.New("Memory error")
		}

		if memory[i] == 1 {
			memory[memory[i+3]] = memory[memory[i+1]] + memory[memory[i+2]]
		}

		if memory[i] == 2 {
			memory[memory[i+3]] = memory[memory[i+1]] * memory[memory[i+2]]
		}
		i += 4
	}
	return nil
}

func solveP2(memory []int) int {
	initialMem := make([]int, len(memory))
	copy(initialMem, memory)

	for i := 0; i < 100; i++ {
		for y := 0; y < 100; y++ {
			copy(memory, initialMem)
			memory[1] = y
			memory[2] = i
			err := runProgram(memory)
			if err != nil {
				continue
			}
			if memory[0] == 19690720 {
				return y*100 + i
			}
		}
	}
	return 0
}
