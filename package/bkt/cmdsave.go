package bkt

import (
	"errors"
	"fmt"
	"os"
	"path"
)

func SaveCmd(labelName string, force bool) error {
	err := defaultSetup()
	if err != nil {
		return err
	}

	currentNamespaceSymlink, err := getCurrentNamespaceSymlink()
	if err != nil {
		return err
	}
	label := path.Join(currentNamespaceSymlink, labelName)

	stat, err := os.Lstat(label)
	if err == nil && stat.Mode()&os.ModeSymlink > 0 {
		if !force {
			return errors.New(fmt.Sprintf("Label '%s' already exists in this namespace, use --force flag to overwrite or switch namespace", labelName))
		}
		err = os.Remove(label)
		if err != nil {
			return err
		}
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	return os.Symlink(cwd, label)
}
