package prisma

import (
	"github.com/goravel/framework/contracts/foundation"
	"github.com/goravel/prisma/commands"

	consolecontract "github.com/goravel/framework/contracts/console"
)

var App foundation.Application

type ServiceProvider struct{}

func (prisma *ServiceProvider) Register(app foundation.Application) {
	App = app
}

func (prisma *ServiceProvider) Boot(app foundation.Application) {
	prisma.registerCommands(app)
}

// register prisma commands
func (prisma *ServiceProvider) registerCommands(app foundation.Application) {
	app.MakeArtisan().Register([]consolecontract.Command{
		commands.NewDBExecCommand(),
		commands.NewDBPullCommand(),
		commands.NewDBPushCommand(),
		commands.NewDBSeedCommand(),
		commands.NewDebugCommand(),
		commands.NewFormatCommand(),
		commands.NewGenerateCommand(),
		commands.NewInitCommand(),
		commands.NewMigrateDeployCommand(),
		commands.NewMigrateDevCommand(),
		commands.NewMigrateResetCommand(),
		commands.NewMigrateStatusCommand(),
		commands.NewValidateCommand(),
		commands.NewVersionCommand(),
	})
}
