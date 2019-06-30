package sixteen

import "testing"

func TestSumOfDigits(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Sum of 1234 is 10", args{1234}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumOfDigits(tt.args.x); got != tt.want {
				t.Errorf("SumOfDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
