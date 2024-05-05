package leetcode1652

import (
	"reflect"
	"testing"
)

func Test_decrypt(t *testing.T) {
	type args struct {
		code []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example 1",
			args: args{
				code: []int{5, 7, 1, 4},
				k:    3,
			},
			want: []int{12, 10, 16, 13},
		},
		{
			name: "example 2",
			args: args{
				code: []int{1, 2, 3, 4},
				k:    0,
			},
			want: []int{0, 0, 0, 0},
		},
		{
			name: "example 3",
			args: args{
				code: []int{2, 4, 9, 3},
				k:    -2,
			},
			want: []int{12, 5, 6, 13},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decrypt(tt.args.code, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
