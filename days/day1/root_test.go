package day1

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func TestDay1(t *testing.T) {
	input, err := ioutil.ReadFile("../../inputs/day1.txt")

	if err != nil {
		t.Error(err)
		return
	}

	results, err := Day1(input)

	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, 633216, results.Part1)
	assert.Equal(t, 68348924, results.Part2)
}

func prepareData() ([]int, error) {
	input, err := ioutil.ReadFile("../../inputs/day1.txt")

	if err != nil {
		return nil, err
	}

	lines := bytes.Split(bytes.TrimSpace(input), []byte("\n"))
	count := len(lines)
	numbers := make([]int, count)
	for i, line := range lines {
		numbers[i], err = strconv.Atoi(strings.TrimSpace(string(line)))
		if err != nil {
			return nil, err
		}
	}

	return numbers, nil
}

func BenchmarkUnsortedPart1(b *testing.B) {
	log.SetLevel(0)
	numbers, err := prepareData()
	if err != nil {
		b.Error(err)
		return
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = Part1(numbers)
	}
}

func BenchmarkSortedAscendingPart1(b *testing.B) {
	log.SetLevel(0)
	numbers, err := prepareData()
	if err != nil {
		b.Error(err)
		return
	}

	sort.Ints(numbers)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = Part1(numbers)
	}
}

func BenchmarkSortedDescendingPart1(b *testing.B) {
	log.SetLevel(0)
	numbers, err := prepareData()
	if err != nil {
		b.Error(err)
		return
	}

	sort.Ints(numbers)
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - 1 - i
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = Part1(numbers)
	}
}

func BenchmarkUnsortedPart2(b *testing.B) {
	log.SetLevel(0)
	numbers, err := prepareData()
	if err != nil {
		b.Error(err)
		return
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = Part2(numbers)
	}
}

func BenchmarkSortedAscendingPart2(b *testing.B) {
	log.SetLevel(0)
	numbers, err := prepareData()
	if err != nil {
		b.Error(err)
		return
	}

	sort.Ints(numbers)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = Part2(numbers)
	}
}

func BenchmarkSortedDescendingPart2(b *testing.B) {
	log.SetLevel(0)
	numbers, err := prepareData()
	if err != nil {
		b.Error(err)
		return
	}

	sort.Ints(numbers)
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - 1 - i
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = Part2(numbers)
	}
}
