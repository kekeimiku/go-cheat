//go:build linux
// +build linux

package cheat

import (
	"log"
	"os"
	"testing"
)

func readNextBytes(file *os.File, number int) []byte {
	bytes := make([]byte, number)

	_, err := file.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}

	return bytes
}

func TestProcMaps(t *testing.T) {

	/* maps, _ := ProcMaps("/proc/1156/maps")
	for _, v := range maps {
		fmt.Println(v)
	} */

	/* err := WriteProcessMemory2(26397, 94725967954592, []byte("okokok"))
	t.Error(err) */

	/* 	fmt.Println(SearchProcessMemory(3252, []byte("hhhhhhton"))) */
	/* 	file, err := os.Open("/proc/" + strconv.Itoa(1156) + "/mem")
	   	if err != nil {
	   		panic(err)
	   	}

	   	file.Seek(94375102029824, 1)

	   	a, _ := io.ReadAll(file)

	   	bytes.Index(a, []byte("ooooooton"))
	   	t.Error(bytes.Index(a, []byte("ooooooton")))
	*/
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

func TestSearchProcessMemory(t *testing.T) {
	c, e := SearchProcessMemory(1156, 94375102029824, []byte("ooooooton"))
	t.Error(c, e)
}
