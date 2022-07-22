package main

import (
	"fmt"

	"github.com/alecthomas/kong"
	"github.com/jfelipearaujo/gvm/commands"
)

type Globals struct {
	Version VersionFlag `name:"version" help:"Print version information and quit"`
}

type VersionFlag string

func (v VersionFlag) Decode(ctx *kong.DecodeContext) error { return nil }
func (v VersionFlag) IsBool() bool                         { return true }
func (v VersionFlag) BeforeApply(app *kong.Kong, vars kong.Vars) error {
	fmt.Println(vars["version"])
	app.Exit(0)
	return nil
}

type CLI struct {
	Globals

	GoPath    commands.GoPathCommand    `cmd:"" short:"p" help:"Set the GOPATH environment variable"`
	Install   commands.InstallCommand   `cmd:"" short:"i" help:"Install (or reinstall) a valid version of Go Lang"`
	List      commands.ListCommand      `cmd:"" short:"l" help:"List installed all versions of Go Lang"`
	Uninstall commands.UninstallCommand `cmd:"" short:"x" help:"Uninstall a version of Go Lang"`
	Use       commands.UseCommand       `cmd:"" short:"u" help:"Use an installed version of Go Lang"`
}

func main() {
	var cli = CLI{}

	ctx := kong.Parse(&cli,
		kong.Name("gvm"),
		kong.Description("A Go Lang Version Manager"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
		}),
		kong.Vars{
			"version": "1.0.3",
		})
	err := ctx.Run(&cli.Globals)
	ctx.FatalIfErrorf(err)
}
