package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func loadDataFromFile(name string) ([]int, []int) {
	file, err := os.Open(name)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var leftList, rightList []int

	for scanner.Scan() {
		list := strings.Fields(scanner.Text())
		num1 := getNumFromText(list[0])
		leftList = append(leftList, num1)

		num2 := getNumFromText(list[1])
		rightList = append(rightList, num2)
	}

	return leftList, rightList
}

func getNumFromText(text string) int {
	num, err := strconv.Atoi(text)
	if err != nil {
		log.Fatalln(err)
	}
	return num
}

// Function to count occurrences of elements from leftList in rightList
func countOccurrences(leftList, rightList []int) map[int]int {
	// Create a map to store the counts of elements in rightList
	countMap := make(map[int]int)
	for _, value := range rightList {
		countMap[value]++ // Count occurrences of each element in rightList
	}

	// Create a result map for elements in leftList
	result := make(map[int]int)
	for _, value := range leftList {
		if count, exists := countMap[value]; exists {
			result[value] += count // If already present add to the count for the value
		} else {
			result[value] = 0
		}
	}

	return result
}

func part1(leftList, rightList []int) {
	sort.Sort(sort.IntSlice(leftList))
	sort.Sort(sort.IntSlice(rightList))

	var sum int
	for i := 0; i < len(leftList); i++ {
		distance := leftList[i] - rightList[i]
		if distance < 0 {
			distance = -distance
		}
		sum += distance
	}
	fmt.Println("total distance - ", sum)
}

func part2(leftList, rightList []int) {
	countMap := countOccurrences(leftList, rightList)

	var similarityTotal int
	for key, count := range countMap {
		similarityTotal += key * count
	}
	fmt.Println("similarity score - ", similarityTotal)

}

func main() {
	leftList, rightList := loadDataFromFile("input.txt")
	//leftList, rightList := loadDataFromFile("sample.txt")

	part1(leftList, rightList)
	part2(leftList, rightList)
}
