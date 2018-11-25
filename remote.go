package gosvn

type RemoteClient struct {
	*CommonClient
}

// NewRemoteClient ...
func NewRemoteClient(url, username, password string) *RemoteClient {
	return &RemoteClient{CommonClient: NewCommonClient(url, username, password)}
}

// CheckOut ...
func (remoteClient *RemoteClient) CheckOut(dir string) error {
	_, err := remoteClient.RunCmd("checkout", remoteClient.URLOrPath, dir)
	if err != nil {
		return err
	}
	return nil
}

// CheckOutWithRevision ...
func (remoteClient *RemoteClient) CheckOutWithRevision(dir string, revision int) error {
	_, err := remoteClient.RunCmd("checkout", dir, "-r", string(revision), remoteClient.URLOrPath)
	if err != nil {
		return err
	}
	return nil
}
