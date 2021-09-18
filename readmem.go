package cheat

import (
	"fmt"
	"strconv"
)

var m []*ProcMap

func (a *App) Clear() *App {
	m = nil
	return a
}

func (a *App) Print() *App {
	for _, v := range m {
		fmt.Printf("%x-%x %v %v %v\n", v.StartAddr, v.EndAddr, v.Perms, v.Offset, v.Pathname)
	}
	return a
}

func (a *App) ReadProcessMemory(t string) *App {

	path := "/proc/" + strconv.Itoa(a.pid) + "/maps"

	switch {
	case t == "ALL":
		ProcMaps(path)
	case t == "B_BAD":
		ReadBad(path)
	case t == "V":
		ReadV(path)
	case t == "C_ALLOC":
		ReadCAlloc(path)
	case t == "C_BSS":
		ReadCBss(path)
	case t == "C_DATA":
		ReadCData(path)
	case t == "C_HEAP":
		ReadCHeap(path)
	case t == "JAVA_HEAP":
		ReadJavaHeap(path)
	case t == "A_ANONMYOUS":
		ReadCAnonmyous(path)
	case t == "CODE_SYSTEM":
		ReadCodeSystem(path)
	case t == "STACK":
		ReadStack(path)
	case t == "ASHMEM":
		ReadAshmem(path)
	case t == "CODE_APP":
		ReadCodeApp(path)
	default:
		fmt.Println("no type selected or invalid type")
	}
	return a
}

func ReadBad(path string) ([]*ProcMap, error) {
	maps, err := ProcMaps(path)
	if err != nil {
		return nil, err
	}

	for _, v := range maps {
		if v.Perms.Read && v.Perms.Write && v.Pathname == "kgsl-3d0" {
			m = append(m, v)
		}
	}
	return m, nil
}

func ReadV(path string) ([]*ProcMap, error) {

	maps, err := ProcMaps(path)
	if err != nil {
		return nil, err
	}
	////m := []*ProcMap{}
	for _, v := range maps {
		if v.Perms.Read && v.Perms.Write && v.Pathname == "/dev/kgsl-3d0" {
			m = append(m, v)
		}
	}
	return m, nil
}

func ReadCAlloc(path string) ([]*ProcMap, error) {
	maps, err := ProcMaps(path)
	if err != nil {
		return nil, err
	}
	////m := []*ProcMap{}
	for _, v := range maps {
		if v.Perms.Read && v.Perms.Write && v.Pathname == "[anon:libc_malloc]" {
			m = append(m, v)
		}
	}
	return m, nil
}

func ReadCBss(path string) ([]*ProcMap, error) {
	maps, err := ProcMaps(path)
	if err != nil {
		return nil, err
	}
	//m := []*ProcMap{}
	for _, v := range maps {
		if v.Perms.Read && v.Perms.Write && v.Pathname == "[anon:.bss]" {
			m = append(m, v)
		}
	}
	return m, nil
}

func ReadCData(path string) ([]*ProcMap, error) {
	maps, err := ProcMaps(path)
	if err != nil {
		return nil, err
	}
	//m := []*ProcMap{}
	for _, v := range maps {
		if v.Perms.Read && v.Perms.Write && v.Pathname == "/data/app/" {
			m = append(m, v)
		}
	}
	return m, nil
}

func ReadCHeap(path string) ([]*ProcMap, error) {
	maps, err := ProcMaps(path)
	if err != nil {
		return nil, err
	}
	//m := []*ProcMap{}
	for _, v := range maps {
		if v.Perms.Read && v.Perms.Write && v.Pathname == "[heap]" {
			m = append(m, v)
		}
	}
	return m, nil
}

func ReadJavaHeap(path string) ([]*ProcMap, error) {
	maps, err := ProcMaps(path)
	if err != nil {
		return nil, err
	}
	//m := []*ProcMap{}
	for _, v := range maps {
		if v.Perms.Read && v.Perms.Write && v.Pathname == "/dev/ashmem/" {
			m = append(m, v)
		}
	}
	return m, nil
}

func ReadCAnonmyous(path string) ([]*ProcMap, error) {
	maps, err := ProcMaps(path)
	if err != nil {
		return nil, err
	}

	//m := []*ProcMap{}
	for _, v := range maps {
		if v.Perms.Read && v.Perms.Write && len(v.Pathname) < 42 {
			m = append(m, v)
		}
	}
	return m, nil
}

func ReadCodeSystem(path string) ([]*ProcMap, error) {
	maps, err := ProcMaps(path)
	if err != nil {
		return nil, err
	}
	//m := []*ProcMap{}
	for _, v := range maps {
		if v.Perms.Read && v.Perms.Write && v.Pathname == "/system" {
			m = append(m, v)
		}
	}
	return m, nil
}

func ReadStack(path string) ([]*ProcMap, error) {
	maps, err := ProcMaps(path)
	if err != nil {
		return nil, err
	}
	//m := []*ProcMap{}

	for _, v := range maps {
		if v.Perms.Read && v.Perms.Write && v.Pathname == "[stack]" {
			m = append(m, v)
		}
	}
	return m, nil
}

func ReadOther(path string) ([]*ProcMap, error) {
	maps, err := ProcMaps(path)
	if err != nil {
		return nil, err
	}
	//m := []*ProcMap{}
	for _, v := range maps {
		if v.Perms.Read && v.Perms.Write && v.Pathname == "[anon:thread signal stack]" {
			m = append(m, v)
		}
	}
	return m, nil
}

func ReadAshmem(path string) ([]*ProcMap, error) {
	maps, err := ProcMaps(path)
	if err != nil {
		return nil, err
	}
	//m := []*ProcMap{}
	for _, v := range maps {

		if v.Perms.Read && v.Perms.Write && v.Pathname == "/dev/ashmem/" || v.Pathname == "dalvik" {
			m = append(m, v)
		}
	}
	return m, nil
}

func ReadCodeApp(path string) ([]*ProcMap, error) {
	maps, err := ProcMaps(path)
	if err != nil {
		return nil, err
	}

	//m := []*ProcMap{}
	for _, v := range maps {

		if v.Perms.Read && v.Perms.Execute && v.Perms.Private && v.Pathname == "/data/app/" {
			m = append(m, v)
		}
	}
	return m, nil
}
