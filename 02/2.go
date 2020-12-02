package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	a()
	b()
}

func a() {
	f, _ := os.Open("2.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	counter := 0

	for scanner.Scan() {
		line := scanner.Text()

		s := strings.Split(line, " ")
		nums := strings.Split(s[0], "-")
		min, _ := strconv.Atoi(nums[0])
		max, _ := strconv.Atoi(nums[1])
		char := s[1][0]
		passwd := s[2]

		matches := 0
		for i := 0; i < len(passwd); i++ {
			if passwd[i] == char {
				matches++
			}
		}

		if min <= matches && matches <= max {
			counter++
		}
	}

	fmt.Println(counter)
}

func b() {
	f, _ := os.Open("2.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	counter := 0

	for scanner.Scan() {
		line := scanner.Text()

		s := strings.Split(line, " ")
		nums := strings.Split(s[0], "-")
		first, _ := strconv.Atoi(nums[0])
		second, _ := strconv.Atoi(nums[1])
		char := s[1][0]
		passwd := s[2]

		if (passwd[first-1] == char) != (passwd[second-1] == char) {
			counter++
		}
	}

	fmt.Println(counter)
}
