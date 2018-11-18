package libFactorioMM

import (
	"github.com/kirsle/configdir"
	"github.com/spf13/viper"
	"path/filepath"
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

func getDefaultModsFolderLocation() {

}
