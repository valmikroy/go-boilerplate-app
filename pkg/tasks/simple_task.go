package tasks

import (
	"fmt"

	"github.com/spf13/cobra"
)

var FirstCmdParam string
var SecondCmdParam string

func SimpleTask(cmd *cobra.Command, args []string) {
	str := String_one(FirstCmdParam) // Initialized and assigned
	fmt.Println(str)

	str = String_two(SecondCmdParam) // now this is just re-assignment
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

func String_one(a string) string {
	ret := "Hello " + a // initialize and assign a string
	return ret
}

func String_two(a string) string {
	var ret = "Hello World " + a // this is another way of initalizing variable compare to line 10 above
	// along with string concatnation example
	return ret
}
