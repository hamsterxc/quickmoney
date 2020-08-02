package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func readValues(prompt string, processor func(string) bool) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(prompt)
		scanner.Scan()
		if processor(scanner.Text()) {
			return
		}
	}
}

func readOneInt(prompt string, check func(int) bool) (result int) {
	readValues(prompt, func(s string) bool {
		s = strings.TrimSpace(s)
		var ignored string
		n, err := fmt.Sscan(s, &result, &ignored)
		return (n == 1) && (err == io.EOF) && check(result)
	})
	return
}

func readOneIntDefault(prompt string, check func(int) bool, def int) (result int) {
	readValues(prompt, func(s string) bool {
		s = strings.TrimSpace(s)
		if s == "" {
			result = def
			return true
		} else {
			var ignored string
			n, err := fmt.Sscan(s, &result, &ignored)
			return (n == 1) && (err == io.EOF) && check(result)
		}
	})
	return
}

func readOneStringDefault(prompt string, check func(string) bool, def string) (result string) {
	readValues(prompt, func(s string) bool {
		s = strings.TrimSpace(s)
		if s == "" {
			result = def
			return true
		} else {
			result = s
			return check(result)
		}
	})
	return
}

func readTwoIntsDefault(prompt string, check func(int, int) bool, defFirst int, defSecond int) (first int, second int) {
	readValues(prompt, func(s string) bool {
		s = strings.TrimSpace(s)
		if s == "" {
			first = defFirst
			second = defSecond
			return true
		} else {
			var ignored string
			n, err := fmt.Sscan(s, &first, &second, &ignored)
			return (n == 2) && (err == io.EOF) && check(first, second)
		}
	})
	return
}
