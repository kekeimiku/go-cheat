package cheat

import (
	"testing"
)

func TestProcMaps(t *testing.T) {
	a, err := ProcMaps("/proc/9230/maps")
	if err != nil {
		t.Error(err)
	}

	for _, v := range a {
		t.Errorf("%x-%x %v %v %v %v %v", v.StartAddr, v.EndAddr, v.Perms, v.Offset, v.Dev, v.Inode, v.Pathname)
	}
}
