package gosvn

import (
	"errors"
	"strings"
	"strconv"
)

type RemoteClient struct {
	*CommonClient
}

// NewRemoteClient ...
func NewRemoteClient(url, username, password string) *RemoteClient {
	return &RemoteClient{CommonClient: NewCommonClient(url, username, password, false)}
}

// CheckOut ...
func (remoteClient *RemoteClient) CheckOut(dir string) error {
	err := verifyDir(dir)
	if err != nil {
		return err
	}

	_, err = remoteClient.runCmd("checkout", remoteClient.URLOrPath, dir)
	if err != nil {
		return err
	}
	return nil
}

// CheckOutWithRevision ...
func (remoteClient *RemoteClient) CheckOutWithRevision(dir string, revision int) error {
	err := verifyDir(dir)
	if err != nil {
		return err
	}

	_, err = remoteClient.runCmd("checkout", remoteClient.URLOrPath, dir, "-r", strconv.Itoa(revision))
	if err != nil {
		return err
	}
	return nil
}

// verifyDir ...
func verifyDir(dir string) error {
	if !strings.HasPrefix(dir, "/") {
		err := errors.New("Checkout dir must start with /")
		return err
	}
	return nil
}
