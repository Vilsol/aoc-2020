package registry

import (
	"errors"
	"fmt"
	"io/ioutil"
)

type Runner func(input []byte) (*Results, error)

type Results struct {
	Part1 interface{}
	Part2 interface{}
}

var registry = make(map[uint8]Runner)

func RegisterDay(day uint8, runner Runner) {
	registry[day] = runner
}

func RunDay(day uint8) error {
	if runner, ok := registry[day]; ok {
		input, err := ioutil.ReadFile(fmt.Sprintf("./inputs/day%d.txt", day))

		if err != nil {
			return err
		}

		if _, err := runner(input); err != nil {
			return err
		}
	} else {
		return errors.New("day not found")
	}

	return nil
}
