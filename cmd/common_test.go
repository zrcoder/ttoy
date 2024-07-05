package cmd

import "testing"

func Test_extWithoutDot(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		{
			path: "xx.yaml",
			want: "yaml",
		},
		{
			path: "xx.json",
			want: "json",
		},
		{
			path: "xx.toml",
			want: "toml",
		},
		{
			path: "yy/xx.json",
			want: "json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extWithoutDot(tt.path); got != tt.want {
				t.Errorf("extWithoutDot() = %v, want %v", got, tt.want)
			}
		})
	}
}
