package soundcloud

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
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
}

func (ts *TrackService) WithID(id string) *TrackService {
	ts.service.queryParams = map[string]string{"ids": id}
	ts.trackID = id

	return ts
}

func (ts *TrackService) FromURL(url string) (*TrackService, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch page: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected return status: %d", resp.StatusCode)
	}

	o, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read body: %w", err)
	}
	out := trackIDRegex.FindSubmatch(o)
	if len(out) != 2 {
		return nil, fmt.Errorf("no match was found")
	}
	id := string(out[1])

	ts.service.queryParams = map[string]string{"ids": id}
	ts.trackID = id
	return ts, nil
}

func (ts *TrackService) Get() (*Track, error) {
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

	return &tracks[0], nil
}

// Ongoing
func (ts *TrackService) StreamURL() (string, error) {
	t, err := ts.Get()
	if err != nil {
		return "", fmt.Errorf("get stream URL: %w", err)
	}
	for _, tc := range t.Media.Transcodings {
		fmt.Println(tc.URL + "?client_id=" + ts.clientID)
	}
	return "", nil
}
