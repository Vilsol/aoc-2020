package day1

import (
	"bytes"
	"github.com/Vilsol/aoc-2020/registry"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

const expectedResult = 2020

func init() {
	registry.RegisterDay(1, Day1)
}

func Day1(input []byte) (*registry.Results, error) {
	var err error
	lines := bytes.Split(bytes.TrimSpace(input), []byte("\n"))
	count := len(lines)
	numbers := make([]int, count)
	for i, line := range lines {
		numbers[i], err = strconv.Atoi(strings.TrimSpace(string(line)))
		if err != nil {
			return nil, err
		}
	}

	var part1 int
	if part1, err = Part1(numbers); err != nil {
		return nil, err
	}

	var part2 int
	if part2, err = Part2(numbers); err != nil {
		return nil, err
	}

	return &registry.Results{
		Part1: part1,
		Part2: part2,
	}, nil
}

func Part1(numbers []int) (int, error) {
	count := len(numbers)
	for i := 0; i < count-1; i++ {
		for j := i + 1; j < count; j++ {
			if numbers[i]+numbers[j] == expectedResult {
				log.WithFields(map[string]interface{}{
					"first":   numbers[i],
					"second":  numbers[j],
					"sum":     numbers[i] + numbers[j],
					"product": numbers[i] * numbers[j],
				}).Info("Day 1 Part 1 Result")
				return numbers[i] * numbers[j], nil
			}
		}
	}

	log.Error("Day 1 Part 1 Result not found")

	return 0, nil
}

func Part2(numbers []int) (int, error) {
	count := len(numbers)
	for i := 0; i < count-2; i++ {
		for j := i + 1; j < count; j++ {
			if numbers[i]+numbers[j] < 2020 {
				for z := j + 1; z < count; z++ {
					if numbers[i]+numbers[j]+numbers[z] == expectedResult {
						log.WithFields(map[string]interface{}{
							"first":   numbers[i],
							"second":  numbers[j],
							"third":   numbers[z],
							"sum":     numbers[i] + numbers[j] + numbers[z],
							"product": numbers[i] * numbers[j] * numbers[z],
						}).Info("Day 1 Part 2 Result")
						return numbers[i] * numbers[j] * numbers[z], nil
					}
				}
			}
		}
	}

	log.Error("Day 1 Part 2 Result not found")

	return 0, nil
}
