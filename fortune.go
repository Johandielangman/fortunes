// ~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~
//      /\_/\
//     ( o.o )
//      > ^ <
//
// Author: Johan Hanekom
// Date: May 2025
//
// ~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// =============== // CONSTANTS // ===============

var fortunes []string
var par []string

// =============== // HELPER FUNCTIONS // ===============

func join(strArr []string) string {
	return strings.Join(strArr, "\n")
}

func main() {
	// ====> OPEN THE FILE AND DEFER
	file, err := os.Open("./fortunes.txt")
	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}
	defer file.Close()

	// ====> CREATE A SCANNER AND READ THE FILE LINE BY LINE
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		switch line {
		case "---":
			if len(par) > 0 {
				fortunes = append(fortunes, join(par))
				par = []string{}
			}
		default:
			if strings.TrimSpace(line) != "" {
				par = append(par, line)
			}
		}
	}

	// ====> TO ACCOUNT FOR THE FACT THAT THERE IS NO "---" AT THE END OF THE FILE
	if len(par) > 0 {
		fortunes = append(fortunes, join(par))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file", err)
	}

	// ====> GET A RANDOM ARRAY ELEMENT
	fmt.Println(fortunes[rand.Intn(len(fortunes))])
}
