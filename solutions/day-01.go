package solutions

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func leftClick(position, clicks int) int {
	if position-(clicks) < 0 {
		return 100 + position - clicks
	} else {
		return position - clicks
	}
}

func rightClick(position, clicks int) int {
	if (clicks%100)+position >= 100 {
		return clicks + position - 100
	} else {
		return position + clicks
	}
}

func Day01() {
	content, _ := os.ReadFile("inputs/day-01.txt")
	input := strings.Split(strings.TrimSpace(string(content)), "\n")

	i := 0
	position := 50
	zeroCount := 0
	for i < len(input) {
		click := input[i]
		direction := string(click[0])
		clicks, _ := strconv.Atoi(click[1:])
		// discard the full rotations as they don't affect the position
		remainingClicks := clicks % 100

		switch direction {
		case "L":
			position = leftClick(position, remainingClicks)
		case "R":
			position = rightClick(position, remainingClicks)
		}

		if position == 0 {
			zeroCount += 1
		}

		i += 1
	}

	fmt.Println("Password:", zeroCount)
}
