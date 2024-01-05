package commands

import (
	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/console/command"
	"github.com/goravel/prisma/helpers"
	"github.com/steebchen/prisma-client-go/cli"
)

type MigrateResetCommand struct{}

func NewMigrateResetCommand() *MigrateResetCommand {
	return &MigrateResetCommand{}
}

// Signature The name and signature of the console command.
func (receiver *MigrateResetCommand) Signature() string {
	return "prisma:migrate:reset"
}

// Description The console command description.
func (receiver *MigrateResetCommand) Description() string {
	return "Reset your database and apply all migrations, all data will be lost"
}

// Extend The console command extend.
func (receiver *MigrateResetCommand) Extend() command.Extend {
	return command.Extend{
		Category: "prisma",
	}
}

// Handle Execute the console command.
func (receiver *MigrateResetCommand) Handle(ctx console.Context) error {
	return cli.Run(helpers.Commands([]string{"migrate", "reset"}, ctx.Arguments()...), true)
}
