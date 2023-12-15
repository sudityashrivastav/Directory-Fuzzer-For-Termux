package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func main() {

	if len(os.Args) < 4 {
		usage()
	}

	var (
		URL          = os.Args[1]
		fileLocation = os.Args[2]
		threadsSTR   = os.Args[3]
	)

	threads, err := strconv.Atoi(threadsSTR)

	if err != nil {
		fmt.Println("Invalid arguments")
		usage()
	}

	semaphore := make(chan struct{}, threads)

	file, err := os.Open(fileLocation)
	if err != nil {
		fmt.Println("Invalid file location")
		usage()
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		word := scanner.Text()
		wg.Add(1)
		go start(URL, word, semaphore)
	}

	wg.Wait()

	fmt.Println("Fuzzing Completed..")

}

func usage() {
	fmt.Println("USAGE: go run main.go <url> <wordlist> <threads>")
	os.Exit(0)
}

func start(url string, word string, semaphore chan struct{}) {

	defer wg.Done()

	semaphore <- struct{}{}
	defer func() { <-semaphore }()

	if !strings.HasSuffix(url, "/") {
		url = strings.TrimSpace(url) + "/"
	}

	word = strings.TrimPrefix(word, "/")

	address := url + word

	client := &http.Client{Timeout: 5 * time.Second}
	res, err := client.Get(address)

	if err != nil {
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 404 {
		fmt.Println(address, "\t\t", res.StatusCode)
	}

}
