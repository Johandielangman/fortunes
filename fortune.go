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
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
)

func dirExists(path string) (bool, error) {
	if _, err := os.Stat(path); err == nil {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, nil
	} else {
		return false, err
	}
}

func listFortuneFiles(path string) []string {
	var fortuneFiles []string

	items, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range items {
		if !item.IsDir() && strings.HasSuffix(item.Name(), ".txt") {
			fortuneFiles = append(fortuneFiles, item.Name())
		}
	}
	return fortuneFiles
}

func chooseRandomElement(slice []string) string {
	return slice[rand.Intn(len(slice))]
}

func readFortunes(path string) []string {
	var (
		fortunes []string
		chunk    []string
	)

	rawFile, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	// Damn Windows
	rawFileString := strings.ReplaceAll(string(rawFile), "\r\n", "\n")

	lines := strings.Split(rawFileString, "\n")

	for _, line := range lines {
		if strings.TrimSpace(line) == "%" {
			// To account for a % at the top with no content above it
			if len(chunk) > 0 {
				// Rebuild the par and add to the final slice
				fortunes = append(fortunes, strings.Join(chunk, "\n"))

				// reset
				chunk = []string{}
			}
		} else {
			// we're still building the chunk
			chunk = append(chunk, line)
		}
	}

	// account for if there is no % at the very end of the file
	if len(chunk) > 0 {
		fortunes = append(fortunes, strings.Join(chunk, "\n"))
	}
	return fortunes
}

func main() {
	// DECLARE SOME VARIABLES
	FORTUNE_COLLECTION_DIR := "fortunes"

	// ====> CHECK TO MAKE SURE THE FORTUNE DIRECORY EXISTS
	if exists, err := dirExists(FORTUNE_COLLECTION_DIR); err != nil {
		log.Fatal(err)
	} else if !exists {
		log.Fatal("Could not find your fortune directory")
	}

	// ====> GET A SLICE OF AVAILABLE FORTUNES
	fortuneFiles := listFortuneFiles(FORTUNE_COLLECTION_DIR)
	fortuneFile := chooseRandomElement(fortuneFiles)

	// ====> READ IN ALL THE FORTUNES
	fortunes := readFortunes(filepath.Join(FORTUNE_COLLECTION_DIR, fortuneFile))

	// ====> THE FINAL OUTPUT PASSED TO COWSAY
	fortune := chooseRandomElement(fortunes)
	fmt.Println(fortune)
}
