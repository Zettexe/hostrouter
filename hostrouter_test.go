package hostrouter

import (
	"slices"
	"testing"
)

func Test_getWildcardHost(t *testing.T) {
	type args struct {
		host string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"no wildcard in 1-part host", args{"com"}, []string{"com"}},
		{"wildcard in 2-part", args{"dot.com"}, []string{"*.com", "dot.*"}},
		{"wildcard in 3-part", args{"amazing.dot.com"}, []string{"*.dot.com", "amazing.*.com", "amazing.dot.*"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getWildcardHosts(tt.args.host); !slices.Equal(got, tt.want) {
				t.Errorf("getWildcardHosts() = %v, want %v", got, tt.want)
			}
		})
	}
}
