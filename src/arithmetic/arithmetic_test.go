package arithmetic

import (
	fmt "fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

//-> copy all annotations from here to add and keep the original annotation
//this tests add
func TestAdd(t *testing.T) {
	assert.Equal(t, add(0,0),  0);
	assert.Equal(t, add(2,0),  2);
	assert.Equal(t, add(2,3),  5);
	assert.Equal(t, add(5,-2), 3);
}


func TestSub(t *testing.T) {
	//-> we add '<' to it looks backwards as to find line 19 and add "only" so we remove the original annotation
	fmt.Print("<this tests sub\n");
	assert.Equal(t, sub(0,0),  0);
	assert.Equal(t, sub(2,0),  2);
	assert.Equal(t, sub(2,3),  -1);
	assert.Equal(t, sub(5,-2), 7);
}

//map to multiply, but by line numbers
//this line+6 tests src/arithmetic.go:11-13
func TestMultiply(t *testing.T) {
	assert.Equal(t, multiply(0,0),  0);
	assert.Equal(t, multiply(2,0),  0);
	assert.Equal(t, multiply(2,3),  6);
	assert.Equal(t, multiply(5,-2), -10);
}

//this tests divide
func TestDivision(t *testing.T) {
	assert.Equal(t, divide(0,1),  0);
	assert.Equal(t, divide(2,2),  1);
	assert.Equal(t, divide(8,4),  2);
	assert.Equal(t, divide(4,0), 0x7FFFFFFF); //this line only tests arithmetic.divide+1+8
}
