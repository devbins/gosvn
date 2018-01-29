package svn

import (
	"encoding/xml"
	"os"
	"os/exec"
)

type Client struct {
	username string
	password string
	svnUrl   string
	svnDir   string
	Env      []string
}

// NewClient
func NewClient(username, password, url string) *Client {
	return &Client{username: username, password: password, svnUrl: url}
}

// NewClientWithEnv ...
func NewClientWithEnv(username, password, url string, env []string) *Client {
	return &Client{username: username, password: password, svnUrl: url, Env: env}
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
func (this *Client) checkout() (string, error) {
	cmd := []string{"checkout", this.svnUrl}
	if this.svnDir != "" {
		cmd = append(cmd, this.svnDir)
	}
	out, err := this.Run(cmd...)
	if err != nil {
		return "", err
	}
	return string(out), nil
}

// client checkout from specific revision
func (this *Client) checkoutWithRevision(revision string) (string, error) {
	cmd := []string{"checkout", this.svnUrl}
	if this.svnDir != "" {
		cmd = append(cmd, this.svnDir)
	}
	cmd = append(cmd, "-r", revision)
	out, err := this.Run(cmd...)
	if err != nil {
		return "", err
	}
	return string(out), nil
}

// info ...
func (this *Client) Info() (*Info, error) {
	cmd := []string{"info", this.svnUrl, "--xml"}
	out, err := this.Run(cmd...)
	if err != nil {
		return nil, err
	}
	info := new(Info)
	err = xml.Unmarshal(out, info)
	if err != nil {
		return nil, err
	}
	return info, nil
}

// Run ...
func (this *Client) Run(args ...string) ([]byte, error) {
	ops := []string{"--username", this.username, "--password", this.password}
	args = append(args, ops...)
	cmd := exec.Command("svn", args...)
	if len(this.Env) > 0 {
		cmd.Env = append(os.Environ(), this.Env...)
	}
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return out, nil
}
