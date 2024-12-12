package day1

import (
	"advent-of-code-2k24/utils"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func CalculateTotalDistance() (int, error) {
	leftList, rightList, err := readInput("day-1/input.txt")
	if err != nil {
		return 0, err
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	distances, err := getDistances(leftList, rightList)
	if err != nil {
		return 0, err
	}

	sum := sum(distances)

	return sum, nil
}

func CalculateSimilarityScore() (int, error) {
	leftList, rightList, err := readInput("day-1/input.txt")
	if err != nil {
		return 0, err
	}
	scores := []int{}

	for _, leftNumber := range leftList {
		occurencesCount := countOccurences(rightList, leftNumber)
		scores = append(scores, leftNumber*occurencesCount)
	}

	sum := sum(scores)

	return sum, nil
}

func readInput(filepath string) ([]int, []int, error) {
	leftList := []int{}
	rightList := []int{}

	content, err := os.ReadFile(filepath)
	if err != nil {
		return leftList, rightList, err
	}

	splittedContent := strings.Split(string(content), "\r\n") // WINDOWS

	for _, entry := range splittedContent {
		var (
			leftNumber  int
			rightNumber int
		)
		if entry == "" {
			continue
		}
		leftNumber, rightNumber, err = extractNumbersFromLine(entry)

		if err != nil {
			break
		}

		leftList = append(leftList, leftNumber)
		rightList = append(rightList, rightNumber)
	}

	return leftList, rightList, err
}

func extractNumbersFromLine(line string) (int, int, error) {
	splittedLine := strings.Split(line, "   ")
	if len(splittedLine) != 2 {
		return 0, 0, fmt.Errorf("did not found 2 numbers in one line")
	}
	elem1, err := strconv.Atoi(splittedLine[0])
	if err != nil {
		return 0, 0, err
	}
	elem2, err := strconv.Atoi(splittedLine[1])
	if err != nil {
		return 0, 0, err
	}
	return elem1, elem2, nil
}

func getDistances(leftList []int, rightList []int) ([]int, error) {
	if len(leftList) != len(rightList) {
		return []int{}, fmt.Errorf("left and right list are not the same size")
	}
	distances := []int{}
	for i := 0; i < len(leftList); i++ {
		distances = append(distances, utils.AbsInt(leftList[i]-rightList[i]))
	}

	return distances, nil
}

func sum(array []int) int {
	sum := 0
	for _, item := range array {
		sum += item
	}
	return sum
}

func countOccurences(slice []int, item int) int {
	count := 0
	for _, s := range slice {
		if s == item {
			count++
		}
	}
	return count
}
