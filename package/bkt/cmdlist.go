package bkt

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func ListCmd(namespaces bool, longformat bool) error {
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

	var current string
	if namespaces {
		symlink, err := getCurrentNamespaceSymlink()
		if err != nil {
			return err
		}
		current, err = filepath.EvalSymlinks(symlink)
		if err != nil {
			return err
		}
	} else {
		current, err = os.Getwd()
		if err != nil {
			return err
		}
	}

	for _, entry := range entries {
		if longformat {
			var dir string
			if namespaces {
				dir = path.Join(listDir, entry.Name())
			} else {
				dir, err = Get(entry.Name())
				if err != nil {
					return err
				}
			}

			pre := "[ ]"
			if current == dir {
				pre = "[*]"
			}

			post := dir
			if namespaces {
				namespaceLabels, err := os.ReadDir(dir)
				if err != nil {
					return err
				}
				labelText := "labels"
				if len(namespaceLabels) == 1 {
					labelText = "label"
				}
				post = fmt.Sprintf("%d %s", len(namespaceLabels), labelText)
			}
			fmt.Println(fmt.Sprintf("%s %s -> %s", pre, entry.Name(), post))
		} else {
			fmt.Println(entry.Name())
		}
	}
	return nil
}
