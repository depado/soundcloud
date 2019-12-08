package soundcloud

import (
	"testing"

	"github.com/go-resty/resty/v2"
)

func Test_service_Limit(t *testing.T) {
	type fields struct {
		client      *resty.Client
		queryParams map[string]string
		path        string
		pathParams  map[string]string
	}
	type args struct {
		limit int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				client:      tt.fields.client,
				queryParams: tt.fields.queryParams,
				path:        tt.fields.path,
				pathParams:  tt.fields.pathParams,
			}
			s.Limit(tt.args.limit)
		})
	}
}

func Test_service_Offset(t *testing.T) {
	type fields struct {
		client      *resty.Client
		queryParams map[string]string
		path        string
		pathParams  map[string]string
	}
	type args struct {
		offset int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				client:      tt.fields.client,
				queryParams: tt.fields.queryParams,
				path:        tt.fields.path,
				pathParams:  tt.fields.pathParams,
			}
			s.Offset(tt.args.offset)
		})
	}
}
