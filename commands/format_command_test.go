package commands

import (
	"testing"

	"github.com/goravel/prisma/mocks"
	"github.com/stretchr/testify/assert"
)

func TestFormatCommand(t *testing.T) {
	fmtCmd := NewFormatCommand()
	ctx := &mocks.Context{}

	// init prisma project
	handleInitPrisma(ctx, t)
	defer removePrisma()

	// fill schema.prisma with data
	fillPrismaSchema()

	// formatting the prisma schema file has no error
	ctx.On("Arguments").Return([]string{}).Once()
	assert.Nil(t, fmtCmd.Handle(ctx))
}
