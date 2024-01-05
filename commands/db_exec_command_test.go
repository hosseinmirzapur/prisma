package commands

import (
	"testing"

	"github.com/goravel/prisma/mocks"

	"github.com/stretchr/testify/assert"
)

func TestDBExecCommand(t *testing.T) {
	// make an instance of db exec command struct
	dbec := &DBExecCommand{}
	mockCtx := &mocks.Context{}

	// no args
	mockCtx.On("Argument", 0).Return("").Once()
	assert.Error(t, dbec.Handle(mockCtx))

	// no --file
	mockCtx.On("Argument", 0).Return("--file").Once()
	assert.Error(t, dbec.Handle(mockCtx))

	// --stdin without sql in stdin
	mockCtx.On("Argument", 0).Return("--stdin").Once()
	assert.Error(t, dbec.Handle(mockCtx))

	mockCtx.On("Argument", 0).Return("-h").Once()
	assert.Nil(t, dbec.Handle(mockCtx))
}
