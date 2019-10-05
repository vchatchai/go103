package linear

import "testing"

func Test_powerSeries(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{"test", args{1}, 1, 1},
		{"test", args{2}, 4, 8},
		{"test", args{3}, 9, 27},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := powerSeries(tt.args.a)
			if got != tt.want {
				t.Errorf("powerSeries() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("powerSeries() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
