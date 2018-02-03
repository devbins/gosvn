package gosvn

import (
	"encoding/xml"
	"os"
	"os/exec"
)

type client struct {
	username string
	password string
	svnUrl   string
	svnDir   string
	Env      []string
}

// NewClient
func NewClient(username, password, url string) *client {
	return &client{username: username, password: password, svnUrl: url}
}

// NewClientWithEnv ...
func NewClientWithEnv(username, password, url string, env []string) *client {
	return &client{username: username, password: password, svnUrl: url, Env: env}
}

// Cat ...
func (this *client) Cat(file string) (string, error) {
	out, err := this.run("cat", this.svnUrl+"/"+file)
	if err != nil {
		return "", err
	}
	return string(out), nil
}

// Export ...
func (this *client) Export(dir string) error {
	_, err := this.run("export", this.svnUrl, dir)
	if err != nil {
		return err
	}
	return nil

}

// Log ...
func (this *client) Log() (*log, error) {
	out, err := this.run("log", this.svnUrl, "--xml", "-v")
	if err != nil {
		return nil, err
	}
	l := new(log)
	err = xml.Unmarshal(out, l)
	if err != nil {
		return nil, err
	}
	return l, nil
}

// list ...
func (this *client) List() (*lists, error) {
	cmd := []string{"list", this.svnUrl, "--xml"}
	out, err := this.run(cmd...)
	if err != nil {
		return nil, err
	}
	l := new(lists)
	err = xml.Unmarshal(out, l)
	if err != nil {
		return nil, err
	}
	return l, nil

}

// Checkout
func (this *client) Checkout() (string, error) {
	cmd := []string{"checkout", this.svnUrl}
	if this.svnDir != "" {
		cmd = append(cmd, this.svnDir)
	}
	out, err := this.run(cmd...)
	if err != nil {
		return "", err
	}
	return string(out), nil
}

// client Checkout from specific revision
func (this *client) CheckoutWithRevision(revision string) (string, error) {
	cmd := []string{"checkout", this.svnUrl}
	if this.svnDir != "" {
		cmd = append(cmd, this.svnDir)
	}
	cmd = append(cmd, "-r", revision)
	out, err := this.run(cmd...)
	if err != nil {
		return "", err
	}
	return string(out), nil
}

// info ...
func (this *client) Info() (*Info, error) {
	cmd := []string{"info", this.svnUrl, "--xml"}
	out, err := this.run(cmd...)
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

// run 运行命令
func (this *client) run(args ...string) ([]byte, error) {
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
