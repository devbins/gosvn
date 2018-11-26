package gosvn

type LocalClient struct {
	*CommonClient
}

// NewLocalClient ...
func NewLocalClient(dir, username, password string) *LocalClient {
	return &LocalClient{CommonClient: NewCommonClient(dir, username, password, false)}
}

// add ...
func (localClient *LocalClient) Add(file string) (string, error) {
	out, err := localClient.runCmd("add", file)
	if err != nil {
		return "", err
	}

	return string(out), nil

}

// Commit ...
func (localClient *LocalClient) Commit(msg string) (string, error) {
	out, err := localClient.runCmd("commit", "-m", msg)
	if err != nil {
		return "", err
	}

	return string(out), nil

}

// Update ...
func (localClient *LocalClient) Update() (string, error) {
	out, err := localClient.runCmd("update")
	if err != nil {
		return "", err
	}
	return string(out), nil
}

// CleanUp ...
func (localClient *LocalClient) CleanUp() error {
	_, err := localClient.runCmd("cleanup")
	if err != nil {
		return err
	}
	return nil
}
