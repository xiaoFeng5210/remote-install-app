package core

import (
	"os/exec"
)

func CheckAdbInstalled() error {
	cmd := exec.Command("adb", "version")

}
