package password

import (
	"testing"
)

func TestCheck(t *testing.T) {
	type args struct {
		password string
		hash     string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test",
			args: args{
				password: "qwerty",
				hash:     "3f6223822c94b6e4cde23eb93c8e0dea356f06b103fef4abc2bd774a52dda17e$cbabc2a614ca2c75b02981620f0e614363cb283fc88d23e620a1090af58a25ea28a643a69a015d14bbeb7003a7f1dd0bc7c80e6df84f42528eeda0e8f3a83e44",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Check(tt.args.password, tt.args.hash); got != tt.want {
				t.Errorf("Check() = %v, want %v", got, tt.want)
			}
		})
	}
}
