package bkt

import (
	"errors"
	"fmt"
	"os"
	"path"
)

func UseCmd(namespace string, create bool) error {
	err := defaultSetup()
	if err != nil {
		return err
	}

	namespacesDir, err := getNamespacesDir()
	if err != nil {
		return err
	}

	namespacePath := path.Join(namespacesDir, namespace)
	stat, err := os.Stat(namespacePath)
	if err == nil && stat.IsDir() {
		if create {
			return errors.New(fmt.Sprintf("Namespace '%s' already exists", namespace))
		}
	} else if !create {
		return errors.New(fmt.Sprintf("Namespace '%s' does not exist, use -c flag to create new", namespace))
	} else {
		err = os.MkdirAll(namespacePath, FILE_DIR_PERMISSION)
		if err != nil {
			return err
		}
	}

	currentNamespaceSymlink, err := getCurrentNamespaceSymlink()
	if err != nil {
		return err
	}

	err = os.Remove(currentNamespaceSymlink)
	if err != nil {
		return err
	}

	return os.Symlink(namespacePath, currentNamespaceSymlink)
}
