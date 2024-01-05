package commands

import (
	"testing"

	"github.com/hosseinmirzapur/prisma/mocks"
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
	ctx.On("Arguments").Return([]string{"--force"}).Once()
	assert.Nil(t, mdc.Handle(ctx))
}

func migrateDev(ctx *mocks.Context) {
	md := NewMigrateDevCommand()
	ctx.On("Arguments").Return([]string{"--name", "init"}).Once()
	md.Handle(ctx)
}
