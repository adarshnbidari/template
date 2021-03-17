package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "template",
	Short: "Create Template for a Project",
	Long:  `Template command line tool creates set's up the intial project`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
