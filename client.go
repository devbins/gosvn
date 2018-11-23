package gosvn

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

type client struct {
	username string
	password string
	svnUrl   string
	svnDir   string
	Env      []string
}

// NewClient
func NewClient(username, password, url, workDir string) *client {

	path, err := filepath.Abs(workDir)
	if err != nil {
		panic(err)
	}
	fmt.Println(path)
	return &client{username: username, password: password, svnUrl: url, svnDir: path}
}

// Cleanup ...
func (this *client) Cleanup() error {
	_, err := this.run("cleanup")
	if err != nil {
		return err
	}
	return nil
}

// Update ...
func (this *client) Update() (string, error) {
	out, err := this.run("update")
	if err != nil {
		return "", err
	}
	return string(out), err

}

// Commit ...
func (this *client) Commit(msg string) (string, error) {
	out, err := this.run("commit", "-m", msg)
	if err != nil {
		return "", err
	}
	return string(out), nil
}

// Add ...
func (this *client) Add(file string) (string, error) {
	out, err := this.run("add", file)
	if err != nil {
		return "", err
	}
	return string(out), nil

}

// Status ...
func (this *client) Status() (*status, error) {
	out, err := this.run("status", "--xml")
	if err != nil {
		return nil, err
	}
	s := new(status)
	err = xml.Unmarshal(out, s)
	if err != nil {
		return nil, err
	}
	return s, nil

}



// Diff ...
func (this *client) Diff(start, end int) (string, error) {
	r := fmt.Sprintf("%d:%d", start, end)
	out, err := this.run("diff", "-r", r, this.svnUrl)
	if err != nil {
		return "", err
	}
	return string(out), nil
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
func (this *client) Log() (*SvnLog, error) {
	out, err := this.run("log", this.svnUrl, "--xml", "-v")
	if err != nil {
		return nil, err
	}
	l := new(SvnLog)
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
	cmd := []string{"checkout", this.svnUrl, this.svnDir}
	out, err := this.run(cmd...)
	if err != nil {
		return "", err
	}
	return string(out), nil
}

// client Checkout from specific revision
func (this *client) CheckoutWithRevision(revision string) (string, error) {
	cmd := []string{"checkout", this.svnUrl, this.svnDir}
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
	ops := []string{"--non-interactive", "--trust-server-cert"}

	if this.username != "" {
		args = append(args, "--username", this.username)
	}

	if this.password != "" {
		args = append(args, "--password", this.password)
	}

	args = append(args, ops...)

	log.Println("cmd args:", args)

	cmd := exec.Command("svn", args...)
	if len(this.Env) > 0 {
		cmd.Env = append(os.Environ(), this.Env...)
	}

	if this.svnDir != "" {
		if _, err := os.Stat(this.svnDir); err == nil {
			cmd.Dir = this.svnDir
			log.Printf("cmd dir %s \n", cmd.Dir)
		}
	}

	out, err := cmd.Output()
	if err != nil {
		log.Println("cmd err:", string(out), err)
		return nil, err
	}
	return out, nil
}
