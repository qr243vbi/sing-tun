//go:build linux && !with_internal_resolvectl

package tun

import "os/exec"

func getResolveCtl() (string, error) {
	return exec.LookPath("resolvectl")
}
