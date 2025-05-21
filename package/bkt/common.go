package bkt

import (
	"os"
	"path"
)

const FILE_DIR_PERMISSION = 0755

func defaultSetup() error {
	namespacesDir, err := getNamespacesDir()
	if err != nil {
		return err
	}
	currentNamespaceSymlink, err := getCurrentNamespaceSymlink()
	if err != nil {
		return err
	}
	_, err = os.Stat(currentNamespaceSymlink)
	if os.IsNotExist(err) {
		defaultNamespace := path.Join(namespacesDir, DEFAULT_NAMESPACE)
		err = os.MkdirAll(defaultNamespace, FILE_DIR_PERMISSION)
		if err != nil {
			return err
		}
		err = os.Symlink(defaultNamespace, currentNamespaceSymlink)
	}
	return err
}
