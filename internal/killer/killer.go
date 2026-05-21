package killer

import (
	"os"
	"syscall"
)

func KillProcess(pid string) error {
	process, err := os.FindProcess(toInt(pid))

	if err != nil {
		return err
	}

	return process.Signal(syscall.SIGTERM)
}

func toInt(value string) int {
	var result int

	for _, char := range value {
		result = result*10 + int(char-'0')
	}

	return result
}
