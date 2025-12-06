package main

import (
	"advent_of_code_2025/solutions"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <day>")
		fmt.Println("Example: go run . 1")
		os.Exit(1)
	}

	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid argument:", os.Args[1], "\nPlease try again")
		os.Exit(1)
	} else if day < 1 || day > 12 {
		fmt.Println("Day number must be between 1 and 12 \nPlease try again")
		os.Exit(1)
	}

	switch day {
	case 1:
		solutions.Day01()
	default:
		fmt.Printf("Day %d not implemented yet\n", day)
		os.Exit(1)
	}
}
