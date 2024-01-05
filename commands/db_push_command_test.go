package commands

import (
	"testing"

	"github.com/goravel/prisma/mocks"
	"github.com/stretchr/testify/assert"
)

func TestDBPushCommand(t *testing.T) {
	// make an instance of db pull command struct
	dbpc := NewDBPushCommand()
	mockCtx := &mocks.Context{}

	// init prisma project
	handleInitPrisma(mockCtx, t)
	defer removePrisma()

	// fill schema.prisma with data
	fillPrismaSchema()

	// requires at least one existing table at database to execute push
	mockCtx.On("Argument", 0).Return("").Once()
	assert.Error(t, dbpc.Handle(mockCtx))

}
