//go:build linux
// +build linux

package cheat

import (
	"testing"
)

func TestProcMaps(t *testing.T) {
	err := WriteProcessMemory2(26397, 94725967954592, []byte("okokok"))
	t.Error(err)
}
