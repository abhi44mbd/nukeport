package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/abhi44mbd/killport/internal/finder"
	"github.com/abhi44mbd/killport/internal/killer"
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

		processes, err := finder.FindByPort(port)

		if err != nil {
			fmt.Println(err)
			return
		}

		for _, process := range processes {
			fmt.Println("Process found")
			fmt.Printf("Name: %s\n", process.Name)
			fmt.Printf("PID: %s\n", process.PID)
			fmt.Printf("Port: %s\n", process.Port)
			fmt.Printf("Command: %s\n", process.Command)

			reader := bufio.NewReader(os.Stdin)

			fmt.Print("Kill this process? (y/n): ")

			input, _ := reader.ReadString('\n')

			input = strings.TrimSpace(strings.ToLower(input))

			if input != "y" {
				fmt.Println("Operation cancelled")
				continue
			}

			err = killer.KillProcess(process.PID)

			if err != nil {
				fmt.Println("Failed to kill process")
				continue
			}

			fmt.Println("Process terminated successfully")
		}
	},
}

func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		panic(err)
	}
}
