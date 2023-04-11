package main

import (
	"testing"
)

func Test_pkValidate(t *testing.T) {
	type args struct {
		config Config
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "default validate command",
			args: args{
				config: Config{
					Template: "foo.json",
				},
			},
			want: "packer validate foo.json",
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
			want: "packer validate -var-file=bar.json -var foo=bar foo.json",
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
			want: "packer validate -except=foo,bar -only=a,b -syntax-only foo.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pkValidate(tt.args.config); got.String() != tt.want {
				t.Errorf("pkValidate() = %v, want %v", got.String(), tt.want)
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
		want string
	}{
		{
			name: "default build command",
			args: args{
				config: Config{
					Template: "foo.json",
				},
			},
			want: "packer build foo.json",
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
			want: "packer build -var-file=bar.json -var foo=bar foo.json",
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
			want: "packer build -parallel=true -color=true -debug foo.json",
		},
		{
			name: "add machine readable flag",
			args: args{
				config: Config{
					Template: "foo.json",
					Readable: true,
				},
			},
			want: "packer build -machine-readable foo.json",
		},
		{
			name: "add force flag",
			args: args{
				config: Config{
					Template: "foo.json",
					Force:    true,
				},
			},
			want: "packer build -force foo.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pkBuild(tt.args.config); got.String() != tt.want {
				t.Errorf("pkBuild() = %v, want %v", got.String(), tt.want)
			}
		})
	}
}

func Test_pkInit(t *testing.T) {
	type args struct {
		config Config
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "default command",
			args: args{
				config: Config{
					Template: "aws-ubuntu.pkr.hcl",
				},
			},
			want: "packer init aws-ubuntu.pkr.hcl",
		},
		{
			name: "upgrade command",
			args: args{
				config: Config{
					Template:  "aws-ubuntu.pkr.hcl",
					IsUpgrade: true,
				},
			},
			want: "packer init -upgrade aws-ubuntu.pkr.hcl",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pkInit(tt.args.config); got.String() != tt.want {
				t.Errorf("pkInit() = %v, want %v", got, tt.want)
			}
		})
	}
}
