package day1_test

import (
	"advent2023/day1"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunPart1_Example(t *testing.T) {
	answer, err := day1.RunPart1("example_1.txt")

	assert.Nil(t, err)
	assert.Equal(t, 142, answer)
}

func TestRunPart1_Input(t *testing.T) {
	answer, err := day1.RunPart1("input_1.txt")

	assert.Nil(t, err)
	assert.Equal(t, 54338, answer)
}
