package fugle_marketdata

import (
	"reflect"
	"testing"
)

func TestNewRestClient(t *testing.T) {
	type args struct {
		option *RestClientOption
	}
	tests := []struct {
		name    string
		args    args
		want    *RestClient
		wantErr bool
	}{
		{
			name:    "invalid url then error",
			args:    args{option: &RestClientOption{Endpoint: "invalid", APIKey: "key"}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "empty api key then error",
			args:    args{option: &RestClientOption{Endpoint: "", APIKey: ""}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRestClient(tt.args.option)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRestClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("err: %v", err)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRestClient() got = %v, want %v", got, tt.want)
			}
		})
	}
}
