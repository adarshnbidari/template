package cmd

import (
	c "github.com/adarshnbidari/template/components"
	"github.com/spf13/cobra"
)

var projectType string

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "command to initialize project",
	RunE: func(cmd *cobra.Command, args []string) error {

		if err := c.Setup(projectType); err != nil {

			return err

		}

		return nil
	},
}

func init() {

	RootCmd.AddCommand(initCmd)

	initCmd.Flags().StringVarP(&projectType, "project", "p", "", "Name of the Project to create")
	initCmd.MarkFlagRequired("project")

}
