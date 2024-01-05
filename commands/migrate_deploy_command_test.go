package commands

import (
	"testing"

	"github.com/goravel/prisma/mocks"
	"github.com/stretchr/testify/assert"
)

func TestMigrateDeployCommand(t *testing.T) {
	ctx := &mocks.Context{}
	mdc := NewMigrateDeployCommand()

	// init prisma
	handleInitPrisma(ctx, t)
	defer removePrisma()

	// init prisma schema
	fillPrismaSchema()

	// no args
	ctx.On("Argument", 0).Return("").Once()
	assert.Nil(t, mdc.Handle(ctx))
}
