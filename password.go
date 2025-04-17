package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func matrixCharArray(m Matrix) [][]string {
	// Initialize a 2D slice
	matrix := make([][]string, m.rows)
	for i := range matrix {
		matrix[i] = make([]string, m.cols)
		for j := range matrix[i] {
			r := rand.Intn(m.max-m.min+1) + m.min
			matrix[i][j] = string(rune(r))
		}
	}

	return matrix
}

type Matrix struct {
	rows int
	cols int
	min  int
	max  int
}

func printPasswordPart(length int, complexity int) (string, [][]string) {
	// generate random numbers
	matrix := matrixCharArray(Matrix{
		rows: length,
		cols: complexity,
		min:  35,
		max:  126,
	})

	delay := time.Duration(1000/complexity) * time.Millisecond

	if delay <= 10 {
		delay = 10 * time.Millisecond
	}

	pass := make([]string, 0, length)

	for _, row := range matrix {
		for ic, char := range row {
			fmt.Print(char)
			time.Sleep(delay)
			fmt.Print("\b")
			if ic == len(row)-1 {
				fmt.Print(char)
				pass = append(pass, char)
			}
		}
	}
	fmt.Println()

	return strings.Join(pass, ""), matrix
}
