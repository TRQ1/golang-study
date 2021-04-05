package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSixNumberTrue(t *testing.T) {
	p := IsValidLottery(1, 2, 3, 4, 5, 6)
	assert.True(t, p, "There are the count of six numbers")
}

func TestUnderSixNumberFalse(t *testing.T) {
	p := IsValidLottery(1, 2, 3, 4, 5)
	assert.False(t, p, "It is under six numbers")
}

func TestSixOverNumberFalse(t *testing.T) {
	p := IsValidLottery(1, 2, 3, 4, 5, 6, 7)
	assert.False(t, p, "It is over six numbers")
}

func TestRangeOverNumber(t *testing.T) {
	p := IsValidLottery(0, 2, 3, 4, 5, 45)
	assert.False(t, p, "It is over a rang of numbers from 1 to 45")
}

func TestSortNumberTest(t *testing.T) {
	p := IsValidLottery(1, 12, 3, 20, 5, 7)
	assert.False(t, p, "It is not to sort the numbers in ascending order")
}

func TestDuplicatedNumber(t *testing.T) {
	p := IsValidLottery(1, 2, 2, 4, 5, 6)
	assert.False(t, p, "It is the sort the numbes in ascending order but not duplication")
}
