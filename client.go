package svn

import (
	"encoding/xml"
	"fmt"
	"os"
	"os/exec"
)

type Cmder interface {
	// Run ...
	Run(cmd string)
}

type client struct {
	Env      []string
	svnUrl   string
	svnDir   string
	username string
	password string
}

// Run ...
func (this *client) Run(args ...string) (string, error) {
	cmd := exec.Command("svn", args, "--username", this.username, "--password", this.password)
	cmd.Env = append(os.Environ(), client.Env...)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return out, nil
}

// log ...
func (this *client) log(cmd string) (logEntry, error) {
	out, err := this.Run(cmd)
	if err != nil {
		return nil, err
	}
	log := new(logEntry)
	err = xml.Unmarshal(out, log)
	if err != nil {
		return nil, err
	}

}

// checkout
func (this *client) checkout(cmd string) (string, error) {

}

// client checkout from specific revision
func (this *client) checkoutWithRevision(url string, revision int) (string, error) {

}
