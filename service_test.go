package soundcloud

import (
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func Test_service_Get(t *testing.T) {
	type fields struct {
		client      *resty.Client
		queryParams map[string]string
		path        string
		pathParams  map[string]string
	}
	type args struct {
		out interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := service{
				client:      tt.fields.client,
				queryParams: tt.fields.queryParams,
				path:        tt.fields.path,
				pathParams:  tt.fields.pathParams,
			}
			if err := s.Get(tt.args.out); (err != nil) != tt.wantErr {
				t.Errorf("service.Get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_Limit(t *testing.T) {
	s := service{}
	assert.NotContains(t, s.queryParams, "limit", "query params should be empty by default")
	s.Limit(1)
	assert.Contains(t, s.queryParams, "limit", "limit should be contained")
	assert.Equal(t, s.queryParams["limit"], "1", "limit should be equal to value")
	s.Limit(2)
	assert.Contains(t, s.queryParams, "limit", "limit should be contained")
	assert.Equal(t, s.queryParams["limit"], "2", "limit should be equal to value")
}

func Test_service_Offset(t *testing.T) {
	s := service{}
	assert.NotContains(t, s.queryParams, "offset", "query params should be empty by default")
	s.Offset(1)
	assert.Contains(t, s.queryParams, "offset", "offset should be contained")
	assert.Equal(t, s.queryParams["offset"], "1", "offset should be equal to value")
	s.Offset(2)
	assert.Contains(t, s.queryParams, "offset", "offset should be contained")
	assert.Equal(t, s.queryParams["offset"], "2", "offset should be equal to value")
}
