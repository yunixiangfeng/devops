package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var cmdEcho = &cobra.Command{
	Use:   "echo [string to echo]",
	Short: "Echo anything to the screen",
	Long:  `echo is for echoing anything back. Echo works a lot like print, except it has a child command.`, Args: cobra.MinimumNArgs(1), PreRun: func(cmd *cobra.Command, args []string) {
		setup()
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Print: " + strings.Join(args, " "))
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(cmdEcho)
}

func setup() {
	fmt.Println("setup echo...")
}
