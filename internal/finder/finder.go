package finder

import (
	"fmt"
	"os/exec"
	"strings"
)

type Process struct {
	PID     string
	Name    string
	Command string
	Port    string
}

func FindByPort(port string) ([]Process, error) {
	cmd := exec.Command("lsof", "-i", ":"+port)

	output, err := cmd.Output()

	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(output), "\n")

	if len(lines) <= 1 {
		return nil, fmt.Errorf("no process found")
	}

	processes := []Process{}

	for _, line := range lines[1:] {
		if strings.TrimSpace(line) == "" {
			continue
		}

		fields := strings.Fields(line)

		if len(fields) < 2 {
			continue
		}

		process := Process{
			Name:    fields[0],
			PID:     fields[1],
			Command: strings.Join(fields, " "),
			Port:    port,
		}

		processes = append(processes, process)
	}

	if len(processes) == 0 {
		return nil, fmt.Errorf("no valid processes found")
	}

	return processes, nil
}
