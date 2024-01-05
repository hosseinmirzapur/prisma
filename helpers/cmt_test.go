package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommands(t *testing.T) {
	// data to provide
	dataToProvide := []string{"1", "2", "3", "4", "5"}

	// data to expect
	dataToExpect := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}

	// test functionalities
	assert.Equal(t, dataToExpect, Commands(dataToProvide, "6", "7", "8", "9", "10"))
}
