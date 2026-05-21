package finder

import (
	"fmt"
	"os/exec"
)

func FindByPort(port string) error {
	cmd := exec.Command("lsof", "-i", ":"+port)

	output, err := cmd.Output()

	if err != nil {
		return err
	}

	fmt.Println(string(output))

	return nil
}
