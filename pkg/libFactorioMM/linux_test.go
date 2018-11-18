// +build linux

package libFactorioMM

import "testing"

func TestSymlink(t *testing.T) {
}

func TestGetFactorioFolderName(t *testing.T) {
	if getFactorioFolderName() != "factorio" {
		t.Error("OS Isn't returning correct name of factorio Configuration Folder")
	}
}
