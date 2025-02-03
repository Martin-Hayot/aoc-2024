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

func calculateDistance(file *os.File) int {
	var left []int
	var right []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, "   ")
		leftNumber, _ := strconv.Atoi(numbers[0])
		rightNumber, _ := strconv.Atoi(numbers[1])
		left = append(left, leftNumber)
		right = append(right, rightNumber)
	}

	sort.Ints(left)
	sort.Ints(right)

	var totalDistance int

	for index := range left {
		result := int(math.Abs(float64(left[index] - right[index])))
		totalDistance += result
	}
	return totalDistance
}

// Calculate the similarity of 2 lists
func calculateSimilarity(file *os.File) int {
	var left []int
	var right []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, "   ")
		leftNumber, _ := strconv.Atoi(numbers[0])
		rightNumber, _ := strconv.Atoi(numbers[1])
		left = append(left, leftNumber)
		right = append(right, rightNumber)
	}

	var similarityScore int

	for i, value := range left {
		counter := 0
		for j := range right {
			if left[i] == right[j] {
				counter++
			}
		}
		similarityScore = similarityScore + (counter * value)
	}
	return similarityScore
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("error reading input file")
	}
	defer file.Close()

	// totalDistance := calculateDistance(file)
	similarityScore := calculateSimilarity(file)
	// fmt.Println(totalDistance)
	fmt.Println(similarityScore)

	fmt.Println(fmt.Sprintf("Welcome %s", "name"))
}
