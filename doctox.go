package docc

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	MSOFFICE_PATH = "C:/Program Files (x86)/Microsoft Office"
)

var ErrNotSupportOldDoc = errors.New("not supporting .doc file")
var office_dirs = []string{"Office12", "Office14", "Office15", "Office16"}

func findConvPath() (string, error) {
	msop := filepath.Clean(MSOFFICE_PATH)
	dirs, err := os.ReadDir(msop)
	if err != nil {
		return "", err
	}

	verDir := ""
	for _, dir := range dirs {
		if !strings.Contains(dir.Name(), "Office1") {
			continue
		}
		for _, d := range office_dirs {
			if dir.Name() == d {
				verDir = dir.Name()
			}
		}
	}
	if verDir == "" {
		return "", ErrNotSupportOldDoc
	}

	msop = filepath.Join(msop, verDir)
	fds, err := os.ReadDir(msop)
	if err != nil {
		return "", err
	}
	for _, fd := range fds {
		if fd.IsDir() {
			continue
		}
		if strings.ToLower(fd.Name()) == "wordconv.exe" {
			return filepath.Join(msop, fd.Name()), nil
		}
	}

	return "", ErrNotSupportOldDoc
}

func docToX(fp string) (docxPath string, err error) {
	docxPath = ""
	err = nil

	if runtime.GOOS != "windows" {
		err = ErrNotSupportOldDoc
		return
	}

	convPath, err := findConvPath()
	if err != nil {
		return "", err
	}

	tmpDir, err := os.MkdirTemp("", "docc_*")
	if err != nil {
		return
	}
	tmp := filepath.Join(tmpDir, "tmp.docx")
	cmd := exec.Command(convPath, "-oice", "-nme", fp, tmp)
	err = cmd.Run()
	if err != nil {
		return
	}

	docxPath = tmp
	return
}
