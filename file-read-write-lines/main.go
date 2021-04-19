package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	data := []string{"one", "two", "three"}
	if err := writeToFile("test.txt", data); err != nil {
		panic(err)
	}
	newData, err := readLines("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(newData)
}

func writeToFile(filename string, data []string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed creating file: %s", err)
	}

	dataWriter := bufio.NewWriter(file)

	for i := range data {
		_, _ = dataWriter.WriteString(data[i] + "\n")
	}

	dataWriter.Flush()
	file.Close()
	return nil
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
