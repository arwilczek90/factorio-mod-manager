package libFactorioMM

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestGetDefaultModsFolderLocation(t *testing.T) {
	modsFolder := getDefaultModsFolderLocation()
	if runtime.GOOS == "windows" {
		if modsFolder != filepath.Join(os.Getenv("APPDATA"), "Factorio", "mods") {
			t.Error(fmt.Sprintf("Not Getting correct mods folder location for factorio config dir: %s", modsFolder))
		}
	} else if runtime.GOOS == filepath.Join(os.Getenv("HOME"), "Library", "Application Support", "factorio", "mods") {
		if modsFolder != "C:\\Users\\start" {
			t.Error(fmt.Sprintf("Not Getting correct mods folder location for factorio config dir: %s", modsFolder))
		}
	} else if runtime.GOOS == "linux" {
		if modsFolder != filepath.Join(os.Getenv("HOME"), ".factorio", "mods") {
			t.Error(fmt.Sprintf("Not Getting correct mods folder location for factorio config dir: %s", modsFolder))
		}
	}

}

func TestGetDefaultSavesFolderLocation(t *testing.T) {
	savesFolder := getDefaultSavesFolderLocation()
	if runtime.GOOS == "windows" {
		if savesFolder != filepath.Join(os.Getenv("APPDATA"), "Factorio", "saves") {
			t.Error(fmt.Sprintf("Not Getting correct mods folder location for factorio config dir: %s", savesFolder))
		}
	} else if runtime.GOOS == filepath.Join(os.Getenv("HOME"), "Library", "Application Support", "factorio", "saves") {
		if savesFolder != "C:\\Users\\start" {
			t.Error(fmt.Sprintf("Not Getting correct mods folder location for factorio config dir: %s", savesFolder))
		}
	} else if runtime.GOOS == "linux" {
		if savesFolder != filepath.Join(os.Getenv("HOME"), ".factorio", "saves") {
			t.Error(fmt.Sprintf("Not Getting correct mods folder location for factorio config dir: %s", savesFolder))
		}
	}

}
