package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func countOcurrences(line string) {
	resp, _ := http.Get(line)
	defer resp.Body.Close()
	html, _ := ioutil.ReadAll(resp.Body)
	htmlStr := string(html)

	regex := regexp.MustCompile(`[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,4}`)

	found := regex.FindAllString(htmlStr, -1)
	sort.Strings(found)
	counter := strings.Count(htmlStr, "@")
	fmt.Println("\nThere are", counter, "\"@\" on this address :", line)

	if len(found) == 0 {
		fmt.Println("I haven't found any email.")
	} else if len(found) == 1 {
		fmt.Println("I have found this email :")
	} else {
		fmt.Println("I found these emails: ")
	}

	for i := 0; i < len(found); i++ {
		fmt.Println(found[i])
	}

	defer wg.Done()
}

func readLine(path string) {
	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		wg.Add(1)
		go countOcurrences(scanner.Text())
	}
	wg.Wait()
}

func main() {
	readLine("urls.txt")
}
