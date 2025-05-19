package bkt

import (
	"errors"
	"fmt"
	"os"
)

func ListCmd(namespaces bool, longformat bool) error {
	if namespaces && longformat {
		return errors.New("Long format is currently not supported for namespaces")
	}

	err := defaultSetup()
	if err != nil {
		return err
	}

	var listDir string
	if namespaces {
		listDir, err = getNamespacesDir()
		if err != nil {
			return err
		}
	} else {
		listDir, err = getCurrentNamespaceSymlink()
		if err != nil {
			return err
		}
	}

	entries, err := os.ReadDir(listDir)
	if err != nil {
		return err
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if longformat {
			dir, err := Get(entry.Name())
			if err != nil {
				dir = err.Error()
			}
			pre := "[ ]"
			if cwd == dir {
				pre = "[*]"
			}
			fmt.Println(fmt.Sprintf("%s %s -> %s", pre, entry.Name(), dir))
		} else {
			fmt.Println(entry.Name())
		}
	}
	return nil
}
