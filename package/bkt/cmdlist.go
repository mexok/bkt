package bkt

import (
	"fmt"
	"os"
)

func ListCmd(namespaces bool) error {
	err := defaultSetup()
	if err != nil {
		return err
	}

	var listDir string
	if namespaces {
		listDir, err = getCurrentNamespaceSymlink()
		if err != nil {
			return err
		}
	} else {
		listDir, err = getNamespacesDir()
		if err != nil {
			return err
		}
	}

	entries, err := os.ReadDir(listDir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		fmt.Println(entry.Name())
	}
	return nil
}
