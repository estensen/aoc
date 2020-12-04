package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var requiresFields = [...]string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

func main() {
	f, _ := os.Open("4.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	currentPassport := ""
	allFieldsPassports := 0
	validPassports := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			currentPassport += " " + line
		} else {
			allFields, valid := validate(currentPassport)

			if allFields {
				allFieldsPassports++
			}

			if valid {
				validPassports++
			}
			currentPassport = ""
			allFields = true
			valid = true
		}
	}

	// a
	fmt.Println(allFieldsPassports)
	// b
	fmt.Println(validPassports)
}

func validate(passport string) (bool, bool) {
	allFields := true
	valid := true
	for _, field := range requiresFields {
		if !strings.Contains(passport, field) {
			allFields = false
		}
		switch field {
		case "byr":
			byr := regexp.MustCompile(`(?:^|\s)(byr):(?:(19[2-9]\d|200[0-2])(?:\s|$))?`).FindStringSubmatch(passport)
			if len(byr) == 0 || byr[2] == "" {
				valid = false
			}
		case "iyr":
			iyr := regexp.MustCompile(`(?:^|\s)(iyr):(?:(201\d|2020)(?:\s|$))?`).FindStringSubmatch(passport)
			if len(iyr) == 0 || iyr[2] == "" {
				valid = false
			}
		case "eyr":
			eyr := regexp.MustCompile(`(?:^|\s)(eyr):(?:(202\d|2030)(?:\s|$))?`).FindStringSubmatch(passport)
			if len(eyr) == 0 || eyr[2] == "" {
				valid = false
			}
		case "hgt":
			hgt := regexp.MustCompile(`(?:^|\s)(hgt):(?:((?:1[5-8]\d|19[0-3])cm|(?:59|6\d|7[0-6])in)(?:\s|$))?`).FindStringSubmatch(passport)
			if len(hgt) == 0 || hgt[2] == "" {
				valid = false
			}
		case "hcl":
			hcl := regexp.MustCompile(`(?:^|\s)(hcl):(?:(#[\da-f]{6})(?:\s|$))?`).FindStringSubmatch(passport)
			if len(hcl) == 0 || hcl[2] == "" {
				valid = false
			}
		case "ecl":
			ecl := regexp.MustCompile(`(?:^|\s)(ecl):(?:(amb|blu|brn|gry|grn|hzl|oth)(?:\s|$))?`).FindStringSubmatch(passport)
			if len(ecl) == 0 || ecl[2] == "" {
				valid = false
			}
		case "pid":
			pid := regexp.MustCompile(`(?:^|\s)(pid):(?:(\d{9})(?:\s|$))?`).FindStringSubmatch(passport)
			if len(pid) == 0 || pid[2] == "" {
				valid = false
			}
		}
	}
	return allFields, valid
}
