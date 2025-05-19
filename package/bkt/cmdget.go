package bkt

import (
	"fmt"
)

func GetCmd(labelName string) error {
	resolvedPath, err := Get(labelName)
	if err != nil {
		return err
	}

	fmt.Println(resolvedPath)
	return nil
}
