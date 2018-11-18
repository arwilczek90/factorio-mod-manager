// +build windows

package libFactorioMM

import (
	"io/ioutil"
	"testing"
)

func TestSymlink(t *testing.T) {
	file1 := "tmp1"
	file2 := "tmp2"
	input := []byte("hello\ngo\n")
	errWrite := ioutil.WriteFile(file1, input, 0644)
	symlink(file1, file2)
	out, errRead := ioutil.ReadFile(file2)
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
