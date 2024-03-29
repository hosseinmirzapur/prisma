package commands

import (
	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/console/command"
	"github.com/hosseinmirzapur/prisma/helpers"
	"github.com/steebchen/prisma-client-go/cli"
)

type DBPushCommand struct{}

func NewDBPushCommand() *DBPullCommand {
	return &DBPullCommand{}
}

// Signature The name and signature of the console command.
func (receiver *DBPushCommand) Signature() string {
	return "prisma:db:pull"
}

// Description The console command description.
func (receiver *DBPushCommand) Description() string {
	return "🙌  Push the state from your Prisma schema to your database (no migrations change)"
}

// Extend The console command extend.
func (receiver *DBPushCommand) Extend() command.Extend {
	return command.Extend{
		Category: "prisma",
	}
}

// Handle Execute the console command
func (r *DBPushCommand) Handle(ctx console.Context) error {
	return cli.Run(helpers.Commands([]string{"db", "push"}, ctx.Arguments()...), true)
}
