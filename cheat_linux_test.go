package cheat

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestFindPidByName(t *testing.T) {
	if FindPidByName("kthreadd") != 2 {
		t.Error("TestFindPidByName fail")
	}
}

func TestWriteProcessMemory(t *testing.T) {
	var a uint64 = 1
	b := fmt.Sprintf("%p", &a)
	c, _ := strconv.ParseUint(b, 0, 64)
	err := WriteProcessMemory(os.Getpid(), c, []byte{9})
	if a != 9 {
		t.Error(a, b, c, err)
	}
}

func TestReadProcessMemory(t *testing.T) {
	var a uint64 = 1
	b := fmt.Sprintf("%p", &a)
	c, _ := strconv.ParseUint(b, 0, 64)
	d, err := ReadProcessMemory(os.Getpid(), c, 1)

	if d != &[]byte{1} && err != nil {
		t.Error(d, err)
	}
}
