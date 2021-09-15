//go:build linux
// +build linux

package cheat

import (
	"os"
	"path/filepath"
	"strconv"
	"syscall"

	"golang.org/x/sys/unix"
)

//FindPidByName return pid based on process name.
//If the return value is -1, it is a failure.
func FindPidByName(name string) int {
	dir, _ := filepath.Glob("/proc/*")
	for _, f := range dir {
		r, _ := os.ReadFile(f + "/comm")
		if string(r) == name+"\n" {
			pid, _ := strconv.Atoi(string(f[6:]))
			return pid
		}
	}
	return -1
}

//FindPidByName find the pids of all processes with the same name and return an array.
//If the length of the returned array is 0, it is not found.
func FindPidByNameGlob(name string) []int {
	var p []int
	dir, _ := filepath.Glob("/proc/*")
	for _, f := range dir {
		r, _ := os.ReadFile(f + "/comm")
		if string(r) == name+"\n" {
			pid, _ := strconv.Atoi(string(f[6:]))
			p = append(p, pid)
		}
	}
	return p
}

//Need root privileges
//WriteProcessMemory provides the function of write data to the process.
func WriteProcessMemory(pid int, addr uint64, data []byte) error {
	if _, err := unix.ProcessVMWritev(pid, []unix.Iovec{{
		Base: &data[0],
		Len:  uint64(len(data)),
	}}, []unix.RemoteIovec{{
		Base: uintptr(addr),
		Len:  int(len(data)),
	}}, 0); err != nil {
		return err
	}
	return nil
}

//Need root privileges
//WriteProcessMemory provides the function of read data to the process.
func ReadProcessMemory(pid int, addr uint64, size uint64) (*[]byte, error) {
	data := make([]byte, size)
	if _, err := unix.ProcessVMReadv(pid, []unix.Iovec{{
		Base: &data[0],
		Len:  size,
	}}, []unix.RemoteIovec{{
		Base: uintptr(addr),
		Len:  int(size),
	}}, 0); err != nil {
		return nil, err
	}
	return &data, nil
}

//KillProcessByPid kill the process by pid.
func KillProcessByPid(pid int) error {
	err := syscall.Kill(pid, syscall.SIGKILL)
	if err != nil {
		return err
	}
	return err
}

func WriteProcessMemory2(pid int, offset int64, data []byte) error {

	file, err := os.OpenFile("/proc/"+strconv.Itoa(pid)+"/mem", os.O_RDWR, 0)

	if err != nil {
		return err
	}
	/* var a string = "56271babc2a0"
	   b := fmt.Sprintf("%p", &a)
	   adr, e := strconv.ParseInt(b, 0, 64)

	   fmt.Println(adr, e) */
	_, err = file.Seek(offset, 1)

	if err != nil {
		return err
	}
	//fmt.Println(ret, err)
	_, err = file.Write(data)
	if err != nil {
		return err
	}

	defer file.Close()

	return nil
}
