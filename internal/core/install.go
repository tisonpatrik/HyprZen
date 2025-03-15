package core

import (
	"os/exec"
	"time"
)

func Install() {
	time.Sleep(5 * time.Second)
}

func IsGrubInstalled() bool {
	_, err := exec.LookPath("grub-install")
	return err == nil
}
