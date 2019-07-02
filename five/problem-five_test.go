package five

import "testing"

func Test_isDivisibleByAll(t *testing.T) {
	type args struct {
		n int
		m int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"6 is evenly divisible from 1 to 3", args{6, 3}, true},
		{"2520 is evenly divisible from 1 to 10", args{2520, 10}, true},
		{"10 is not evenly divisible from 1 to 3", args{10, 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isDivisibleByAll(tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("isDivisibleByAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
