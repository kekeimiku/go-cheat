//go:build linux
// +build linux

package cheat

import (
	"testing"
)

func TestProcMaps(t *testing.T) {

	/* m, err := ProcMaps("/proc/34689/maps")
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range m {
		if v.Pathname == "[heap]" {
			fmt.Println(v)
		}
	} */
	//ReadCHeap("/proc/34689/maps")

	/* a, e := ReadProcessMemory2("A_ANONMYOUS", 1912)

	if e != nil {
		fmt.Println(e)
	}
	for _, v := range a {

		fmt.Printf("%x-%x %v %v %v\n", v.StartAddr, v.EndAddr, v.Perms, v.Offset, v.Pathname)
	} */

	var app = App{pid: 4989}
	app.ReadProcessMemory("A_ANONMYOUS").ReadProcessMemory("C_HEAP").Print()

	/* maps, err := ProcMaps("/proc/34689/maps")
	if err != nil {
		fmt.Println(err)
	}
	ma := []*ProcMap{}
	for _, v := range maps {
		if v.Perms.Read == true && v.Perms.Write == true && len(v.Pathname) < 42 {
			ma = append(ma, v)
		}
	}

	for _, v := range ma {
		fmt.Println(v)
	} */

	/* maps, _ := ProcMaps("/proc/30604/maps")
	for _, v := range maps {
		fmt.Println(v.Pathname)
	} */

	/* err := WriteProcessMemory2(26397, 94725967954592, []byte("okokok"))
	t.Error(err) */

	/* 	fmt.Println(SearchProcessMemory(3252, []byte("hhhhhhton"))) */

	//buf := make([]byte, 135168)

	//io.ReadAtLeast(file, buf, 135168)

	/* file, err := os.Open("/proc/" + strconv.Itoa(21475) + "/mem")
	if err != nil {
		panic(err)
	}

	file.Seek(824633720832, 1)

	a, _ := io.ReadAll(file)

	t.Error(bytes.Index(a, []byte("Holberton")))

	t.Error(len(a))

	defer file.Close() */

	/* reader := bufio.NewReader(io.ReadSeeker) */
	/* t.Error(e)
	t.Error(bytes.Index(data, []byte("ooooooton"))) */

	/* file.Seek(94375102030496, 0)
	_, e := file.Write([]byte("oooooo"))
	t.Error(e) */
	//55d56a6ea000-55d56a70b000

	/* var st string = "55d56a6ea000"
	b := fmt.Sprintf("%p", &st)
	stc, _ := strconv.ParseInt(b, 0, 64)

	var ed string = "55d56a70b000"
	eda := fmt.Sprintf("%p", &ed)
	edc, _ := strconv.ParseInt(eda, 0, 64) */

	/* 	a, er := file.Seek(int64(94375102029824), 0)
	   	t.Error(a, er)

	   	data := make([]byte, 9)

	   	file.ReadAt(data, int64(135168)) */
	/* t.Error(edc)
	t.Error(stc)
	t.Error(edc - stc) */
	/* 	t.Error(data)

	   	t.Error(bytes.Index(data, []byte("Holberton")))
	*/
	//94375102029824
	//data := make([]byte, 9)

	/* var ed string = "55d56a70b000"
	eda := fmt.Sprintf("%p", &ed)
	edc, _ := strconv.ParseInt(eda, 0, 64)
	*/
	//t.Error()
	/* 	data := make([]byte, 9)

	   	file.Read(data)
	   	t.Error(data)
	   	n, e := file.ReadAt(data, 135168)
	   	t.Error(n, e)
	   	t.Error(bytes.Index(data, []byte("Holberton")))
	*/
	//defer file.Close()

	//t.Error(n, err1)
}

/* func TestSearchProcessMemory(t *testing.T) {
	c, e := SearchProcessMemory(21745, 94375102029824, []byte("ooooooton"))
	t.Error(c, e)
}
*/
