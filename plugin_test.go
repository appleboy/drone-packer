package main

import (
	"os/exec"
	"reflect"
	"testing"
)

func Test_pkValidate(t *testing.T) {
	type args struct {
		config Config
	}
	tests := []struct {
		name string
		args args
		want *exec.Cmd
	}{
		{
			name: "default validate command",
			args: args{
				config: Config{
					Template: "foo.json",
					Actions:  []string{"validate"},
				},
			},
			want: exec.Command("packer", "validate", "foo.json"),
		},
		{
			name: "add vars flag",
			args: args{
				config: Config{
					Template: "foo.json",
					Actions:  []string{"validate"},
					Vars: map[string]string{
						"foo": "bar",
					},
					VarFiles: []string{"bar.json"},
				},
			},
			want: exec.Command("packer", "validate", "-var-file", "bar.json", "-var", "foo=bar", "foo.json"),
		},
		{
			name: "add except only color flag",
			args: args{
				config: Config{
					Template:   "foo.json",
					Actions:    []string{"validate"},
					Except:     []string{"foo", "bar"},
					Only:       []string{"a", "b"},
					SyntaxOnly: true,
				},
			},
			want: exec.Command("packer", "validate", "-except=foo,bar", "-only=a,b", "-syntax-only", "foo.json"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pkValidate(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pkValidate() = %v, want %v", got, tt.want)
			}
		})
	}
}
