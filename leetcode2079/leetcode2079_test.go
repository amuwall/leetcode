package leetcode2079

import "testing"

func Test_wateringPlants(t *testing.T) {
	type args struct {
		plants   []int
		capacity int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				plants:   []int{2, 2, 3, 3},
				capacity: 5,
			},
			want: 14,
		},
		{
			name: "example 2",
			args: args{
				plants:   []int{1, 1, 1, 4, 2, 3},
				capacity: 4,
			},
			want: 30,
		},
		{
			name: "example 3",
			args: args{
				plants:   []int{7, 7, 7, 7, 7, 7, 7},
				capacity: 8,
			},
			want: 49,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wateringPlants(tt.args.plants, tt.args.capacity); got != tt.want {
				t.Errorf("wateringPlants() = %v, want %v", got, tt.want)
			}
		})
	}
}
