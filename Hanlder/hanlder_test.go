package hanlder

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestAddFunc(t *testing.T) {
	assert.Equal(t, AddFunc(1, 2), 3, "should equal")
}
