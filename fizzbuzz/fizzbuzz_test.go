package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoFizzNoBuzz(t *testing.T) {
	assert.Equal(t, FizzBuzz(1), "1")
	assert.Equal(t, FizzBuzz(2), "2")
	assert.Equal(t, FizzBuzz(4), "4")
}

func TestFizzNoBuzz(t *testing.T) {
	assert.Equal(t, FizzBuzz(3), "Fizz")
	assert.Equal(t, FizzBuzz(9), "Fizz")
	assert.Equal(t, FizzBuzz(12), "Fizz")
}

func TestNoFizzBuzz(t *testing.T) {
	assert.Equal(t, FizzBuzz(5), "Buzz")
	assert.Equal(t, FizzBuzz(10), "Buzz")
}

func TestFizzBuzz(t *testing.T) {
	assert.Equal(t, FizzBuzz(15), "FizzBuzz")
}
