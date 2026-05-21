package finder

import (
	"fmt"
	"os/exec"
	"strings"
)

type Process struct {
	PID  string
	Name string
	Port string
}

func FindByPort(port string) (*Process, error) {
	cmd := exec.Command("lsof", "-i", ":"+port)

	output, err := cmd.Output()

	if err != nil {
		return nil, err
	}

	// The output of lsof -i :<port> is typically in the format:
	// COMMAND   PID USER   FD   TYPE DEVICE SIZE/OFF NODE NAME
	// process1  123 user   10u  IPv4 0x12345678      0t0  TCP *:port (LISTEN)
	// We need to parse this output to extract the process name and PID.
	lines := strings.Split(string(output), "\n")

	if len(lines) < 2 {
		return nil, fmt.Errorf("no process found")
	}

	// The first line is the header, so we take the second line for the process information.
	fields := strings.Fields(lines[1])

	if len(fields) < 2 {
		return nil, fmt.Errorf("unable to parse process")
	}

	process := &Process{
		Name: fields[0],
		PID:  fields[1],
		Port: port,
	}

	return process, nil
}
