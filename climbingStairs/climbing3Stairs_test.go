package climbingstairs

import "testing"

func Test_climbStairs3Steps(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "a simple test 1",
			args: args{
				n: 3,
			},
			want: 4,
		},
		{
			name: "a simple test 2",
			args: args{
				n: 5,
			},
			want: 13,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climb3Stairs(tt.args.n); got != tt.want {
				t.Errorf("climbStairs() = %v , want %v", got, tt.want)
			}
		})
	}
}
