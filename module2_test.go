package module2_test

import (
	"module2"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	actual := module2.Hello("Bundit")
	assert.Equal(t, "Hello Bundit", actual)
}
