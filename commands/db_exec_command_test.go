package commands

import (
	"testing"

	"github.com/hosseinmirzapur/prisma/mocks"

	"github.com/stretchr/testify/assert"
)

func TestDBExecCommand(t *testing.T) {
	// make an instance of db exec command struct
	dbec := &DBExecCommand{}
	mockCtx := &mocks.Context{}

	// no args
	mockCtx.On("Arguments").Return([]string{""}).Once()
	assert.Error(t, dbec.Handle(mockCtx))

	// no --file
	mockCtx.On("Arguments").Return([]string{"--file"}).Once()
	assert.Error(t, dbec.Handle(mockCtx))

	// --stdin without sql in stdin
	mockCtx.On("Arguments").Return([]string{"--stdin"}).Once()
	assert.Error(t, dbec.Handle(mockCtx))

	mockCtx.On("Arguments").Return([]string{"--help"}).Once()
	assert.Nil(t, dbec.Handle(mockCtx))
}
