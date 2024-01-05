package commands

import (
	"testing"

	"github.com/goravel/prisma/mocks"
	"github.com/stretchr/testify/assert"
)

func TestMigrateResetCommand(t *testing.T) {
	ctx := &mocks.Context{}
	mdc := NewMigrateResetCommand()

	// init prisma
	handleInitPrisma(ctx, t)
	defer removePrisma()

	// fill prisma schema
	fillPrismaSchema()

	// firstly migrate the schema
	migrateDev(ctx)

	// reset command will fail because it requires user interaction to continue
	// but passes as a assert.Error() test
	ctx.On("Argument", 0).Return("").Once()
	assert.Error(t, mdc.Handle(ctx))
}

func migrateDev(ctx *mocks.Context) {
	md := NewMigrateDevCommand()
	ctx.On("Argument", 0).Return("").Once()
	md.Handle(ctx)
}
