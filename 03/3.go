package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// a
	fmt.Println(getTrees(3, 1))

	// b
	a := getTrees(1, 1)
	fmt.Println(a)
	b := getTrees(3, 1)
	fmt.Println(b)
	c := getTrees(5, 1)
	fmt.Println(c)
	d := getTrees(7, 1)
	fmt.Println(d)
	e := getTrees(1, 2)
	fmt.Println(e)
	fmt.Println(a * b * c * d * e)
}

func getTrees(right int, down int) int {
	f, _ := os.Open("3.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	x := 0
	y := 0
	lineNum := 0
	trees := 0

	for scanner.Scan() {
		line := scanner.Text()

		if lineNum%down == 0 {
			if string(line[x%(len(line))]) == "#" {
				trees++
			}
			y += down
			x += right
		}
		lineNum++
	}
	return trees
}
