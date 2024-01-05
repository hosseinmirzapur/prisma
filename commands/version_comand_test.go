package commands

import (
	"testing"

	"github.com/hosseinmirzapur/prisma/mocks"
	"github.com/stretchr/testify/assert"
)

func TestVersionCommand(t *testing.T) {
	ctx := &mocks.Context{}
	mdc := NewVersionCommand()

	// no args
	ctx.On("Arguments").Return([]string{""}).Once()
	assert.Nil(t, mdc.Handle(ctx))
}
