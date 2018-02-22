package imath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	assert := assert.New(t)

	expected := 3
	assert.Equal(expected, Max(2, 3))

	assert.Equal(expected, Max(3, 2))
}

func TestMin(t *testing.T) {
	assert := assert.New(t)

	expected := 3

	assert.Equal(expected, Min(5, 3))

	assert.Equal(expected, Min(3, 5))
}

func TestClamp(t *testing.T) {
	assert := assert.New(t)

	_, err := Clamp(4, 3, 2)
	assert.Error(err)

	actual, _ := Clamp(3, 5, 10)
	assert.Equal(5, actual)

	actual, _ = Clamp(12, 5, 10)
	assert.Equal(10, actual)

	actual, _ = Clamp(7, 5, 10)
	assert.Equal(7, actual)
}

func TestAbs(t *testing.T) {
	assert := assert.New(t)

	expected := 3
	assert.Equal(expected, Abs(-3))

	expected = 3
	assert.Equal(expected, Abs(3))
}
