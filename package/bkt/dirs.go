package bkt

import (
	"os"
	"path"
)

const DEFAULT_NAMESPACE = "default"

func getBaseDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return path.Join(homeDir, ".local", "share", "bkt"), nil
}

func getNamespacesDir() (string, error) {
	baseDir, err := getBaseDir()
	if err != nil {
		return "", err
	}
	return path.Join(baseDir, "namespaces"), nil
}

func getCurrentNamespaceSymlink() (string, error) {
	baseDir, err := getBaseDir()
	if err != nil {
		return "", err
	}
	return path.Join(baseDir, "current"), nil
}
