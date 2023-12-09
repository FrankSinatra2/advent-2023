package problem5

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

var types = []string{
	"seed",
	"soil",
	"fertilizer",
	"water",
	"light",
	"temperature",
	"humidity",
	"location",
}

type Point struct {
	FromType string
	FromID   int

	ToType string
	ToID   int
}

type Range struct {
	Source int
	Dest   int
	R      int
}

func (p Point) FromKey() string {
	return fmt.Sprintf("%v%v", p.FromType, p.FromID)
}

func (p Point) ToKey() string {
	return fmt.Sprintf("%v%v", p.ToType, p.ToID)
}

func ParseStartingSeeds(s string) []int {
	parts := strings.Fields(s)[1:]
	result := make([]int, len(parts))
	for i, v := range parts {
		result[i], _ = strconv.Atoi(v)
	}

	return result
}

func ParsePointMap(s string) []Range {
	// fmt.Printf("s: %v\n", s)
	chunk := strings.Split(s, "\n")

	result := []Range{}

	for _, v := range chunk[1 : len(chunk)-1] {
		var source int
		var dest int
		var r int

		fmt.Sscanf(v, "%d %d %d", &dest, &source, &r)
		result = append(result, Range{Source: source, Dest: dest, R: r})

		// fmt.Printf("i: %v\n", i)
	}

	return result
}

func ParseFile(path string) ([]int, map[string][]Range) {
	file, err := os.Open(path)

	if err != nil {
		fmt.Printf("err: %v\n", err.Error())
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	seeds := ParseStartingSeeds(scanner.Text())
	max := 0
	for _, s := range seeds {
		if s > max {
			max = s
		}
	}
	scanner.Scan()
	m := make(map[string][]Range)

	// fmt.Printf("seeds: %v\n", seeds)

	c := ""
	typeIndex := 0
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			p := ParsePointMap(c)
			m[types[typeIndex]] = p

			c = ""
			typeIndex++
			continue
		}

		c += line
		c += "\n"
	}

	return seeds, m
}

func GetID(s string) int {
	var i int
	for i = 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			break
		}
	}
	n, _ := strconv.Atoi(s[i:])

	return n
}

func Part1() {
	s, m := ParseFile("./problem-5/input/prod.txt")

	fmt.Println("Starting 1")

	results := []int{}
	for _, seed := range s {

		search := seed

		for typeIndex := 0; typeIndex < len(types); typeIndex++ {
			ranges := m[types[typeIndex]]

			for _, r := range ranges {
				if r.Source <= search && search < r.Source+r.R {
					diff := int(math.Abs(float64(search - r.Source)))
					search = r.Dest + diff
					break
				}
			}
		}
		results = append(results, search)
	}
	fmt.Printf("result: %d\n", slices.Min(results))
}

// slow
func Part2() {
	s, m := ParseFile("./problem-5/input/prod.txt")
	fmt.Println("Starting 2")
	lowest := -1
	for i := 0; i < len(s); i += 2 {
		for seed := s[i]; seed < s[i]+s[i+1]; seed++ {
			search := seed
			for typeIndex := 0; typeIndex < len(types); typeIndex++ {
				ranges := m[types[typeIndex]]

				for _, r := range ranges {
					if r.Source <= search && search < r.Source+r.R {

						diff := int(math.Abs(float64(search - r.Source)))
						search = r.Dest + diff
						break
					}
				}
			}

			if lowest == -1 || search < lowest {
				lowest = search
			}
		}
	}

	fmt.Printf("result: %d\n", lowest)
}
func Run() {
	Part1()
	Part2()
}
