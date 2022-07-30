package print

import "testing"

func TestFormat(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test for 1", args{1}, "The number is 1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Format(tt.args.num); got != tt.want {
				t.Errorf("Format() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpace(t *testing.T) {
	type args struct {
		length int
		parts  []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"empty parts", args{10, []string{}}, ""},
		{"one element parts", args{10, []string{"hello"}}, "hello"},
		{"two element parts", args{10, []string{"hi", "you"}}, "hi     you"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Space(tt.args.length, tt.args.parts...); got != tt.want {
				t.Errorf("Space() = %v, want %v", got, tt.want)
			}
		})
	}
}
