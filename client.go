package svn

import (
	"encoding/xml"
	"fmt"
	"os"
	"os/exec"
)

type Client struct {
	Username string
	Password string
	SvnUrl   string
	SvnDir   string
	Env      []string
}

// log ...
func (this *Client) log(cmd string) (*logEntry, error) {
	out, err := this.Run(cmd)
	if err != nil {
		return nil, err
	}
	log := new(logEntry)
	err = xml.Unmarshal([]byte(out), log)
	if err != nil {
		return nil, err
	}
	return log, nil
}

// checkout
func (this *Client) checkout(cmd string) (string, error) {
	fmt.Println("checkout")
	return "", nil
}

// client checkout from specific revision
func (this *Client) checkoutWithRevision(url string, revision int) (string, error) {
	fmt.Println("checkoutwithrevision")
	return "", nil
}

// info ...
func (this *Client) Info() (string, error) {
	cmd := "info"
	out, err := this.Run(cmd)
	if err != nil {
		return "", err
	}
	info := new(info)
	err = xml.Unmarshal([]byte(out), info)
	if err != nil {
		return "", err
	}
	return out, nil
}

// Run ...
func (this *Client) Run(args ...string) (string, error) {
	ops := []string{this.SvnUrl, "--xml", "--username", this.Username, "--password", this.Password}
	args = append(args, ops...)
	cmd := exec.Command("svn", args...)
	cmd.Env = append(os.Environ(), this.Env...)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("error")
		fmt.Println(err)
		return "", err
	}
	return string(out), nil
}
