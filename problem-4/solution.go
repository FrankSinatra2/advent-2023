package problem4

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	ID       int
	Revealed []int
	Winning  []int
}

func ScoreCard(c Card) int {
	count := 0
	for _, r := range c.Revealed {
		for _, w := range c.Winning {
			if r == w {
				count += 1
				break
			}
		}
	}

	return int(math.Pow(2.0, float64(count-1)))
}

func WinCardIds(c Card) []int {
	result := []int{}
	id := c.ID

	for _, r := range c.Revealed {
		for _, w := range c.Winning {
			if r == w {
				id++
				result = append(result, id)
				break
			}
		}
	}

	return result
}

// Card #: # # # ... # | # # # ... #
func ParseCard(s string) Card {
	var ID int
	revealed := []int{}
	winning := []int{}

	parts := strings.FieldsFunc(s, func(r rune) bool {
		return r == ':' || r == '|'
	})

	// fmt.Printf("parts: %v\n", parts)

	fmt.Sscanf(parts[0], "Card %d", &ID)

	for _, e := range strings.Fields(parts[1]) {
		asInt, _ := strconv.Atoi(e)
		revealed = append(revealed, asInt)
	}

	for _, e := range strings.Fields(parts[2]) {
		asInt, _ := strconv.Atoi(e)
		winning = append(winning, asInt)
	}

	return Card{
		ID:       ID,
		Revealed: revealed,
		Winning:  winning,
	}
}

func Part1() {
	file, err := os.Open("./problem-4/inputs/prod.txt")

	if err != nil {
		fmt.Printf("err: %v\n", err.Error())
		return
	}

	scanner := bufio.NewScanner(file)

	cards := []Card{}

	for scanner.Scan() {
		t := scanner.Text()
		cards = append(cards, ParseCard(t))
	}

	sum := 0

	for _, c := range cards {
		sum += ScoreCard(c)
	}

	fmt.Printf("part1: %v\n", sum)
}

func Part2() {
	file, err := os.Open("./problem-4/inputs/prod.txt")

	if err != nil {
		fmt.Printf("err: %v\n", err.Error())
		return
	}

	scanner := bufio.NewScanner(file)

	baseCards := []Card{}

	for scanner.Scan() {
		t := scanner.Text()
		baseCards = append(baseCards, ParseCard(t))
	}

	cardCounts := make([]int, len(baseCards))

	for _, c := range baseCards {
		ids := WinCardIds(c)

		cardCounts[c.ID-1]++
		repeat := cardCounts[c.ID-1]

		for k := 0; k < repeat; k++ {
			for _, i := range ids {
				cardCounts[i-1]++
			}
		}
	}

	count := 0
	for _, c := range cardCounts {
		count += c
	}

	fmt.Printf("part2: %v\n", count)
}

func Run() {
	Part1()
	Part2()
}
