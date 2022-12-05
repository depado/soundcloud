package soundcloud

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

const (
	PlaylistPath = "/playlists/{playlistID}"
)

var playlistIDRegex = regexp.MustCompile(`"soundcloud://playlists:(.*?)"`)

// Playlist will create a new playlist service
func (c *Client) Playlist() *PlaylistService {
	return &PlaylistService{
		service: service{
			client: c.c,
			path:   PlaylistPath,
		},
		extra: service{
			client: c.c,
			path:   "/tracks",
		},
	}
}

// PlaylistService is the structure holding information about the current
// playlist operation
type PlaylistService struct {
	service
	extra      service
	playlistID string
}

// WithID will set the current playlist ID to the given value
func (ps *PlaylistService) WithID(id string) *PlaylistService {
	ps.service.pathParams = map[string]string{"playlistID": id}
	ps.playlistID = id
	return ps
}

// FromURL will set the current playlist ID according to the provided URL
// This function will fetch the given URL and attempt to find a meta tag
// containing said ID
// If it fails, an error will always be returned
func (ps *PlaylistService) FromURL(url string) (*PlaylistService, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch page: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected return status: %d", resp.StatusCode)
	}

	o, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read body: %w", err)
	}
	out := playlistIDRegex.FindSubmatch(o)
	if len(out) != 2 {
		return nil, fmt.Errorf("no match was found")
	}
	id := string(out[1])

	ps.service.pathParams = map[string]string{"playlistID": id}
	ps.playlistID = id
	return ps, nil
}

// Get will fetch all the information about a playlist including all of its
// tracks. The first few tracks are already filled during the first request, the
// rest of the tracks needs to be fetched with another query
// WithID or FromURL must be called beforehand
func (ps *PlaylistService) Get() (*Playlist, error) {
	if ps.playlistID == "" {
		return nil, fmt.Errorf("no playlist ID or URL was provided")
	}

	pl := &Playlist{}
	if err := ps.service.Get(pl); err != nil {
		return pl, fmt.Errorf("get playlist %s: %w", ps.playlistID, err)
	}

	tracks := Tracks{}
	ids := []string{}
	for _, t := range pl.Tracks {
		if t.PermalinkURL == "" {
			ids = append(ids, strconv.Itoa(t.ID))
		} else {
			t.Playlist = pl
			tracks = append(tracks, t)
		}
	}

	chunks := [][]string{}
	chunkSize := 15
	for i := 0; i < len(ids); i += chunkSize {
		end := i + chunkSize
		if end > len(ids) {
			end = len(ids)
		}
		chunks = append(chunks, ids[i:end])
	}

	pl.Tracks = Tracks{}
	pl.Tracks = append(pl.Tracks, tracks...)
	for _, c := range chunks {
		ps.extra.queryParams = map[string]string{"ids": strings.Join(c, ",")}
		tr := &Tracks{}
		if err := ps.extra.Get(tr); err != nil {
			return pl, fmt.Errorf("get playlist %s: %w", ps.playlistID, err)
		}
		for _, t := range *tr {
			t.Playlist = pl
			pl.Tracks = append(pl.Tracks, t)
		}
	}

	return pl, nil
}
