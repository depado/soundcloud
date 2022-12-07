package soundcloud

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-resty/resty/v2"
)

const (
	TrackPath = "/tracks"
)

var trackIDRegex = regexp.MustCompile(`"soundcloud://sounds:(.*?)"`)

func (c *Client) Track() *TrackService {
	return &TrackService{
		service: service{
			client: c.c,
			path:   TrackPath,
		},
		clientID: c.clientID,
	}
}

type TrackService struct {
	service
	clientID string
	trackID  string
	current  *Track
}

func (ts *TrackService) WithID(id string) *TrackService {
	ts.service.queryParams = map[string]string{"ids": id}
	ts.trackID = id

	return ts
}

func (ts *TrackService) FromTrack(t *Track, fetch bool) (*TrackService, *Track, error) {
	ts.trackID = strconv.Itoa(t.ID)
	ts.current = t
	if !fetch {
		return ts, t, nil
	}
	t, err := ts.get()
	return ts, t, err
}

func (ts *TrackService) FromURL(url string) (*TrackService, *Track, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to fetch page: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, nil, fmt.Errorf("unexpected return status: %d", resp.StatusCode)
	}

	o, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to read body: %w", err)
	}
	out := trackIDRegex.FindSubmatch(o)
	if len(out) != 2 {
		return nil, nil, fmt.Errorf("no match was found")
	}
	id := string(out[1])

	ts.service.queryParams = map[string]string{"ids": id}
	ts.trackID = id
	t, err := ts.get()
	return ts, t, err
}

func (ts *TrackService) get() (*Track, error) {
	if ts.trackID == "" {
		return nil, fmt.Errorf("get track: no track ID or URL was provided")
	}

	tracks := Tracks{}
	if err := ts.service.Get(&tracks); err != nil {
		return nil, fmt.Errorf("get track %s: %w", ts.trackID, err)
	}

	if len(tracks) < 1 {
		return nil, fmt.Errorf("get track %s: array of size 0", ts.trackID)
	}

	ts.current = &tracks[0]

	return ts.current, nil
}

func (ts *TrackService) Stream(st StreamType) (string, error) {
	if ts.current == nil || len(ts.current.Media.Transcodings) == 0 {
		return "", fmt.Errorf("no track assigned or no transcoding found")
	}

	var furl string
	for _, tc := range ts.current.Media.Transcodings {
		if strings.HasPrefix(tc.Preset, st.PresetPrefix) && tc.Format.Protocol == st.Protocol {
			furl = tc.URL
			break
		}
	}
	if furl == "" {
		return "", fmt.Errorf("no matching stream type found")
	}

	resp, err := resty.New().SetQueryParam("client_id", ts.clientID).R().Get(furl)
	if err != nil {
		return "", fmt.Errorf("query endpoint %s: %w", furl, err)
	}

	if !resp.IsSuccess() {
		return "", fmt.Errorf("query endpoint %s: status code %d", furl, resp.StatusCode())
	}
	out := struct {
		URL string `json:"url"`
	}{}
	if err = json.Unmarshal(resp.Body(), &out); err != nil {
		return "", fmt.Errorf("unmarshal: %w", err)
	}

	return out.URL, nil
}
