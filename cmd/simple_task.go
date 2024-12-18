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

	simpleTaskCmd.PersistentFlags().StringVarP(&tasks.FirstCmdParam, "first", "f", "", "First Param")
	simpleTaskCmd.MarkPersistentFlagRequired("first")
	simpleTaskCmd.PersistentFlags().StringVarP(&tasks.SecondCmdParam, "second", "s", "", "Second Param")
	// Register version command
	rootCmd.AddCommand(simpleTaskCmd)
}
