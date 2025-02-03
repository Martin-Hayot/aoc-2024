package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func isSafe(report string) bool {
	slice := strings.Split(report, " ")
	intSlice := make([]int, len(slice))
	for i, v := range slice {
		var err error
		intSlice[i], err = strconv.Atoi(v)
		if err != nil {
			return false
		}
	}
	slices.IsSortedFunc(intSlice, func(a int, b int) int {
		return
	})
	return true

	// fmt.Println(intSlice)
	// return false
}

func nbrOfSafeReports(file *os.File) int {
	scanner := bufio.NewScanner(file)
	counter := 0
	for scanner.Scan() {
		if isSafe(scanner.Text()) {
			counter++
		} else {
			continue
		}

	}

	return counter
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("cannot open file")
	}
	defer file.Close()

	fmt.Println(nbrOfSafeReports(file))

}
