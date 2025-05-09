package core

import (
	"hyprzen/internal/core/bootloaders"
	"hyprzen/internal/utils"
)

func PreInstall() {
	nvidiaDetected := utils.IsNvidiaDetected()
	bootloaders.ConfigureGrub(nvidiaDetected)
	ConfigurePacman()
}
