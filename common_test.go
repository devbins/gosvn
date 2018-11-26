package gosvn

import (
	"testing"
	"fmt"
)

var commonclient = NewCommonClient("/home/manjaro/Desktop/svntest", "test", "test",false)


func TestCommonClient_Log(t *testing.T) {
	out,err := commonclient.Log()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(out)
}

func TestCommonClient_Status(t *testing.T) {
	out,err :=  commonclient.Status()
	if err != nil {
		t.Error(err)
	}

	t.Log(out)
	fmt.Println(out)


}

func TestCommonClient_Properties(t *testing.T) {
	out,err := commonclient.Properties()
	if err != nil {
		t.Error(err)
	}
	t.Log(out)
	fmt.Println(out)
}

func TestCommonClient_Cat(t *testing.T) {
	out,err := commonclient.Cat("index.js")
	if err != nil {
		t.Error(err)
	}
	t.Log(out)
	fmt.Println(out)

}

func TestCommonClient_List(t *testing.T) {
	out,err := commonclient.List()
	if err != nil {
		t.Error(err)
	}
	t.Log(out)
	fmt.Println(out)

}

// TestCommonClient_DiffSummary ...
func TestCommonClient_DiffSummary(t *testing.T)  {
	out,err := commonclient.DiffSummary(0, 1)
	if err != nil {
		t.Error(err)
	}
	t.Log(out)
	fmt.Println(out)
}


// TestCommonClient_Diff ...
func TestCommonClient_Diff(t *testing.T)  {
	out,err := commonclient.Diff(0, 1)
	if err != nil {
		t.Error(err)
	}
	t.Log(out)
	fmt.Println(out)

}
