package env

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewConfig(t *testing.T) {
	type args struct {
		configPath string
	}
	tests := []struct {
		name      string
		args      args
		want      func(config IConfig)
		wantError assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{
				configPath: "testdata/.env.test.success",
			},
			want: func(config IConfig) {
				assert.Equal(t, "url", config.GetRepositoryURL())
				assert.Equal(t, "user", config.GetRepositoryUser())
				assert.Equal(t, "pass", config.GetRepositoryPass())
				assert.Equal(t, "/path/to/file", config.GetPasswordFile())
			},
			wantError: assert.NoError,
		},
		{
			name: "no url",
			args: args{
				configPath: "testdata/.env.test.nourl",
			},
			want: func(config IConfig) {
				assert.Nil(t, config)
			},
			wantError: assert.Error,
		},
		{
			name: "no user",
			args: args{
				configPath: "testdata/.env.test.nouser",
			},
			want: func(config IConfig) {
				assert.Nil(t, config)
			},
			wantError: assert.Error,
		},
		{
			name: "no pass",
			args: args{
				configPath: "testdata/.env.test.nopass",
			},
			want: func(config IConfig) {
				assert.Nil(t, config)
			},
			wantError: assert.Error,
		},
		{
			name: "no file",
			args: args{
				configPath: "testdata/.env.test.nofile",
			},
			want: func(config IConfig) {
				assert.Nil(t, config)
			},
			wantError: assert.Error,
		},
		{
			name: "env file not found",
			args: args{
				configPath: "testdata/.env.test.notfound",
			},
			want: func(config IConfig) {
				assert.Nil(t, config)
			},
			wantError: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewConfig(tt.args.configPath)
			if err != nil {
				tt.wantError(t, err)
			}
			tt.want(got)
		})
	}
}
