package bkt

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func GetCmd(labelName string) error {
	currentNamespaceSymlink, err := getCurrentNamespaceSymlink()
	if err != nil {
		return err
	}

	label := path.Join(currentNamespaceSymlink, labelName)
	stat, err := os.Lstat(label)
	if err != nil || stat.Mode()&os.ModeSymlink == 0 {
		return errors.New(fmt.Sprintf("Label '%s' doesn't exist in this namespace", labelName))
	}

	resolvedPath, err := filepath.EvalSymlinks(label)
	if err != nil {
		return err
	}

	fmt.Println(resolvedPath)
	return nil
}
