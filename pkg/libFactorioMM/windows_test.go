// +build windows

package libFactorioMM

import (
	"io/ioutil"
	"testing"
)

func TestSymlink(t *testing.T) {
	input := []byte("hello\ngo\n")
	errWrite := ioutil.WriteFile("/tmp/dat1", input, 0644)
	symlink("/tmp/dat1", "/tmp/dat2")
	out, errRead := ioutil.ReadFile("/tmp/dat2")
	if errWrite != nil {
		t.Error("Error writing to files test bad")
	}

	if errRead != nil {
		t.Error("Error reading from file test bad")
	}

	if string(out) != "hello\ngo\n" {
		t.Error("Symlink Didn't work correctly")
	}

}

func TestGetFactorioFolderName(t *testing.T) {
	if getFactorioFolderName() != "Factorio" {
		t.Error("OS Isn't returning correct name of factorio Configuration Folder")
	}
}
