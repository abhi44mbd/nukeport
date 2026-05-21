package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "killport",
	Short: "Find and kill processes running on ports",
	Long: `killport is a fast and safe CLI tool
to find and terminate processes
running on specific ports.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide a port")
			return
		}

		port := args[0]

		fmt.Printf("Searching process on port %s...\n", port)
	},
}

func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		panic(err)
	}
}
