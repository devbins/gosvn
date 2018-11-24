package gosvn

type remoteClient struct {
	CommonClient
}

// CheckOut ...
func (remoteClient *remoteClient) CheckOut(dir string) error {
	_, err := remoteClient.RunCmd("checkout", dir, "--xml", remoteClient.URLOrPath)
	if err != nil {
		return err
	}
	return nil
}

// CheckOutWithRevision ...
func (remoteClient *remoteClient) CheckOutWithRevision(dir string, revision int) error {
	_, err := remoteClient.RunCmd("checkout", dir, "-r", string(revision), "--xml", remoteClient.URLOrPath)
	if err != nil {
		return err
	}
	return nil
}
