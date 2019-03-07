package goclear

import (
	"os"
	"os/exec"
	"runtime"
)

func Clear() error {
	var command *exec.Cmd
	if operatingSystem := runtime.GOOS; operatingSystem == "linux" || operatingSystem == "darwin" {
		command = exec.Command("clear")
	} else if operatingSystem == "windows" {
		command = exec.Command("cmd", "/c", "cls")
	} else {
		panic("Unsupported OS: " + runtime.GOOS)
	}
	command.Stdout = os.Stdout
	return command.Run()
}
