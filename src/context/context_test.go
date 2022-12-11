package context

import (
	"github.com/kengo-k/password-manager/context/runmode"
	"github.com/kengo-k/password-manager/env"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestNewContext(t *testing.T) {
	type args struct {
		mode   runmode.RunMode
		config env.IConfig
	}
	tests := []struct {
		name string
		args args
		want func(context Context)
	}{
		{
			name: "file,file",
			args: args{
				mode: runmode.FILE_TO_FILE,
			},
			want: func(context Context) {
				load1 := reflect.ValueOf(loadFile).Pointer()
				load2 := reflect.ValueOf(context.LoadFn).Pointer()
				save1 := reflect.ValueOf(saveFile).Pointer()
				save2 := reflect.ValueOf(context.SaveFn).Pointer()
				assert.Equal(t, load1, load2)
				assert.Equal(t, save1, save2)
			},
		},
		{
			name: "git,git",
			args: args{
				mode: runmode.GIT_TO_GIT,
			},
			want: func(context Context) {
				load1 := reflect.ValueOf(loadRepository).Pointer()
				load2 := reflect.ValueOf(context.LoadFn).Pointer()
				save1 := reflect.ValueOf(saveRepository).Pointer()
				save2 := reflect.ValueOf(context.SaveFn).Pointer()
				assert.Equal(t, load1, load2)
				assert.Equal(t, save1, save2)
			},
		},
		{
			name: "git,file",
			args: args{
				mode: runmode.GIT_TO_FILE,
			},
			want: func(context Context) {
				load1 := reflect.ValueOf(loadRepository).Pointer()
				load2 := reflect.ValueOf(context.LoadFn).Pointer()
				save1 := reflect.ValueOf(saveFile).Pointer()
				save2 := reflect.ValueOf(context.SaveFn).Pointer()
				assert.Equal(t, load1, load2)
				assert.Equal(t, save1, save2)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewContext(tt.args.mode, tt.args.config)
			context := got.(*Context)
			tt.want(*context)
		})
	}
}

func Test_loadFile(t *testing.T) {

	tests := []struct {
		name      string
		getConfig func() env.IConfig
		want      func([]string)
		wantErr   assert.ErrorAssertionFunc
	}{
		{
			name: "password not found",
			getConfig: func() env.IConfig {
				config, _ := env.NewConfig("testdata/.test.notfound.env")
				return config
			},
			wantErr: assert.Error,
		},
		{
			name: "success",
			getConfig: func() env.IConfig {
				config, _ := env.NewConfig("testdata/.test.loadfile.env")
				return config
			},
			wantErr: assert.NoError,
			want: func(lines []string) {
				assert.Equal(t, 2, len(lines))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := tt.getConfig()
			got, err := loadFile(config)
			if tt.wantErr != nil {
				tt.wantErr(t, err)
			}
			if tt.want != nil {
				tt.want(got)
			}
		})
	}
}
