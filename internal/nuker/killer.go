package nuker

import (
	"os"
	"syscall"
)

func NukeProcess(pid string, force bool) error {
	process, err := os.FindProcess(toInt(pid))

	if err != nil {
		return err
	}
	// SIGTERM is the default signal for graceful termination.
	signal := syscall.SIGTERM

	if force {
		signal = syscall.SIGKILL
	}

	return process.Signal(signal)
}

func toInt(value string) int {
	var result int

	for _, char := range value {
		result = result*10 + int(char-'0')
	}

	return result
}
