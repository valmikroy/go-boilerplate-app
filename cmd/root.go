package cmd

import (
	"os"
	"runtime/debug"
	"strings"

	"github.com/spf13/cobra"

	"go-boilerplate-app/pkg/config"
	"go-boilerplate-app/pkg/logger"
)

var Version string
var runtimeInfo, _ = debug.ReadBuildInfo()
var runtimeModuleInfo = strings.Split(runtimeInfo.Main.Path, "/")
var runtimeModuleName = runtimeModuleInfo[len(runtimeModuleInfo)-1]

var rootCmd = &cobra.Command{
	Use:   runtimeModuleName,
	Short: runtimeModuleName + " cli",
	Long: `
____   ____      .__          .__ __     __________              
\   \ /   /____  |  |   _____ |__|  | __ \______   \ ____ ___.__.
 \   Y   /\__  \ |  |  /     \|  |  |/ /  |       _//  _ <   |  |
  \     /  / __ \|  |_|  Y Y  \  |    <   |    |   (  <_> )___  |
   \___/  (____  /____/__|_|  /__|__|_ \  |____|_  /\____// ____|
               \/           \/        \/         \/       \/   
` + runtimeModuleName + ` cli`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	rootCmd.PersistentFlags().BoolVarP(&config.DevModeFlag, "dev", "d", false, "Run in development mode")
	rootCmd.PersistentFlags().BoolVarP(&config.EnvModeFlag, "env", "e", false, "Print environment before execution")
	rootCmd.PersistentFlags().StringVarP(&config.LogLevelFlag, "log", "l", "", "Log Level")

	cobra.OnInitialize(initEnv)

	// Register persistent function for all commands
	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		execRootPersistentPreRun()
	}
}

func initEnv() {
	if config.DevModeFlag {
		// Changing to Dev log level
		logger.SetDevMode()
	}
}

func execRootPersistentPreRun() {
	logger.Debug("Executing root cmd persistent pre run ...")

	// You can initialize other features here ...
	// this will run before any command, make sure to put only global initializations here
	// to avoid running into nil pointers or undefined variables
	// ...

}
