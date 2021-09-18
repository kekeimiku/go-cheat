package cheat

import (
	"errors"
	"strconv"
)

func ReadProcessMemory2(t string, pid int) ([]*ProcMap, error) {

	path := "/proc/" + strconv.Itoa(pid) + "/maps"

	switch {
	case t == "ALL":
		return ProcMaps(path)
	case t == "B_BAD":
		return ReadBad(path)
	case t == "V":
		return ReadV(path)
	case t == "C_ALLOC":
		return ReadCAlloc(path)
	case t == "C_BSS":
		return ReadCBss(path)
	case t == "C_DATA":
		return ReadCData(path)
	case t == "C_HEAP":
		return ReadCHeap(path)
	case t == "JAVA_HEAP":
		return ReadJavaHeap(path)
	case t == "A_ANONMYOUS":
		return ReadCAnonmyous(path)
	case t == "CODE_SYSTEM":
		return ReadCodeSystem(path)
	case t == "STACK":
		return ReadStack(path)
	case t == "ASHMEM":
		return ReadAshmem(path)
	case t == "CODE_APP":
		return ReadCodeApp(path)
	default:
		return nil, errors.New("no type selected or invalid type")

	}
}

func ReadBad(path string) ([]*ProcMap, error) {
	maps, err := ProcMaps(path)
	if err != nil {
		return nil, err
	}
	m := []*ProcMap{}

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
	m := []*ProcMap{}
	for _, v := range maps {
		if v.Perms.Read == true && v.Perms.Write == true && v.Pathname == "/dev/kgsl-3d0" {
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
	m := []*ProcMap{}
	for _, v := range maps {
		if v.Perms.Read == true && v.Perms.Write == true && v.Pathname == "[anon:libc_malloc]" {
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
	m := []*ProcMap{}
	for _, v := range maps {
		if v.Perms.Read == true && v.Perms.Write == true && v.Pathname == "[anon:.bss]" {
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
	m := []*ProcMap{}
	for _, v := range maps {
		if v.Perms.Read == true && v.Perms.Write == true && v.Pathname == "/data/app/" {
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
	m := []*ProcMap{}
	for _, v := range maps {
		if v.Perms.Read == true && v.Perms.Write == true && v.Pathname == "[heap]" {
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
	m := []*ProcMap{}
	for _, v := range maps {
		if v.Perms.Read == true && v.Perms.Write == true && v.Pathname == "/dev/ashmem/" {
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

	m := []*ProcMap{}
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
	m := []*ProcMap{}
	for _, v := range maps {
		if v.Perms.Read == true && v.Perms.Write == true && v.Pathname == "/system" {
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
	m := []*ProcMap{}

	for _, v := range maps {
		if v.Perms.Read == true && v.Perms.Write == true && v.Pathname == "[stack]" {
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
	m := []*ProcMap{}
	for _, v := range maps {
		if v.Perms.Read == true && v.Perms.Write == true && v.Pathname == "[anon:thread signal stack]" {
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
	m := []*ProcMap{}
	for _, v := range maps {

		if v.Perms.Read == true && v.Perms.Write == true && v.Pathname == "/dev/ashmem/" || v.Pathname == "dalvik" {
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

	m := []*ProcMap{}
	for _, v := range maps {

		if v.Perms.Read == true && v.Perms.Execute == true && v.Perms.Private == true && v.Pathname == "/data/app/" {
			m = append(m, v)
		}
	}
	return m, nil
}
