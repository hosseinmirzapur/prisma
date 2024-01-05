package commands

import (
	"testing"

	"github.com/goravel/prisma/mocks"
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
	ctx.On("Argument", 0).Return("")
	assert.Nil(t, mdc.Handle(ctx))

	assert.DirExists(t, "prisma/migrations")
}
