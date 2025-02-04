package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func convertToIntSlice(slice []string) []int {
	intSlice := make([]int, len(slice))

	for i, v := range slice {
		var err error
		intSlice[i], err = strconv.Atoi(v)
		if err != nil {
			return nil
		}
	}
	return intSlice
}

func isSafe(report string) bool {
	slice := strings.Split(report, " ")
	intSlice := convertToIntSlice(slice)

	isAscending := sort.SliceIsSorted(intSlice, func(i, j int) bool {
		return intSlice[i] < intSlice[j]
	})

	isDescending := sort.SliceIsSorted(intSlice, func(i, j int) bool {
		return intSlice[i] > intSlice[j]
	})

	if !isAscending && !isDescending {
		return false
	}

	for i := range intSlice {
		if i == len(intSlice)-1 {
			continue
		}
		if math.Abs(float64(intSlice[i+1]-intSlice[i])) < 1 || math.Abs(float64(intSlice[i+1]-intSlice[i])) > 3 {
			return false
		}
	}

	return true
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
