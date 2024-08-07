package util_test

import (
	"testing"

	"github.com/lostsnow/keqing/pkg/util"
)

func TestRandomString(t *testing.T) {
	t.Parallel()

	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"random-length", struct{ n int }{n: 10}, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := util.RandomString(tt.args.n); tt.args.n != tt.want {
				t.Errorf("RandomString() = %v, length want %v", got, tt.want)
			}
		})
	}
}
