package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func convertToIntSlice(report string) []int {
	slice := strings.Split(report, " ")
	intSlice := make([]int, len(slice))

	for i, v := range slice {
		num, err := strconv.Atoi(v)
		if err != nil {
			return nil
		}
		intSlice[i] = num
	}
	return intSlice
}

func isSafe(intSlice []int) bool {
	if len(intSlice) < 2 {
		return false
	}

	isAscending := sort.SliceIsSorted(intSlice, func(i, j int) bool {
		return intSlice[i] < intSlice[j]
	})
	isDescending := sort.SliceIsSorted(intSlice, func(i, j int) bool {
		return intSlice[i] > intSlice[j]
	})

	if !isAscending && !isDescending {
		return false
	}

	for i := 0; i < len(intSlice)-1; i++ {
		diff := intSlice[i+1] - intSlice[i]
		if diff < -3 || diff > 3 || diff == 0 {
			return false
		}
	}

	return true
}

func isSafeAfterOneRemoval(report string) bool {
	intSlice := convertToIntSlice(report)

	if isSafe(intSlice) {
		return true
	}

	for i := range intSlice {
		modifiedSlice := append(intSlice[:i:i], intSlice[i+1:]...)
		if isSafe(modifiedSlice) {
			return true
		}
	}
	return false
}

func nbrOfSafeReports(file *os.File) int {
	scanner := bufio.NewScanner(file)
	counter := 0
	for scanner.Scan() {
		if isSafeAfterOneRemoval(scanner.Text()) {
			counter++
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
