package types

import "testing"

func TestReverseString(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test 1",
			args: args{
				text: "Hello",
			},
			want: "olleH",
		},
		{
			name: "test 2",
			args: args{
				text: "Hello,中国#& 1",
			},
			want: "1 &#国中,olleH",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseString(tt.args.text); got != tt.want {
				t.Errorf("ReverseString() = %v, want %v", got, tt.want)
			}
		})
	}
}
