package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/abhi44mbd/killport/internal/finder"
	"github.com/abhi44mbd/killport/internal/killer"
	"github.com/abhi44mbd/killport/internal/output"
	"github.com/spf13/cobra"
)

var force bool
var yes bool
var dryRun bool

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
		spinner, _ := output.StartSpinner("Searching for processes...")

		processes, err := finder.FindByPort(port)

		spinner.Stop()

		if err != nil {
			fmt.Println(err)
			return
		}

		for _, process := range processes {
			output.Success("Process Found")
			fmt.Printf(
				"\nName:    %s\nPID:     %s\nPort:    %s\nCommand: %s\n\n",
				process.Name,
				process.PID,
				process.Port,
				process.Command,
			)

			if dryRun {
				output.Info("Dry run enabled")
				fmt.Printf("Would terminate PID %s\n", process.PID)
				continue
			}

			if !yes {
				reader := bufio.NewReader(os.Stdin)

				fmt.Print("\nContinue and terminate this process? [y/N]: ")
				input, _ := reader.ReadString('\n')

				input = strings.TrimSpace(strings.ToLower(input))

				if input != "y" {
					output.Warning("Operation cancelled")
					continue
				}
			}

			err = killer.KillProcess(process.PID, force)

			if err != nil {
				output.Error(fmt.Sprintf(
					"Failed to terminate PID %s",
					process.PID,
				))
				continue
			}

			output.Success("Process terminated successfully")
		}
	},
}

func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		panic(err)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&force, "force", "f", false, "Force kill process")
	rootCmd.Flags().BoolVarP(&yes, "yes", "y", false, "Skip confirmation")
	rootCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Preview actions without killing")
}
