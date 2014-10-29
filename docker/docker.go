package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func boot2docker(arg string) string {
	cmd := exec.Command("boot2docker", arg)
	b, err := cmd.Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return strings.TrimSpace(string(b))
}

func docker(args []string) {
	f, err := os.Open("Dockerfile")
	if err == nil {
		cmd := exec.Command("boot2docker", "ssh", "tee", "Dockerfile")
		cmd.Stdout = os.Stdout
		cmd.Stdin = f
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		f.Close()
	}

	cmd := exec.Command("boot2docker", append([]string{"ssh", "-t", "docker"}, args...)...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func main() {
	if boot2docker("status") != "running" {
		boot2docker("up")
	}
	docker(os.Args[1:])
}
