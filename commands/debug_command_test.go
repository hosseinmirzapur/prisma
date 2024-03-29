package commands

import (
	"testing"

	"github.com/hosseinmirzapur/prisma/mocks"
	"github.com/stretchr/testify/assert"
)

func TestDebugCommand(t *testing.T) {
	debugCmd := NewDebugCommand()
	mockCtx := &mocks.Context{}

	// init prisma
	handleInitPrisma(mockCtx, t)
	defer removePrisma()

	// check debug info
	mockCtx.On("Arguments").Return([]string{}).Once()
	assert.Nil(t, debugCmd.Handle(mockCtx))
}
