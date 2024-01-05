package commands

import (
	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/console/command"
	"github.com/hosseinmirzapur/prisma/helpers"
	"github.com/steebchen/prisma-client-go/cli"
)

type MigrateStatusCommand struct{}

func NewMigrateStatusCommand() *MigrateStatusCommand {
	return &MigrateStatusCommand{}
}

// Signature The name and signature of the console command.
func (receiver *MigrateStatusCommand) Signature() string {
	return "prisma:migrate:status"
}

// Description The console command description.
func (receiver *MigrateStatusCommand) Description() string {
	return "Check the status of your database migrations"
}

// Extend The console command extend.
func (receiver *MigrateStatusCommand) Extend() command.Extend {
	return command.Extend{
		Category: "prisma",
	}
}

// Handle Execute the console command.
func (receiver *MigrateStatusCommand) Handle(ctx console.Context) error {
	return cli.Run(helpers.Commands([]string{"migrate", "status"}, ctx.Arguments()...), true)
}
