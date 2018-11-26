package gosvn

import "testing"
import "fmt"

var local = NewLocalClient("/home/manjaro/Desktop/svntest", "test", "test")

func TestLocalClient_Add(t *testing.T) {
	out, err := local.Add("js/index.js")
	if err != nil {
		t.Error(err)
	}
	t.Log(out)
	fmt.Println(out)

}

func TestLocalClient_Commit(t *testing.T) {
	out, err := local.Commit("添加 js file")
	if err != nil {
		t.Error(err)
	}
	t.Log(out)
	fmt.Println(out)

}

func TestLocalClient_Update(t *testing.T) {
	out, err := local.Update()
	if err != nil {
		t.Error(err)
	}
	t.Log(out)
	fmt.Println(out)

}

func TestLocalClient_CleanUp(t *testing.T) {
	err := local.CleanUp()
	if err != nil {
		t.Error(err)
	}

}
