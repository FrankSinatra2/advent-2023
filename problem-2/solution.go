package problem2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type GameState struct {
	ID         int
	RedBalls   int
	BlueBalls  int
	GreenBalls int
}

/*
* Game #: # blue, # red, # green; # blue, # green
*
 */
func ParseGameStates(s string) []GameState {
	result := make([]GameState, 0)

	tokens := strings.FieldsFunc(s, func(r rune) bool {
		return r == ':' || r == ';'
	})

	var ID int

	fmt.Sscanf(tokens[0], "Game %d", &ID)
	// fmt.Printf("ID: %v\n", ID)

	for _, value := range tokens[1:] {
		var red int
		var blue int
		var green int

		for _, countColor := range strings.Split(value, ",") {
			var count int
			var color string

			fmt.Sscanf(countColor, "%d %s", &count, &color)

			switch color {
			case "red":
				red = count
			case "blue":
				blue = count
			case "green":
				green = count
			default:
				fmt.Printf("Failed to parse: %s", countColor)
			}
		}

		result = append(result, GameState{
			ID:         ID,
			RedBalls:   red,
			BlueBalls:  blue,
			GreenBalls: green,
		})

		// fmt.Printf("%d red, %d blue, %d green\n", red, blue, green)
	}

	return result
}

func Part1() {
	file, err := os.Open("./problem-2/inputs/prod.txt")

	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	maxRed := 12
	maxBlue := 14
	maxGreen := 13

	sum := 0

	for scanner.Scan() {
		games := ParseGameStates(scanner.Text())
		valid := true

		for _, g := range games {
			if g.RedBalls > maxRed || g.BlueBalls > maxBlue || g.GreenBalls > maxGreen {
				valid = false
				break
			}
		}

		if valid {
			sum += games[0].ID
		}
	}

	fmt.Printf("sum: %v\n", sum)
}

func Part2() {
	file, err := os.Open("./problem-2/inputs/prod.txt")

	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		games := ParseGameStates(scanner.Text())

		maxRed := 0
		maxBlue := 0
		maxGreen := 0

		for _, g := range games {
			if g.RedBalls > maxRed {
				maxRed = g.RedBalls
			}

			if g.BlueBalls > maxBlue {
				maxBlue = g.BlueBalls
			}
			if g.GreenBalls > maxGreen {
				maxGreen = g.GreenBalls
			}
		}

		sum += maxRed * maxBlue * maxGreen
	}

	fmt.Printf("sum: %v\n", sum)

}

func Run() {
	Part1()
	Part2()
}
