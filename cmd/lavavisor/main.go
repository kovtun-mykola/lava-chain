package main

import (
	"fmt"
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/lavanet/lava/app"
	"github.com/lavanet/lava/cmd/lavad/cmd"
	lvcmd "github.com/lavanet/lava/ecosystem/lavavisor/cmd"
	"github.com/lavanet/lava/protocol/upgrade"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := cmd.NewLavaVisorRootCmd()

	// version cobra command
	cmdVersion := versionCommand()
	// lavavisor init cobra command
	cmdLavavisorInit := lvcmd.CreateLavaVisorInitCobraCommand()
	// lavavisor start cobra command
	cmdLavavisorStart := lvcmd.CreateLavaVisorStartCobraCommand()
	// lavavisor start cobra command
	cmdLavavisorCreateService := lvcmd.CreateLavaVisorCreateServiceCobraCommand()

	// Add Version Command
	rootCmd.AddCommand(cmdVersion)
	// Add Lavavisor Init
	rootCmd.AddCommand(cmdLavavisorInit)
	// Add Lavavisor Start
	rootCmd.AddCommand(cmdLavavisorStart)
	// Add Lavavisor Create Service
	rootCmd.AddCommand(cmdLavavisorCreateService)

	if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}

func versionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number",
		Run: func(cmd *cobra.Command, args []string) {
			// Print the lavap version
			version := upgrade.GetCurrentVersion()
			fmt.Println(version.ProviderVersion) // currently we have only one version.
		},
	}
}
