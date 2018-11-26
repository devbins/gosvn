package gosvn

import (
	"testing"
)

var remoteClient = NewRemoteClient("svn://127.0.0.1", "test", "test")

func TestRemoteClient_CheckOut(t *testing.T) {
	err := remoteClient.CheckOut("/home/manjaro/Desktop/gosvn")
	if err != nil {
		t.Error(err)
	}

}

func TestRemoteClient_CheckOutWithRevision(t *testing.T) {

	err := remoteClient.CheckOutWithRevision("/home/manjaro/Desktop/gore", 2)
	if err != nil {
		t.Error(err)
	}

}
