package test

import "testing"

func TestSum(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test1", args{1,2}, 3},
		{"test2", args{4,10}, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSum2(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
	}{
		// TODO: Add test cases.
		{"test1", args{1,2}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Sum2(tt.args.a, tt.args.b); gotResult != tt.wantResult {
				t.Errorf("Sum2() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}


func BenchmarkSub1 (b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(1,2)
	}
}

func BenchmarkSum2 (b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum2(1, 2)
	}
}

func BenchmarkMap (b *testing.B) {
	for i:= 0; i< b.N; i++ {
		b.StopTimer()
		m := make(map[int]int)
		b.StartTimer()
		test(m)
	}
}

func BenchmarkCapacityMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		m := make(map[int]int, 100000)
		b.StartTimer()
		test(m)
	}
}