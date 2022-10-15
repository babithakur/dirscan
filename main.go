package main

import (
	"bufio"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"syscall"
)

func main() {
	fmt.Println("****Gobuster Lite****")
	fmt.Println("******By: A$TRA******")
	var web_url string
	var file_path string
	fmt.Println("Enter the website url: ")
	fmt.Scanln(&web_url)
	fmt.Println("Enter the directory path: ")
	fmt.Scanln(&file_path)
	lines := LinesInFile(file_path)
	valid_url, _ := url.Parse(web_url)
	if valid_url.Scheme == "" {
		fmt.Println("Invalid url!")
		os.Exit(1)
	}

	fmt.Println("Gobuster lite is running....")
	for dir := 0; dir < len(lines); dir++ {
		d := lines[dir]
		response, err := http.Get(web_url + "/" + d)
		if errors.Is(err, syscall.ECONNREFUSED) {
			fmt.Println("Unable to connect to host!")
			os.Exit(2)
		}
		if err != nil {
			fmt.Println("Something went wrong :(")
			os.Exit(2)
		}
		defer response.Body.Close()
		if response.StatusCode == 200 {
			fmt.Println("[+] Found: " + "/" + d)
		}
	}
	fmt.Println("[+] Scan finished!")
}

func LinesInFile(fileName string) []string {
	f, _ := os.Open(fileName)
	scanner := bufio.NewScanner(f)
	result := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}
	return result
}
