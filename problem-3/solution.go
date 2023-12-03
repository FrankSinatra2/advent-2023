package problem3

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"unicode"
)

type NumberInfo struct {
	Left  int
	Right int

	Value int
}

func IsSymbol(r rune) bool {
	return !unicode.IsDigit(r) && r != '.'
}

func Surround(n NumberInfo, rowLength int) []int {
	l := n.Right - n.Left
	result := make([]int, (l)*2+6)
	resultIndex := 0

	// top & bottom
	for i := 0; i < l+2; i++ {
		result[resultIndex] = (n.Left - 1 + i) - rowLength
		result[resultIndex+1] = (n.Left - 1 + i) + rowLength
		resultIndex += 2
	}

	result[len(result)-2] = n.Left - 1
	result[len(result)-1] = n.Right

	return result
}

func Part1() {
	file, err := os.Open("./problem-3/inputs/prod.txt")

	if err != nil {
		fmt.Printf("err: %v\n", err.Error())
		return
	}

	scanner := bufio.NewScanner(file)

	rowLength := 0

	text := ""

	for scanner.Scan() {
		t := scanner.Text()
		rowLength = len(t)
		text += t
	}

	valid := []NumberInfo{}

	for i := 0; i < len(text); i++ {
		if !unicode.IsDigit(rune(text[i])) {
			continue
		}

		cur := NumberInfo{
			Left:  i,
			Right: -1,
			Value: -1,
		}

		for unicode.IsDigit(rune(text[i])) {
			i++
		}

		cur.Right = i
		cur.Value, _ = strconv.Atoi(text[cur.Left:cur.Right])

		border := Surround(cur, rowLength)

		for _, v := range border {
			if 0 < v && v < len(text) && IsSymbol(rune(text[v])) {
				valid = append(valid, cur)
				break
			}
		}
	}

	sum := 0

	for _, v := range valid {
		sum += v.Value
	}

	fmt.Printf("part1: %v\n", sum)
}

func Part2() {
	file, err := os.Open("./problem-3/inputs/prod.txt")

	if err != nil {
		fmt.Printf("err: %v\n", err.Error())
		return
	}

	scanner := bufio.NewScanner(file)

	rowLength := 0

	text := ""

	for scanner.Scan() {
		t := scanner.Text()
		rowLength = len(t)
		text += t
	}

	nums := []NumberInfo{}

	// find numbers
	for i := 0; i < len(text); i++ {
		if !unicode.IsDigit(rune(text[i])) {
			continue
		}

		cur := NumberInfo{
			Left:  i,
			Right: -1,
			Value: -1,
		}

		for unicode.IsDigit(rune(text[i])) {
			i++
		}

		cur.Right = i
		cur.Value, _ = strconv.Atoi(text[cur.Left:cur.Right])

		nums = append(nums, cur)
	}

	// do the problem
	sum := 0
	for i, v := range text {
		if rune(v) != '*' {
			continue
		}

		c := NumberInfo{
			Left:  i,
			Right: i + 1,
			Value: -1,
		}

		border := Surround(c, rowLength)
		gear := make([]int, 0)
		for _, n := range nums {
			for j := n.Left; j < n.Right; j++ {
				if slices.Contains(border, j) {
					gear = append(gear, n.Value)
					break
				}
			}
		}

		if len(gear) == 0 || len(gear) == 1 {
			continue
		}

		product := 1
		for _, v := range gear {
			product *= v
		}

		sum += product
	}

	fmt.Printf("part2: %v\n", sum)
}

func Run() {
	Part1()
	Part2()
}
