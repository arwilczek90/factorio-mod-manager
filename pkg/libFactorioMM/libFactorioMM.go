package libFactorioMM

import (
	"fmt"
	"github.com/kirsle/configdir"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"runtime"
)

func linkModsFolder(profileName string, modsFolderPath string) {

}

func getDefaultConfigPath() string {
	configPath := configdir.LocalConfig("FactorioModManager")
	err := configdir.MakePath(configPath) // Ensure it exists.
	if err != nil {
		panic(err)
	}
	return configPath
}

func loadConfig(vipr *viper.Viper, configPath string) {
	if len(configPath) > 0 {
		vipr.AddConfigPath(configPath)
	} else {
		confFile := filepath.Join(getDefaultConfigPath(), "settings.json")
		vipr.AddConfigPath(confFile)
	}
}

func getDefaultModsFolderLocation() string {
	dirPath := ""
	if runtime.GOOS == "linux" {
		dirPath = filepath.Join(os.Getenv("HOME"), fmt.Sprintf(".%s", getFactorioFolderName()))

	} else {
		dirPath = configdir.LocalConfig(getFactorioFolderName())

	}
	modsFolderPath := filepath.Join(dirPath, "mods")
	return modsFolderPath
}

func getDefaultSavesFolderLocation() string {
	dirPath := ""
	if runtime.GOOS == "linux" {
		dirPath = filepath.Join(os.Getenv("HOME"), fmt.Sprintf(".%s", getFactorioFolderName()))

	} else {
		dirPath = configdir.LocalConfig(getFactorioFolderName())

	}
	modsFolderPath := filepath.Join(dirPath, "saves")
	return modsFolderPath
}
