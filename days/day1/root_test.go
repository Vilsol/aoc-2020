package day1

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
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
