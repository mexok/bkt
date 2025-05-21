package bkt

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func ListCmd(namespaces bool, longformat bool, namespaceToUse string) error {
	err := defaultSetup()
	if err != nil {
		return err
	}

	namespacePath, err := resolveNamespacePathToUse(namespaceToUse)
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
		listDir = namespacePath
	}

	entries, err := os.ReadDir(listDir)
	if err != nil {
		return err
	}

	var current string
	if namespaces {
		current, err = filepath.EvalSymlinks(namespacePath)
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
				dir, err = get(entry.Name(), namespacePath)
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
