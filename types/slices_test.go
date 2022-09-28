package types

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicate(t *testing.T) {
	type args struct {
		sliceList []string
	}
	type testCase struct {
		name string
		args args
		want []string
	}
	tests := []testCase{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{sliceList: []string{"aa", "bb", "cc", "aa"}},
			want: []string{"aa", "bb", "cc"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveDuplicate(tt.args.sliceList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveDuplicateInt(t *testing.T) {
	type args struct {
		sliceList []int
	}
	type testCase struct {
		name string
		args args
		want []int
	}
	tests := []testCase{
		// TODO: Add test cases.
		{
			name: "test2",
			args: args{sliceList: []int{44, 66, 26, 44}},
			want: []int{44, 66, 26},
		},
		{
			name: "test3",
			args: args{sliceList: []int{44, 66, 26, 44, 26}},
			want: []int{44, 66, 26},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveDuplicate(tt.args.sliceList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
