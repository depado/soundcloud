package soundcloud

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/go-resty/resty/v2"
)

type service struct {
	client      *resty.Client
	path        string
	queryParams map[string]string
	pathParams  map[string]string
}

func (s service) Get(out interface{}) error {
	resp, err := s.client.R().SetPathParams(s.pathParams).SetQueryParams(s.queryParams).Get(s.path)
	if err != nil {
		return fmt.Errorf("query endpoint %s: %w", s.path, err)
	}

	if !resp.IsSuccess() {
		return fmt.Errorf("query endpoint %s: status code %d", s.path, resp.StatusCode())
	}

	if err = json.Unmarshal(resp.Body(), out); err != nil {
		return fmt.Errorf("unmarshal: %w", err)
	}

	return nil
}

func (s *service) Limit(limit int) {
	if s.queryParams == nil {
		s.queryParams = map[string]string{"limit": strconv.Itoa(limit)}
	} else {
		s.queryParams["limit"] = strconv.Itoa(limit)
	}
}

func (s *service) Offset(offset int) {
	if s.queryParams == nil {
		s.queryParams = map[string]string{"offset": strconv.Itoa(offset)}
	} else {
		s.queryParams["offset"] = strconv.Itoa(offset)
	}
}
