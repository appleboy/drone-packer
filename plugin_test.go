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
				},
			},
			want: exec.Command("packer", "validate", "foo.json"),
		},
		{
			name: "add vars flag",
			args: args{
				config: Config{
					Template: "foo.json",
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
			t.Parallel()
			if got := pkValidate(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pkValidate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pkBuild(t *testing.T) {
	type args struct {
		config Config
	}
	tests := []struct {
		name string
		args args
		want *exec.Cmd
	}{
		{
			name: "default build command",
			args: args{
				config: Config{
					Template: "foo.json",
				},
			},
			want: exec.Command("packer", "build", "foo.json"),
		},
		{
			name: "default build vars flag",
			args: args{
				config: Config{
					Template: "foo.json",
					Vars: map[string]string{
						"foo": "bar",
					},
					VarFiles: []string{"bar.json"},
				},
			},
			want: exec.Command("packer", "build", "-var-file", "bar.json", "-var", "foo=bar", "foo.json"),
		},
		{
			name: "add Parallel, Color and debug flag",
			args: args{
				config: Config{
					Template: "foo.json",
					Parallel: true,
					Color:    true,
					Debug:    true,
				},
			},
			want: exec.Command("packer", "build", "-parallel=true", "-color=true", "-debug", "foo.json"),
		},
		{
			name: "add machine readable flag",
			args: args{
				config: Config{
					Template: "foo.json",
					Readable: true,
				},
			},
			want: exec.Command("packer", "build", "-machine-readable", "foo.json"),
		},
		{
			name: "add force flag",
			args: args{
				config: Config{
					Template: "foo.json",
					Force:    true,
				},
			},
			want: exec.Command("packer", "build", "-force", "foo.json"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := pkBuild(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pkBuild() = %v, want %v", got, tt.want)
			}
		})
	}
}
