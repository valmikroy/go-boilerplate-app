package cmd

import (
	"go-boilerplate-app/pkg/tasks"

	"go-boilerplate-app/pkg/config"
	"go-boilerplate-app/pkg/utils/constants"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var startServiceCmd = &cobra.Command{
	Use:   "service",
	Short: "Simple Service Task",
	Long:  `Simple Service Task of Epoc`,
	Run:   tasks.SimpleServiceTask,
}

func init() {

	startServiceCmd.PersistentFlags().StringVarP(&config.HostFlag, "host", "H", constants.DEFAULT_HOST, "service host")
	startServiceCmd.PersistentFlags().StringVarP(&config.PortFlag, "port", "P", constants.DEFAULT_PORT, "service port")

	// Register version command to upstream rootCmd
	rootCmd.AddCommand(startServiceCmd)
}
