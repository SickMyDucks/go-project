package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode"
)

func main() {
	info, _ := os.Stat(os.Args[1])
	fmt.Printf("File : %s\n", info.Name())
	fmt.Printf("File size: %d bytes\n", info.Size())
	content, _ := ioutil.ReadFile(os.Args[1])
	ContentMap := maping(content)
	count(ContentMap)

}

func count(Map map[rune]int) {
	var maxCount int
	var maxLetter []rune
	minCount := 99999999
	var minLetter []rune
	for i, j := range Map {
		if unicode.IsLetter(i) {
			if j > maxCount {
				maxCount = j
				maxLetter = []rune{i}
			} else if j >= maxCount {
				maxCount = j
				maxLetter = append(maxLetter, i)
			}
			if j < minCount {
				minCount = j
				minLetter = []rune{i}
			} else if j <= minCount {
				minCount = j
				minLetter = append(minLetter, i)
			}
		}
	}
	if len(minLetter) < 2 {
		fmt.Printf("Least recurrent letter is %s with %d occurrences\n", strings.Split(string(minLetter), ""), minCount)
	} else {
		fmt.Printf("Least recurrent letters are %s with %d occurrences\n", strings.Split(string(minLetter), ""), minCount)
	}

	if len(maxLetter) < 2 {
		fmt.Printf("Most recurrent letter is %s with %d occurrences\n", strings.Join(strings.Split(string(maxLetter), ""), ", "), maxCount)
	} else {
		fmt.Printf("Most recurrent letters are %q with %d occurrences\n", strings.Join(strings.Split(string(maxLetter), ""), ", "), maxCount)
	}

}

func maping(contentByte []byte) map[rune]int {
	contentrune := []rune(string(contentByte))
	mymap := make(map[rune]int)
	for _, j := range contentrune {
		mymap[j] = mymap[j] + 1
	}
	return mymap
}
