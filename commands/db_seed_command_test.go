package commands

import (
	"testing"

	"github.com/hosseinmirzapur/prisma/mocks"
	"github.com/stretchr/testify/assert"
)

func TestDBSeedCommand(t *testing.T) {
	dbsc := NewDBSeedCommand()
	mockCtx := &mocks.Context{}

	// init prisma before any test
	handleInitPrisma(mockCtx, t)

	// fill schema.prisma with data
	fillPrismaSchema()
	defer removePrisma()

	// test on user model
	mockCtx.On("Arguments").Return([]string{""})
	assert.Nil(t, dbsc.Handle(mockCtx))

}
