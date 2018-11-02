package gosvn

import (
	"os/exec"
)

type remoteClient struct {
	url      string
	username string
	password string
}

// CheckOut ...
func (this *remoteClient) CheckOut(dir string) error {
	args := []string{"checkout", this.url, dir, "--username", this.username, "--password", this.password}
	_, err := runcmd(args...)
	return err
}

// run ...
func runcmd(args ...string) ([]byte, error) {
	cmd := exec.Command("svn", args...)
	out, err := cmd.Output()
	return out, err
}
