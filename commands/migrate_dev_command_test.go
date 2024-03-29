package commands

import (
	"testing"

	"github.com/hosseinmirzapur/prisma/mocks"
	"github.com/stretchr/testify/assert"
)

func TestMigrateDevCommand(t *testing.T) {
	ctx := &mocks.Context{}
	mdc := NewMigrateDevCommand()

	// init prisma
	handleInitPrisma(ctx, t)
	defer removePrisma()

	// fill prisma schema
	fillPrismaSchema()

	// --name unspecified
	ctx.On("Arguments").Return([]string{"--name", "init"}).Once()
	assert.Nil(t, mdc.Handle(ctx))

	assert.DirExists(t, "prisma/migrations")
}
