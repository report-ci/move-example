package arithmetic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	assert.Equal(t, add(0,0),  0);
	assert.Equal(t, add(2,0),  2);
	assert.Equal(t, add(2,3),  5);
	assert.Equal(t, add(5,-2), 3);
}

func TestSub(t *testing.T) {
	assert.Equal(t, sub(0,0),  0);
	assert.Equal(t, sub(2,0),  2);
	assert.Equal(t, sub(2,3),  -1);
	assert.Equal(t, sub(5,-2), 7);
}

func TestMultiply(t *testing.T) {
	assert.Equal(t, multiply(0,0),  0);
	assert.Equal(t, multiply(2,0),  0);
	assert.Equal(t, multiply(2,3),  6);
	assert.Equal(t, multiply(5,-2), -10);
}

func TestDivision(t *testing.T) {
	assert.Equal(t, divide(0,1),  0);
	assert.Equal(t, divide(2,2),  1);
	assert.Equal(t, divide(8,4),  2);
	assert.Equal(t, divide(4,-2), -2);
}
