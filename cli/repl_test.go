package cli

import (
	"reflect"
	"testing"
)

func Test_cleanInput(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				input: "hello world",
			},
			want: []string{"hello", "world"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cleanInput(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cleanInput() = %v, want %v", got, tt.want)
			}
		})
	}
}
