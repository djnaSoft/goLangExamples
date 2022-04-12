package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	charWordMap := map[string][]string{}

	// assumes short lines, can adjust Scanner buffer size
	for scanner.Scan() {
		var word = scanner.Text()
		var sortedWord string = processOneLine(word)
		if _, ok := charWordMap[sortedWord]; !ok {
			charWordMap[sortedWord] = []string{}
		}
		charWordMap[sortedWord] = append(charWordMap[sortedWord], word)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//fmt.Println("result: ", charWordMap)

	for _, w := range charWordMap {
		if len(w) > 2 {
			fmt.Println(len(w), " ", w)
		}
	}
}

func processOneLine(word string) string {
	lowerWord := strings.ToLower(word)
	s := []rune(lowerWord)
	sort.Slice(s, func(i int, j int) bool { return s[i] < s[j] })

	alphaWord := string(s)
	//fmt.Println(lowerWord, "=", alphaWord)
	return alphaWord
}
