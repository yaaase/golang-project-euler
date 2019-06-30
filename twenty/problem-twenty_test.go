package twenty

import (
	"math/big"
	"reflect"
	"testing"
)

func Test_sumOfDigits(t *testing.T) {
	type args struct {
		x string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Sum of digits 10055 is 11", args{"10055"}, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfDigits(tt.args.x); got != tt.want {
				t.Errorf("sumOfDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFactorial(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want big.Int
	}{
		{"Factorial of 5 is 120", args{5}, *big.NewInt(120)},
		{"Factorial of 10 is 3628800", args{10}, *big.NewInt(3628800)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Factorial(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}
