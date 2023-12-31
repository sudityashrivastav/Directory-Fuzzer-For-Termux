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
	wg                  sync.WaitGroup
	status_codes_string = "404"
)

func main() {
	banner()
	if len(os.Args) < 4 {
		usage("Provide all values")
	}

	var (
		URL          = os.Args[1]
		fileLocation = os.Args[2]
		threadsSTR   = os.Args[3]
	)

	if len(os.Args) == 5 {
		status_codes_string = os.Args[4]
	}

	status_codes := strings.Split(status_codes_string, ",")
	threads, err := strconv.Atoi(threadsSTR)

	if err != nil {
		usage("Invalid arguments")
	}

	semaphore := make(chan struct{}, threads)

	file, err := os.Open(fileLocation)
	if err != nil {
		usage("Invalid file location")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fmt.Println("\033[36m", "Status Code\tURL", "\033[0m")
	for scanner.Scan() {
		word := scanner.Text()
		wg.Add(1)
		go start(URL, word, semaphore, status_codes)
	}

	wg.Wait()

	fmt.Println("Fuzzing Completed..")
}

func start(url string, word string, semaphore chan struct{}, status_codes []string) {
	defer wg.Done()

	semaphore <- struct{}{}
	defer func() { <-semaphore }()

	if !strings.HasSuffix(url, "/") {
		url = url + "/"
	}

	word = strings.TrimPrefix(word, "/")
	address := url + word

	address = strings.TrimSpace(address)

	req, err := http.NewRequest("GET", address, nil)

	if err != nil {
		//	fmt.Println("Error at making request: ", err.Error())
		return
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36")

	client := &http.Client{Timeout: 5 * time.Second}
	res, err := client.Do(req)

	if err != nil {
		// fmt.Println("Error at making request: ", err.Error())
		return
	}
	defer res.Body.Close()

	result := true
	for _, status_string := range status_codes {
		status, _ := strconv.Atoi(status_string)
		if res.StatusCode == status {
			result = false
			break
		}
	}

	if result {
		if res.StatusCode == 200 {
			fmt.Println("\033[32m", res.StatusCode, "\t\t", address, "\033[0m")
		} else {
			fmt.Println("\033[33m", res.StatusCode, "\t\t", address, "\033[0m")
		}
	}
}

func usage(msg string) {
	fmt.Println("\nError:\033[31m ", msg, "\033[0m")
	fmt.Println("\033[32mfuzz <url> <wordlist> <threads> <negative status codes>")
	fmt.Println("fuzz https://google.com wordlist.txt 40 404,500,302\033[0m")
	os.Exit(0)
}

func banner() {
	fmt.Println("\033[31m", `
███████╗██╗   ██╗███████╗███████╗███████╗██████╗
██╔════╝██║   ██║╚══███╔╝╚══███╔╝██╔════╝██╔══██╗
█████╗  ██║   ██║  ███╔╝   ███╔╝ █████╗  ██████╔╝
██╔══╝  ██║   ██║ ███╔╝   ███╔╝  ██╔══╝  ██╔══██╗
██║     ╚██████╔╝███████╗███████╗███████╗██║  ██║
╚═╝      ╚═════╝ ╚══════╝╚══════╝╚══════╝╚═╝  ╚═╝`)
	fmt.Println("\tby Anon Shrivastav\n", "\033[0m")

}
