package stuff

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//this tests do_stuff
func TestStuff(t *testing.T) {
	assert.Equal(t, do_real_stuff(),  42);
}