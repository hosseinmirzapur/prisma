package commands

import (
	"testing"

	"github.com/goravel/prisma/mocks"
	"github.com/stretchr/testify/assert"
)

func TestValidateCommand(t *testing.T) {
	ctx := &mocks.Context{}
	mdc := NewValidateCommand()

	// init prisma
	handleInitPrisma(ctx, t)
	defer removePrisma()

	// create prisma schema
	fillPrismaSchema()

	// no args
	ctx.On("Arguments").Return([]string{""}).Once()
	assert.Nil(t, mdc.Handle(ctx))
}
