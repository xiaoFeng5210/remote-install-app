package main

import (
	"os/exec"
)

func main() {
	apkPath := "app-release.apk"

	cmd := exec.Command("adb", "install", apkPath)
}
