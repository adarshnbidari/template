package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Version: "v0.0.1",
	Use:     "template",
	Short:   "Create Template for a Project",
	Long:    `Template command line tool creates set's up the initial project`,
}

//Execute to execute the commands
func Execute() {
	cobra.CheckErr(RootCmd.Execute())
}
