package test

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
