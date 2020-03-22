package soundcloud

import (
	"net/http"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func Test_service_Get(t *testing.T) {
	c := resty.New()
	e := struct{}{} // Empty struct to unmarshal to

	httpmock.ActivateNonDefault(c.GetClient())
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "/ok", httpmock.NewStringResponder(http.StatusOK, "{}"))
	httpmock.RegisterResponder("GET", "/nok", httpmock.NewStringResponder(http.StatusBadRequest, "{}"))

	s := service{client: c, path: "/ok"}
	assert.EqualError(t, s.Get(nil), "unmarshal: json: Unmarshal(nil)", "empty value as parameter returns an unmarshal error")
	assert.NoError(t, s.Get(&e), "should unmarshal an empty JSON into an empty struct")

	s = service{client: c, path: "/nok"}
	assert.EqualError(t, s.Get(&e), "query endpoint /nok: status code 400", "invalid code should yield an error")
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
