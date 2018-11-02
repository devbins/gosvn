package gosvn

import (
	"encoding/xml"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

type CommonClient struct {
	UrlOrPath string
	Username  string
	Password  string
	TrustCert bool
	Env       []string
}

// NewCommonClient ...
func NewCommonClient(urlOrPath, username, password string) *CommonClient {
	return &CommonClient{UrlOrPath: urlOrPath, Username: username, Password: password}
}

// RunCmd ...
func (client *CommonClient) RunCmd(args ...string) ([]byte, error) {
	args = append(args, "--non-interactive")
	if client.Username != "" && client.Password != "" {
		args = append(args, "--username", client.Username)
		args = append(args, "--password", client.Password)
	}

	log.Println(args)

	cmd := exec.Command("svn", args...)
	if len(client.Env) > 0 {
		cmd.Env = append(os.Environ(), client.Env...)
	}

	output, err := cmd.Output()

	if err != nil {
		return nil, err
	}

	return output, nil

}

// info ...
func (client *CommonClient) Info() (*Info, error) {
	output, err := client.RunCmd("info", "--xml", client.UrlOrPath)
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
	output, err := client.RunCmd("log", client.UrlOrPath, "--xml", "-v")
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
	output, err := client.RunCmd("status", "--xml")
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
	out, err := client.RunCmd("proplist", "--xml")
	if err != nil {
		return nil, err
	}
	return out, nil

}

// Cat ...
func (client *CommonClient) Cat(file string) (string, error) {
	filePath := filepath.Join(client.UrlOrPath, file)
	out, err := client.RunCmd("cat", filePath)
	if err != nil {
		return "", err
	}
	return string(out), nil
}

// List ...
func (client *CommonClient) List() (*lists, error) {
	out, err := client.RunCmd("ls", "--xml", client.UrlOrPath)
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
