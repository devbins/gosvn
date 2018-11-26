package gosvn

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type CommonClient struct {
	URLOrPath string
	Username  string
	Password  string
	TrustCert bool
	Env       []string
}

// NewCommonClient ...
func NewCommonClient(urlOrPath, username, password string, trustCert bool) *CommonClient {
	return &CommonClient{URLOrPath: urlOrPath, Username: username, Password: password, TrustCert: trustCert}
}

// runCmd ...
func (client *CommonClient) runCmd(args ...string) ([]byte, error) {
	args = append(args, "--non-interactive")

	if client.TrustCert {
		args = append(args, "--trust-server-cert")
	}

	if client.Username != "" && client.Password != "" {
		args = append(args, "--username", client.Username)
		args = append(args, "--password", client.Password)
	}

	cmd := exec.Command("svn", args...)
	if len(client.Env) > 0 {
		cmd.Env = append(os.Environ(), client.Env...)
	}

	if strings.HasPrefix(client.URLOrPath, "/") {
		if _, err := os.Stat(client.URLOrPath); err == nil {
			cmd.Dir = client.URLOrPath
		}
	}

	output, err := cmd.Output()

	if err != nil {
		return nil, err
	}

	return output, nil

}

// info ...
func (client *CommonClient) Info() (*Info, error) {
	output, err := client.runCmd("info", "--xml", client.URLOrPath)
	if err != nil {
		return nil, err
	}

	info := new(Info)
	err = xml.Unmarshal(output, info)

	if err != nil {
		return nil, err
	}
	return info, nil
}

// Log ...
func (client *CommonClient) Log() (*SvnLog, error) {
	output, err := client.runCmd("log", client.URLOrPath, "--xml", "-v")
	if err != nil {
		return nil, err
	}
	log := new(SvnLog)

	err = xml.Unmarshal(output, log)
	if err != nil {
		return nil, err
	}

	return log, nil

}

// Status ...
func (client *CommonClient) Status() (*status, error) {
	output, err := client.runCmd("status", "--xml")
	if err != nil {
		return nil, err
	}
	status := new(status)
	err = xml.Unmarshal(output, status)
	if err != nil {
		return nil, err
	}
	return status, nil

}

// Properties ...
func (client *CommonClient) Properties() ([]byte, error) {
	out, err := client.runCmd("proplist", "--xml", client.URLOrPath)
	if err != nil {
		return nil, err
	}
	return out, nil

}

// Cat ...
func (client *CommonClient) Cat(file string) (string, error) {
	filePath := filepath.Join(client.URLOrPath, file)
	out, err := client.runCmd("cat", filePath)
	if err != nil {
		return "", err
	}
	return string(out), nil
}

// List ...
func (client *CommonClient) List() (*lists, error) {
	out, err := client.runCmd("ls", "--xml", client.URLOrPath)
	if err != nil {
		log.Println("error:", err)
		return nil, err
	}
	ls := new(lists)
	err = xml.Unmarshal(out, ls)
	if err != nil {
		log.Println("Unmarshal error:", err)
		return nil, err
	}
	return ls, nil
}

// Diff ...
func (client *CommonClient) Diff(start, end int) (string, error) {
	r := fmt.Sprintf("%d:%d", start, end)
	out, err := client.runCmd("diff", "-r", r, client.URLOrPath)
	if err != nil {
		return "", err
	}
	return string(out), nil
}

// DiffSummary ...
func (client *CommonClient) DiffSummary(start, end int) (*diffPath, error) {
	r := fmt.Sprintf("%d:%d", start, end)
	out, err := client.runCmd("diff", "-r", r, client.URLOrPath, "--xml", "--summarize")
	if err != nil {
		return nil, err
	}
	dp := new(diffPath)
	err = xml.Unmarshal(out, dp)
	if err != nil {
		return nil, err
	}
	return dp, nil
}
