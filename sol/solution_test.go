package sol

import "testing"

func BenchmarkTest(b *testing.B) {
	m, n := 5, 6
	for idx := 0; idx < b.N; idx++ {
		uniquePaths(m, n)
	}
}
func Test_uniquePaths(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "m = 3, n = 7",
			args: args{m: 3, n: 7},
			want: 28,
		},
		{
			name: "m = 3, n = 2",
			args: args{m: 3, n: 2},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePaths(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
