package bkt

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func get(labelName string, namespacePath string) (string, error) {
	label := path.Join(namespacePath, labelName)
	stat, err := os.Lstat(label)
	if err != nil || stat.Mode()&os.ModeSymlink == 0 {
		return "", errors.New(fmt.Sprintf("Label '%s' doesn't exist in this namespace", labelName))
	}

	resolvedPath, err := filepath.EvalSymlinks(label)
	if err != nil {
		return "", err
	}

	return resolvedPath, nil
}
