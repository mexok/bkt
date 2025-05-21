package bkt

import (
	"fmt"
)

func GetCmd(labelName string, namespaceToUse string) error {
	namespacePath, err := resolveNamespacePathToUse(namespaceToUse)
	if err != nil {
		return err
	}

	resolvedPath, err := get(labelName, namespacePath)
	if err != nil {
		return err
	}

	fmt.Println(resolvedPath)
	return nil
}
