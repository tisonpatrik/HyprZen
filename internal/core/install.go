package core

import "hyprzen/internal/utils"

func PreInstall() {
	nvidiaDetected := utils.IsNvidiaDetected()
	ConfigureGrub(nvidiaDetected)
	ConfigurePacman()
}
