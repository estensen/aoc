package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("6.txt")
	defer f.Close()

	sumSingleYes := 0
	sumConsensusYes := 0

	yeses := make(map[rune]int, 26)
	groupSize := 0

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			sumSingleYes += len(yeses)

			for _, c := range yeses {
				if c == groupSize {
					sumConsensusYes++
				}
			}

			yeses = make(map[rune]int, 26)
			groupSize = 0
		} else {

			for _, c := range line {
				yeses[c]++
			}

			groupSize++
		}

	}

	// EOF
	sumSingleYes += len(yeses)

	for _, c := range yeses {
		if c == groupSize {
			sumConsensusYes++
		}
	}

	yeses = make(map[rune]int, 26)
	groupSize = 0

	fmt.Println(sumSingleYes)
	fmt.Println(sumConsensusYes)
}
