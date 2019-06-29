package four

import "testing"

func TestIsPalindrome(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"11 is a palindrome", args{11}, true},
		{"12 is not a palindrome", args{12}, false},
		{"404 is a palindrome", args{404}, true},
		{"1234321 is a palindrome", args{1234321}, true},
		{"1234567 is not a palindrome", args{1234567}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.args.n); got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
