package gosvn

import (
	"fmt"
	"testing"
)

var c = NewClient("test", "test", "svn://10.211.55.52", "")

// TestInfo ...
func TestInfo(t *testing.T) {
	out, err := c.Info()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(out)
}

// TestCheckOut ...
func TestCheckOut(t *testing.T) {
	out, err := c.Checkout()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(out)
}

// TestList ...
func TestList(t *testing.T) {
	l, err := c.List()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(l)
}

// TestLog ...
func TestLog(t *testing.T) {
	out, err := c.Log()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(out)
}

// TestExport ...
func TestExport(t *testing.T) {
	err := c.Export("")
	if err != nil {
		t.Error(err)
	}
}

// TestCat ...
func TestCat(t *testing.T) {
	out, err := c.Cat("readme.md")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(out)
}

// TestDiff ...
func TestDiff(t *testing.T) {
	out, err := c.Diff(39, 40)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(out)
}

// TestDiffSummary ...
func TestDiffSummary(t *testing.T) {
	out, err := c.DiffSummary(0, 40)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(out)

}

// TestStatus ...
func TestStatus(t *testing.T) {
	out, err := c.Status()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(out)

}

// TestAdd ...
func TestAdd(t *testing.T) {
	out, err := c.Add("add.md")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(out)

}

// TestCommit ...
func TestCommit(t *testing.T) {
	out, err := c.Commit("test commit")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(out)

}

// TestCleanup ...
func TestCleanup(t *testing.T) {
	out, err := c.Cleanup()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(out)
}
