package uniquepaths

import (
	"reflect"
	"testing"
)

func Test_maxProfitPath(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "simple test",
			args: args{grid: [][]int{
				{0, 2, 2, 1},
				{3, 1, 1, 1},
				{4, 4, 2, 0},
			}},
			want: [][]int{
				{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}, {2, 3},
			},
		},
		{
			name: "simple test 2",
			args: args{grid: [][]int{
				{0, 2, 2, 50},
				{3, 1, 1, 100},
				{4, 4, 2, 0},
			}},
			want: [][]int{
				{0, 0}, {0, 1}, {0, 2}, {0, 3}, {1, 3}, {2, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
