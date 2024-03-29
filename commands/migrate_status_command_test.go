package commands

import (
	"testing"

	"github.com/hosseinmirzapur/prisma/mocks"
	"github.com/stretchr/testify/assert"
)

func TestMigrateStatusCommand(t *testing.T) {
	ctx := &mocks.Context{}
	mdc := NewMigrateStatusCommand()

	// init prisma
	handleInitPrisma(ctx, t)
	defer removePrisma()

	// fill prisma schema
	fillPrismaSchema()

	// migrate dev to get database tables ready
	migrateDev(ctx)

	// no args
	ctx.On("Arguments").Return([]string{""}).Once()
	assert.Nil(t, mdc.Handle(ctx))
}
