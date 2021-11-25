package tool

import "testing"

func TestPickDomainFromUrl(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test-Http",
			args: args{
				url: "http://192.168.0.1:81/abc/123",
			},
			want:    "192.168.0.1:81",
			wantErr: false,
		},
		{
			name: "Test-Https",
			args: args{
				url: "https://192.168.0.1:81/abc/123",
			},
			want:    "192.168.0.1:81",
			wantErr: false,
		},
		{
			name: "Test-Error",
			args: args{
				url: "http:/",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PickDomainFromUrl(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("PickDomainFromUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PickDomainFromUrl() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubStr(t *testing.T) {
	type args struct {
		s      string
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "123456789",
			args: args{
				s:      "123456789",
				length: 5,
			},
			want: "12345",
		},
		{
			name: "中文字符截断测试",
			args: args{
				s:      "中文字符截断测试",
				length: 5,
			},
			want: "中文字符截",
		},
		{
			name: "123中文测试",
			args: args{
				s:      "123中文测试",
				length: 5,
			},
			want: "123中文",
		},
		{
			name: "123中",
			args: args{
				s:      "123中",
				length: 5,
			},
			want: "123中",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubStr(tt.args.s, tt.args.length); got != tt.want {
				t.Errorf("SubStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
