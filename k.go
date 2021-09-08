package k

import (
	"syscall"
	"unsafe"
)

func Syscall6(trap, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err syscall.Errno)

type Iovec struct {
	Base *byte
	Len  uint64
}
type RemoteIovec struct {
	Base uintptr
	Len  int
}

var _zero uintptr

type Errno uintptr

var (
	errEAGAIN error = syscall.EAGAIN
	errEINVAL error = syscall.EINVAL
	errENOENT error = syscall.ENOENT
)

func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return nil
	case syscall.Errno(0xb):
		return errEAGAIN
	case syscall.Errno(0x16):
		return errEINVAL
	case syscall.Errno(0x2):
		return errENOENT
	}
	return e
}

func WriteProcessMemory(pid int, addr uint64, buffer []byte) error {

	if _, err := ProcessVMWritev(pid, []Iovec{{
		Base: &buffer[0],
		Len:  uint64(len(buffer)),
	}}, []RemoteIovec{{
		Base: uintptr(addr),
		Len:  int(len(buffer)),
	}}, 0); err != nil {
		return err
	}

	return nil
}

func ReadProcessMemory(pid int, addr uint64, size uint64) (*[]byte, error) {
	data := make([]byte, size)

	if _, err := ProcessVMReadv(pid, []Iovec{{
		Base: &data[0],
		Len:  size,
	}}, []RemoteIovec{{
		Base: uintptr(addr),
		Len:  int(size),
	}}, 0); err != nil {
		return nil, err
	}

	return &data, nil
}

func ProcessVMWritev(pid int, localIov []Iovec, remoteIov []RemoteIovec, flags uint) (n int, err error) {
	var _p0 unsafe.Pointer
	if len(localIov) > 0 {
		_p0 = unsafe.Pointer(&localIov[0])
	} else {
		_p0 = unsafe.Pointer(&_zero)
	}
	var _p1 unsafe.Pointer
	if len(remoteIov) > 0 {
		_p1 = unsafe.Pointer(&remoteIov[0])
	} else {
		_p1 = unsafe.Pointer(&_zero)
	}
	r0, _, e1 := Syscall6(311, uintptr(pid), uintptr(_p0), uintptr(len(localIov)), uintptr(_p1), uintptr(len(remoteIov)), uintptr(flags))
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func ProcessVMReadv(pid int, localIov []Iovec, remoteIov []RemoteIovec, flags uint) (n int, err error) {
	var _p0 unsafe.Pointer
	if len(localIov) > 0 {
		_p0 = unsafe.Pointer(&localIov[0])
	} else {
		_p0 = unsafe.Pointer(&_zero)
	}
	var _p1 unsafe.Pointer
	if len(remoteIov) > 0 {
		_p1 = unsafe.Pointer(&remoteIov[0])
	} else {
		_p1 = unsafe.Pointer(&_zero)
	}
	r0, _, e1 := Syscall6(310, uintptr(pid), uintptr(_p0), uintptr(len(localIov)), uintptr(_p1), uintptr(len(remoteIov)), uintptr(flags))
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
