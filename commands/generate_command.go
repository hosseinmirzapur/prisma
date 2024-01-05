package commands

import (
	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/console/command"
	"github.com/hosseinmirzapur/prisma/helpers"
	"github.com/steebchen/prisma-client-go/cli"
)

type GenerateCommand struct{}

func NewGenerateCommand() *GenerateCommand {
	return &GenerateCommand{}
}

// Signature The name and signature of the console command.
func (receiver *GenerateCommand) Signature() string {
	return "prisma:generate"
}

// Description The console command description.
func (receiver *GenerateCommand) Description() string {
	return "Generate artifacts (e.g. Prisma Client)"
}

// Extend The console command extend.
func (receiver *GenerateCommand) Extend() command.Extend {
	return command.Extend{
		Category: "prisma",
	}
}

// Handle Execute the console command
func (r *GenerateCommand) Handle(ctx console.Context) error {
	return cli.Run(helpers.Commands([]string{"generate"}, ctx.Arguments()...), true)
}
