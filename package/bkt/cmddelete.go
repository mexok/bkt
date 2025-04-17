package bkt

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func DeleteCmd(name string, namespace bool, yes bool) error {
	currentNamespaceSymlink, err := getCurrentNamespaceSymlink()
	if err != nil {
		return err
	}

	stat, err := os.Lstat(currentNamespaceSymlink)
	if err != nil || stat.Mode()&os.ModeSymlink == 0 {
		return errors.New("Current namespace doesn't exist")
	}

	resolvedNamespace, err := filepath.EvalSymlinks(currentNamespaceSymlink)
	if err != nil {
		return err
	}

	if namespace {
		if !yes {
			return errors.New("Please confirm deletion of namespace using -y")
		}
		err = os.RemoveAll(resolvedNamespace)
		if err != nil {
			return err
		}
		return os.Remove(currentNamespaceSymlink)
	} else {
		pathOfName := path.Join(resolvedNamespace, name)
		labelStat, err := os.Lstat(pathOfName)
		if err != nil || labelStat.Mode()&os.ModeSymlink == 0 {
			return errors.New(fmt.Sprintf("Label '%s' doesn't exist in current namespace. Cannot delete", name))
		}
		return os.Remove(pathOfName)
	}
}
