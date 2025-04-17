package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/inancgumus/screen"
)

const (
	defaultLength int = 8
	complexity    int = defaultLength * 2
)

func main() {
	screen.Clear()
	screen.MoveTopLeft()

	fmt.Println("Password Generator With MATRIX way")

	length := defaultLength
	complexity := complexity

	args := os.Args[1:]
	if len(args) > 0 {
		if argLength, err := strconv.Atoi(args[0]); err == nil {
			length = argLength
			complexity = argLength * 2
		}
		if len(args) > 1 {
			if complex, err := strconv.Atoi(args[1]); err == nil {
				complexity = complex
			}
		}
	}

	fmt.Printf("Generating password length %v and complexity of %v \n", length, complexity)
	fmt.Println(strings.Repeat("=", 50))

	password, matrix := printPasswordPart(length, complexity)

	fmt.Println(strings.Repeat("-", 50))

	err := clipboard.WriteAll(password)
	if err != nil {
		fmt.Println("Error copying to clipboard:", err)
	}

	fmt.Println("Copied to clipboard ✅")

	fmt.Println("Tries:")

	for _, row := range matrix {
		for _, char := range row {
			fmt.Print(char)
		}
		fmt.Println()
	}

}
