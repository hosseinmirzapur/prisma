package commands

import (
	"testing"

	"github.com/goravel/prisma/mocks"
	"github.com/stretchr/testify/assert"
)

func TestVersionCommand(t *testing.T) {
	ctx := &mocks.Context{}
	mdc := NewVersionCommand()

	// no args
	ctx.On("Argument", 0).Return("").Once()
	assert.Nil(t, mdc.Handle(ctx))
}
