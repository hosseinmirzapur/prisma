package commands

import (
	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/console/command"
	"github.com/goravel/prisma/helpers"
	"github.com/steebchen/prisma-client-go/cli"
)

type DBExecCommand struct{}

func NewDBExecCommand() *DBExecCommand {
	return &DBExecCommand{}
}

// Signature The name and signature of the console command.
func (receiver *DBExecCommand) Signature() string {
	return "prisma:db:exec"
}

// Description The console command description.
func (receiver *DBExecCommand) Description() string {
	return "🏋️  Manage your database schema and lifecycle during development."
}

// Extend The console command extend.
func (receiver *DBExecCommand) Extend() command.Extend {
	return command.Extend{
		Category: "prisma",
	}
}

// Handle Execute the console command
func (r *DBExecCommand) Handle(ctx console.Context) error {
	return cli.Run(helpers.Commands([]string{"db", "execute"}, ctx.Arguments()...), true)
}
