package main

import (
	"fmt"
	"os"
	"os/exec"
)

func install(f string) error {
	command := fmt.Sprintf("sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf %s", f)
	cmd := exec.Command("/bin/sh", "-c", command)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
