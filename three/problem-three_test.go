package three

import (
	"testing"
)

func Test_IsPrime(t *testing.T) {
	type args struct {
		candidate int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"4 is not prime", args{4}, false},
		{"7 is prime", args{7}, true},
		{"10 is not prime", args{10}, false},
		{"11 is prime", args{11}, true},
		{"29 is prime", args{29}, true},
		{"50 is not prime", args{50}, false},
		{"100 is not prime", args{100}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPrime(tt.args.candidate); got != tt.want {
				t.Errorf("isPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_LargestPrimeFactorOf(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"LPF of 13195 is 29", args{13195}, 29},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LargestPrimeFactorOf(tt.args.n); got != tt.want {
				t.Errorf("LargestPrimeFactorOf() = %v, want %v", got, tt.want)
			}
		})
	}
}
