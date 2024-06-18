package notes

import "errors"

var (
	ErrProjectNotFound       = errors.New("project not found")
	ErrDirectoryDoesNotExist = errors.New("directory does not exist")
)
