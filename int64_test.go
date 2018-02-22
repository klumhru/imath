package mathy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax64(t *testing.T) {
	assert := assert.New(t)

	expected := int64(3)
	assert.Equal(expected, Max64(2, 3))

	assert.Equal(expected, Max64(3, 2))
}

func TestMin64(t *testing.T) {
	assert := assert.New(t)

	expected := int64(3)

	assert.Equal(expected, Min64(5, 3))

	assert.Equal(expected, Min64(3, 5))
}

func TestClamp64(t *testing.T) {
	assert := assert.New(t)

	_, err := Clamp64(4, 3, 2)
	assert.Error(err)

	actual, _ := Clamp64(3, 5, 10)
	assert.Equal(int64(5), actual)

	actual, _ = Clamp64(12, 5, 10)
	assert.Equal(int64(10), actual)

	actual, _ = Clamp64(7, 5, 10)
	assert.Equal(int64(7), actual)
}

func TestAbs64(t *testing.T) {
	assert := assert.New(t)

	expected := int64(3)
	assert.Equal(expected, Abs64(-3))

	expected = int64(3)
	assert.Equal(expected, Abs64(3))
}
