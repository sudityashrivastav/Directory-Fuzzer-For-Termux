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

	if len(os.Args) < 5 {
		usage()
	}

	var (
		URL                 = os.Args[1]
		fileLocation        = os.Args[2]
		threadsSTR          = os.Args[3]
		status_codes_string = os.Args[4]
	)

	status_codes := strings.Split(status_codes_string, ",")

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
		go start(URL, word, semaphore, status_codes)
	}

	wg.Wait()

	fmt.Println("Fuzzing Completed..")

}

func usage() {
	fmt.Println("\nExample:")
	fmt.Println("./fuzzer <url> <wordlist> <threads> <status codes>")
	fmt.Println("./fuzzer https://google.com wordlist.txt 40 200,202,206\n")
	os.Exit(0)
}

func start(url string, word string, semaphore chan struct{}, status_codes []string) {

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

	for _, status_string := range status_codes {
		status, _ := strconv.Atoi(status_string)
		if res.StatusCode == status {
			fmt.Println(res.StatusCode, "\t", address)
		}
	}

}
