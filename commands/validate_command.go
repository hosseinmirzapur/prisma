package commands

import (
	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/console/command"
	"github.com/hosseinmirzapur/prisma/helpers"
	"github.com/steebchen/prisma-client-go/cli"
)

type ValidateCommand struct{}

func NewValidateCommand() *ValidateCommand {
	return &ValidateCommand{}
}

// Signature The name and signature of the console command.
func (receiver *ValidateCommand) Signature() string {
	return "prisma:validate"
}

// Description The console command description.
func (receiver *ValidateCommand) Description() string {
	return "Validate a Prisma schema."
}

// Extend The console command extend.
func (receiver *ValidateCommand) Extend() command.Extend {
	return command.Extend{
		Category: "prisma",
	}
}

// Handle Execute the console command.
func (receiver *ValidateCommand) Handle(ctx console.Context) error {
	return cli.Run(helpers.Commands([]string{"validate"}, ctx.Arguments()...), true)
}
