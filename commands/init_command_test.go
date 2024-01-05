package commands

import (
	"os"
	"testing"

	"github.com/gookit/color"
	"github.com/goravel/prisma/mocks"
	"github.com/stretchr/testify/assert"
)

func TestInitCommand(t *testing.T) {
	mockCtx := &mocks.Context{}

	// init prisma project
	handleInitPrisma(mockCtx, t)
	defer removePrisma()

	// check directories created
	assert.DirExists(t, "prisma")
	assert.FileExists(t, "prisma/schema.prisma")
	assert.FileExists(t, ".env")
	assert.FileExists(t, ".gitignore")
}

func handleInitPrisma(ctx *mocks.Context, t *testing.T) {
	ic := NewInitCommand()

	ctx.On("Arguments").Return([]string{"--datasource-provider=sqlite", "--url=file:dev.db"}).Once()
	assert.NoError(t, ic.Handle(ctx))
	os.Create("prisma/dev.db")
}

func removePrisma() {
	os.RemoveAll("prisma")
	os.Remove(".env")
	os.Remove(".gitignore")
}

func fillPrismaSchema() {
	file, err := os.OpenFile("prisma/schema.prisma", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		color.Redln("unable to open schema.prisma")
		os.Exit(1)
	}
	defer file.Close()

	txt := `

model User {
	id    Int     @id @default(autoincrement())
	email String  @unique
	name  String?
}
`
	if _, err = file.WriteString(txt); err != nil {
		color.Redln("unable to write to schema.prisma")
		os.Exit(1)
	}
}
