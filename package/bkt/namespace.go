package bkt

import "path"

func resolveNamespacePathToUse(namespaceToUse string) (string, error) {
	if namespaceToUse != "" {
		namespacesDir, err := getNamespacesDir()
		if err != nil {
			return "", err
		}
		return path.Join(namespacesDir, namespaceToUse), nil
	} else {
		return getCurrentNamespaceSymlink()
	}
}
