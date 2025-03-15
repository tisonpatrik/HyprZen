package core

func PreInstall() {
	ConfigureGrub()
	ConfigureSystemdBoot()
	ConfigurePacman()
}
