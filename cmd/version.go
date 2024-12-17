package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print service version",
	Long:  `Print service version`,
	Run:   execVersionCmd,
}

func init() {
	// Register version command
	rootCmd.AddCommand(versionCmd)
}

func execVersionCmd(cmd *cobra.Command, args []string) {
	// Version is set in main.go
	fmt.Printf("v%s\n", Version)
}
