package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func loadDataFromFile(name string) [][]int {
	file, err := os.Open(name)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var reportData [][]int

	for scanner.Scan() {
		strLine := strings.Fields(scanner.Text())
		var intValues []int
		for _, str := range strLine {
			intValue := getNumFromText(str)
			intValues = append(intValues, intValue)
		}
		reportData = append(reportData, intValues)
	}

	return reportData
}

func getNumFromText(text string) int {
	num, err := strconv.Atoi(text)
	if err != nil {
		log.Fatalln(err)
	}
	return num
}

func part1(reports [][]int) {
	//fmt.Println("Part 1 answer:", reports)

	totalSafe := 0

	for _, report := range reports {
		isIncreasing := 0
		isDecreasing := 0
		//fmt.Println(report)

		for i := 1; i < len(report); i++ {
			if report[i] > report[i-1] && (report[i]-report[i-1] >= 1 && report[i]-report[i-1] <= 3) {
				//fmt.Println("isDecreasing", report, report[i], report[i-1])
				isDecreasing++
			} else if report[i] < report[i-1] && (report[i-1]-report[i] >= 1 && report[i-1]-report[i] <= 3) {
				//fmt.Println("isIncreasing", report, report[i], report[i-1])
				isIncreasing++
			}
		}

		if isIncreasing == len(report)-1 || isDecreasing == len(report)-1 {
			totalSafe++
		}
	}
	fmt.Println("Part 1 answer:", totalSafe)
}

func main() {
	//reports := loadDataFromFile("sample.txt")
	reports := loadDataFromFile("input.txt")

	part1(reports)
	//part2(reportData)
}
