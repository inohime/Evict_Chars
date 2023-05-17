package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	ch1 := make(chan []string)
	defer close(ch1)

	_, file, _, _ := runtime.Caller(0)
	dataPath := filepath.Join(filepath.Dir(file), "../")

	go ReadFile(dataPath+"/test_data_2.txt", ch1)
	go ReadFile(dataPath+"/test_data_1.txt", ch1)

	x := EvictChars(strings.Join(<-ch1, ", "))
	y := EvictChars(strings.Join(<-ch1, ", "))

	fmt.Printf("first:\n%s\n\n", x)
	fmt.Printf("second:\n%s\n\n", y)
}

func ReadFile(filePath string, ch chan []string) {
	f1, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Failed to read file:", err.Error())
		return
	}
	defer f1.Close()

	var lines []string
	scanner := bufio.NewScanner(f1)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	ch <- lines
}

func EvictChars(str string) string {
	if len(str) <= 1024 {
		fmt.Println("Length of string is less than 84 words (1024 chars)")
		return str
	}

	fmt.Println("[ Before ] length of string:", len(str))

	for i := 0; i < len(str); i++ {
		if i >= 1024 { // 84 words
			str = strings.Join(
				strings.Split(
					str[:strings.LastIndex(str[:i], ",")],
					", ",
				),
				", ",
			)
			break
		}
	}

	fmt.Println("[ After ] length of string:", len(str))

	return str
}
