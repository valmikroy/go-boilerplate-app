package cmd

import (
	"fmt"

	"go-boilerplate-app/pkg/utils"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var simpleTaskCmd = &cobra.Command{
	Use:   "simple",
	Short: "Simple Task",
	Long:  `Simple task of printing numbers`,
	Run:   taskCmd,
}

func init() {
	// Register version command
	rootCmd.AddCommand(simpleTaskCmd)
}

func taskCmd(cmd *cobra.Command, args []string) {
	str := utils.String_one() // Initialized and assigned
	fmt.Println(str)

	str = utils.String_two() // now this is just re-assignment
	fmt.Println(str)

	var num int // int variable gets initialized to 0
	fmt.Println(num)

	var flag bool // Boolean variable gets initialized to false
	fmt.Println(flag)

	var MaxInt uint64 = 1<<64 - 1                      // maximum int number (2^64 - 1)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt) // Printing type of variable

	const d = 1e19 // define constatnt
	fmt.Println(d)
	fmt.Println(uint64(d)) // typecasting of constant
}
