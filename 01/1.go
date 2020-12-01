package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sum := 2020
	f, _ := os.Open("1.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	list := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		i, _ := strconv.Atoi(line)
		list = append(list, i)
	}

	// First
	for _, a := range list {
		b := sum - a

		if contains(list, b) {
			fmt.Println(a)
			fmt.Println(b)
			fmt.Println(a * b)
			break
		}
	}

	// Second
out:
	for _, a := range list {
		sum2 := sum - a

		for _, b := range list {
			c := sum2 - b

			if contains(list, c) {
				fmt.Println(a)
				fmt.Println(b)
				fmt.Println(c)
				fmt.Println(a * b * c)
				break out
			}
		}
	}
}

func contains(list []int, e int) bool {
	for _, a := range list {
		if a == e {
			return true
		}
	}
	return false
}
