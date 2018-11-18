// +build windows

package libFactorioMM

import "testing"

func TestSymlink(t *testing.T) {
}

func TestGetFactorioFolderName(t *testing.T) {
	if getFactorioFolderName() != "Factorio" {
		t.Error("OS Isn't returning correct name of factorio Configuration Folder")
	}
}
