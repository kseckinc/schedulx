package router

import (
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
)

func TestBridgX_Login(t *testing.T) {
	type LoginInfo struct {
		UserName string `json:"username"`
		PassWord string `json:"password"`
	}
	type HttpResp struct {
		Code int64  `json:"code"`
		Msg  string `json:"msg"`
		Data string `json:"data"`
	}
	type args struct {
		Method string
		Url    string
		Param  *LoginInfo
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "UserLogin",
			args: args{
				Method: "POST",
				Url:    "http://bridgx-api.internal.galaxy-future.org/user/login",
				Param: &LoginInfo{
					UserName: "schedulx",
					PassWord: "123456",
				},
			},
		},
	}
	httpClient := resty.New().SetTimeout(2 * time.Second)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &HttpResp{}
			resp2, err := httpClient.R().SetBody(tt.args.Param).SetResult(resp).Post(tt.args.Url)
			t.Log(resp)
			t.Logf("resp2.Body:%s", resp2.Body())
			if err != nil {
				t.Fatal(err)
			}
			t.Log(err)
		})
	}

}
