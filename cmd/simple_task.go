package cmd

import (
	"go-boilerplate-app/pkg/tasks"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var simpleTaskCmd = &cobra.Command{
	Use:   "simple",
	Short: "Simple Task",
	Long:  `Simple task of printing numbers`,
	Run:   tasks.SimpleTask,
}

func init() {

	// Persistant flag
	simpleTaskCmd.PersistentFlags().StringVarP(&tasks.FirstCmdParam, "first", "f", "", "First Param")
	// Required of persistant flag
	simpleTaskCmd.MarkPersistentFlagRequired("first")

	// Non-Persistant flag - will not passed down to subcommand
	simpleTaskCmd.Flags().StringVarP(&tasks.SecondCmdParam, "second", "s", "", "Second Param")
	// Register version command
	rootCmd.AddCommand(simpleTaskCmd)
}
