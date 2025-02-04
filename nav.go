package main

import (
	"os"

	"github.com/coaguila/eval-env/tools"
)

func initDirectory(src string, name string) error {
	err := os.MkdirAll(src+name, tools.OS_ALL_RWX)

	return err
}
