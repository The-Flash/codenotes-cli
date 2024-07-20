package notes

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

const (
	RootDir  = ".codenotes"
	NotesDir = "notes"
)

// InitializeProject creates the required directories
// necessary for a codenotes project
func InitializeProject(directory string) (string, error) {
	return createProject(directory)
}

// AddNote adds notes to the repo
func AddNote(nc *noteConfig) (string, error) {
	fmt.Println(nc)
	return "", nil
}

// GetProject get the closest codenotes projects.
// can be used in subdirectories
func GetProject() (string, error) {
	d, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return getProject(d)
}

func getProject(directory string) (string, error) {
	projectDir := path.Join(directory, RootDir)
	if dirExists(projectDir) {
		return directory, nil
	}
	parent := filepath.Dir(directory)
	if isRootDirectory(parent) {
		return "", ErrProjectNotFound
	}
	return getProject(parent)
}

func createProject(directory string) (string, error) {
	if !dirExists(directory) {
		return "", ErrDirectoryDoesNotExist
	}
	codenotesPath := path.Join(directory, RootDir)
	err := os.MkdirAll(codenotesPath, 0755)
	if err != nil {
		return "", err
	}
	err = os.MkdirAll(path.Join(codenotesPath, NotesDir), 0755)
	if err != nil {
		return "", err
	}
	fullPath, err := filepath.Abs(codenotesPath)
	return fullPath, err
}

func isRootDirectory(p string) bool {
	cleanPath := filepath.Clean(p)
	if runtime.GOOS == "windows" {
		return filepath.VolumeName(cleanPath)+`\` == cleanPath
	}
	return cleanPath == "/"
}

func dirExists(p string) bool {
	info, err := os.Stat(p)
	if errors.Is(err, fs.ErrNotExist) {
		return false
	}
	return info.IsDir()
}
