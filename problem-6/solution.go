package problem6

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseNumberArray(s string) []int {
	result := []int{}

	for _, v := range strings.Fields(s)[1:] {
		asInt, _ := strconv.Atoi(v)
		result = append(result, asInt)
	}

	return result
}

func ParseNumber(s string) int {
	n := strings.ReplaceAll(strings.Split(s, ":")[1], " ", "")
	asInt, _ := strconv.Atoi(n)

	return asInt
}

func LoadData(path string) ([]int, []int) {
	file, err := os.Open(path)

	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	time := ParseNumberArray(scanner.Text())

	scanner.Scan()
	distance := ParseNumberArray(scanner.Text())

	file.Close()

	return time, distance
}

func LoadData2(path string) (int, int) {
	file, err := os.Open(path)

	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	time := ParseNumber(scanner.Text())
	scanner.Scan()
	distance := ParseNumber(scanner.Text())

	file.Close()

	return time, distance
}

func Solve(t int, d int) []int {
	result := []int{}

	for i := 0; i < t; i++ {
		if i*(t-i) > d {
			result = append(result, i)
		}
	}

	return result
}

func Part1() {
	time, distance := LoadData("./problem-6/inputs/prod.txt")

	result := 1

	for i := range time {
		result *= len(Solve(time[i], distance[i]))
	}

	fmt.Printf("result: %v\n", result)
}

func Part2() {
	time, distance := LoadData2("./problem-6/inputs/prod.txt")
	result := len(Solve(time, distance))
	fmt.Printf("result: %v\n", result)

}

func Run() {
	Part1()
	Part2()
}
