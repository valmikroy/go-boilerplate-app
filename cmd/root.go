package cmd

import (
	"os"
	"runtime/debug"
	"strings"

	"github.com/spf13/cobra"
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
