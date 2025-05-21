package bkt

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func DeleteCmd(name string, namespace bool, yes bool, namespaceToUse string) error {
	currentNamespaceSymlink, err := getCurrentNamespaceSymlink()
	if err != nil {
		return err
	}

	stat, err := os.Lstat(currentNamespaceSymlink)
	if err != nil || stat.Mode()&os.ModeSymlink == 0 {
		return errors.New("Current namespace doesn't exist")
	}

	namespacePath, err := resolveNamespacePathToUse(namespaceToUse)
	if err != nil {
		return err
	}

	if namespace {
		if !yes {
			return errors.New("Please confirm deletion of namespace using -y")
		}

		resolvedNamespacePath, err := filepath.EvalSymlinks(namespacePath)
		if err != nil {
			return err
		}

		err = os.RemoveAll(resolvedNamespacePath)
		if err != nil {
			return err
		}

		_, err = filepath.EvalSymlinks(currentNamespaceSymlink)
		if err != nil {
			err = os.Remove(currentNamespaceSymlink)
			if err != nil {
				return err
			}
			return defaultSetup()
		}
		return nil
	} else {
		pathOfName := path.Join(namespacePath, name)
		labelStat, err := os.Lstat(pathOfName)
		if err != nil || labelStat.Mode()&os.ModeSymlink == 0 {
			return errors.New(fmt.Sprintf("Label '%s' doesn't exist in current namespace. Cannot delete", name))
		}
		return os.Remove(pathOfName)
	}
}
