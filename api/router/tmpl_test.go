package router

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/galaxy-future/schedulx/api/handler"
	"github.com/galaxy-future/schedulx/api/types"
	"github.com/galaxy-future/schedulx/service"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

func TestTmplate_Create(t *testing.T) {
	type args struct {
		Method string
		Url    string
		Param  *handler.TmplExpandRequest
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TemplateCreate pier",
			args: args{
				Method: "POST",
				Url:    "/api/v1/schedulx/template/expand/create",
				Param: &handler.TmplExpandRequest{
					EndStep: service.ExpandStepMount,
					TmplInfo: &types.TmpInfo{
						TmplName:         "pier 扩容模板_tmpl_info",
						ServiceClusterId: 7,
						Describe:         "pier 阿里云扩容模板",
						BridgxClusname:   "gf.scheduler.test",
					},
					BaseEnv: &types.BaseEnv{
						IsContainer: true,
					},
					ServiceEnv: &types.ServiceEnv{
						ImageStorageType: "harbor",
						ImageUrl:         "172.16.16.172:12380/dtexpress/prod/pier:0.1.5",
						Port:             80,
						Cmd:              "docker run -d --net=host -p 80:80 --name gf.dtexpress.pier",
					},
					Mount: &types.ParamsMount{
						MountType:  "alb",
						MountValue: "sgp-u11hcmcfay3h226ryg",
					},
				},
			},
		},
		{
			name: "TemplateCreate downloader",
			args: args{
				Method: "POST",
				Url:    "/api/v1/schedulx/template/expand/create",
				Param: &handler.TmplExpandRequest{
					EndStep: service.ExpandStepTmplInfo,
					TmplInfo: &types.TmpInfo{
						TmplName:         "downloader 扩容模板_tmpl_info",
						ServiceClusterId: 7,
						Describe:         "downloader 阿里云扩容模板",
						BridgxClusname:   "gf.scheduler.test",
					},
					BaseEnv: &types.BaseEnv{
						IsContainer: true,
					},
					ServiceEnv: &types.ServiceEnv{
						ImageStorageType: "harbor",
						ImageUrl:         "172.16.48.179:12380/dtexpress/prod/downloader:0.1.15",
						Port:             80,
						Cmd:              "docker run -d --net=host -p 80:80 --name gf.dtexpress.downloader",
					},
					Mount: &types.ParamsMount{
						MountType:  "alb",
						MountValue: "xxx",
					},
				},
			},
		},
		{
			name: "TemplateCreate net_detector",
			args: args{
				Method: "POST",
				Url:    "/api/v1/schedulx/template/expand/create",
				Param: &handler.TmplExpandRequest{
					EndStep: service.ExpandStepTmplInfo,
					TmplInfo: &types.TmpInfo{
						TmplName:         "net_detector 扩容模板_tmpl_info",
						ServiceClusterId: 7,
						Describe:         "net_detector 阿里云扩容模板",
						BridgxClusname:   "gf.scheduler.test",
					},
					BaseEnv: &types.BaseEnv{
						IsContainer: true,
					},
					ServiceEnv: &types.ServiceEnv{
						ImageStorageType: "harbor",
						ImageUrl:         "172.16.16.172:12380/detector/boe/gf.rd.net_detector:1.0.0",
						Port:             80,
						Cmd:              "docker run -d --net=host --name gf.rd.net_detector",
					},
					Mount: &types.ParamsMount{
						MountType:  "alb",
						MountValue: "xxx",
					},
				},
			},
		},
	}

	r := Init()
	w := httptest.NewRecorder()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonByte, _ := jsoniter.Marshal(tt.args.Param)
			req, _ := http.NewRequest(tt.args.Method, tt.args.Url, bytes.NewReader(jsonByte))
			req.Header.Set("Authorization", "Bearer "+authToken)
			r.ServeHTTP(w, req)
			t.Log(w.Body.String())
			assert.Equal(t, 200, w.Code)
		})
	}
}
