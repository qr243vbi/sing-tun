//go:build linux && with_internal_resolvectl

package tun

import (
	"fmt"
	"os"
)

func getResolveCtl() (string, error) {
	path, err := os.Executable()
	fmt.Println("Path is:", path)
	return path, err
}
