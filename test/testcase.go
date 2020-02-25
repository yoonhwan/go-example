package test

import (
	"bufio"
	"fmt"
	"os"
)

func Sum(a,b int) int{
	return a + b
}
func Sum2(a, b int) (result int) {
	result = a + b
	return result
}

func test(m map[int]int) {
	for i:=0; i < 100000; i++ {
		m[i]= i
	}
}


func check(e error) {
	if e != nil {
		panic(e)
	}
}
func makefile(buff []byte) {
	fo, err := os.Create("/tmp/dat")
	check(err)
	defer fo.Close()

	_, err = fo.Write(buff)
	check(err)
}

func deletefile() {
	err := os.Remove("/tmp/dat")
	check(err)
}
func file() {
	f, err := os.Open("/tmp/dat")
	check(err)
	defer f.Close()

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	_ = fmt.Sprintf("5 bytes: %s\n", string(b4))
}