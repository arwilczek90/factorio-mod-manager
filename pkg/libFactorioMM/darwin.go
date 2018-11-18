// +build darwin

package libFactorioMM

func symlink(target string, destination string) {
	os.Symlink(target, destination)
}

func getFactorioFolderName() string {
	return "factorio"
}
