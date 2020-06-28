package main

import (
	"os/exec"
)

func main() {

}

func getCmdUserPosition() (string, error) {
	out, err := exec.Command("cmd", "/C", "echo", "%CD%").CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
