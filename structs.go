package soundcloud

import "time"

type User struct {
	AvatarURL            string      `json:"avatar_url"`
	City                 interface{} `json:"city"`
	CommentsCount        int         `json:"comments_count"`
	CountryCode          interface{} `json:"country_code"`
	CreatedAt            time.Time   `json:"created_at"`
	CreatorSubscriptions []struct {
		Product struct {
			ID string `json:"id"`
		} `json:"product"`
	} `json:"creator_subscriptions"`
	CreatorSubscription struct {
		Product struct {
			ID string `json:"id"`
		} `json:"product"`
	} `json:"creator_subscription"`
	Description        interface{} `json:"description"`
	FollowersCount     int         `json:"followers_count"`
	FollowingsCount    int         `json:"followings_count"`
	FirstName          string      `json:"first_name"`
	FullName           string      `json:"full_name"`
	GroupsCount        int         `json:"groups_count"`
	ID                 int         `json:"id"`
	Kind               string      `json:"kind"`
	LastModified       time.Time   `json:"last_modified"`
	LastName           string      `json:"last_name"`
	LikesCount         int         `json:"likes_count"`
	PlaylistLikesCount int         `json:"playlist_likes_count"`
	Permalink          string      `json:"permalink"`
	PermalinkURL       string      `json:"permalink_url"`
	PlaylistCount      int         `json:"playlist_count"`
	RepostsCount       interface{} `json:"reposts_count"`
	TrackCount         int         `json:"track_count"`
	URI                string      `json:"uri"`
	Urn                string      `json:"urn"`
	Username           string      `json:"username"`
	Verified           bool        `json:"verified"`
	Visuals            struct {
		Urn     string `json:"urn"`
		Enabled bool   `json:"enabled"`
		Visuals []struct {
			Urn       string `json:"urn"`
			EntryTime int    `json:"entry_time"`
			VisualURL string `json:"visual_url"`
		} `json:"visuals"`
	} `json:"visuals"`
}

type Track struct {
	CommentCount int       `json:"comment_count"`
	FullDuration int       `json:"full_duration"`
	Downloadable bool      `json:"downloadable"`
	CreatedAt    time.Time `json:"created_at"`
	Description  string    `json:"description"`
	Media        struct {
		Transcodings []struct {
			URL      string `json:"url"`
			Preset   string `json:"preset"`
			Duration int    `json:"duration"`
			Snipped  bool   `json:"snipped"`
			Format   struct {
				Protocol string `json:"protocol"`
				MimeType string `json:"mime_type"`
			} `json:"format"`
			Quality string `json:"quality"`
		} `json:"transcodings"`
	} `json:"media"`
	Title             string `json:"title"`
	PublisherMetadata struct {
		Urn           string `json:"urn"`
		ContainsMusic bool   `json:"contains_music"`
		ID            int    `json:"id"`
	} `json:"publisher_metadata"`
	Duration          int         `json:"duration"`
	HasDownloadsLeft  bool        `json:"has_downloads_left"`
	ArtworkURL        string      `json:"artwork_url"`
	Public            bool        `json:"public"`
	Streamable        bool        `json:"streamable"`
	TagList           string      `json:"tag_list"`
	DownloadURL       interface{} `json:"download_url"`
	Genre             string      `json:"genre"`
	ID                int         `json:"id"`
	RepostsCount      int         `json:"reposts_count"`
	State             string      `json:"state"`
	LabelName         interface{} `json:"label_name"`
	LastModified      time.Time   `json:"last_modified"`
	Commentable       bool        `json:"commentable"`
	Policy            string      `json:"policy"`
	Visuals           interface{} `json:"visuals"`
	Kind              string      `json:"kind"`
	PurchaseURL       interface{} `json:"purchase_url"`
	Sharing           string      `json:"sharing"`
	URI               string      `json:"uri"`
	SecretToken       interface{} `json:"secret_token"`
	DownloadCount     int         `json:"download_count"`
	LikesCount        int         `json:"likes_count"`
	Urn               string      `json:"urn"`
	License           string      `json:"license"`
	PurchaseTitle     interface{} `json:"purchase_title"`
	DisplayDate       time.Time   `json:"display_date"`
	EmbeddableBy      string      `json:"embeddable_by"`
	ReleaseDate       interface{} `json:"release_date"`
	UserID            int         `json:"user_id"`
	MonetizationModel string      `json:"monetization_model"`
	WaveformURL       string      `json:"waveform_url"`
	Permalink         string      `json:"permalink"`
	PermalinkURL      string      `json:"permalink_url"`
	User              User        `json:"user"`
	PlaybackCount     int         `json:"playback_count"`
}

type Tracks []Track

type PlaylistsLikedOwned struct {
	Collection []Collection `json:"collection"`
	NextHref   string       `json:"next_href"`
	QueryUrn   interface{}  `json:"query_urn"`
}

type Playlist struct {
	ManagedByFeeds bool        `json:"managed_by_feeds"`
	Kind           string      `json:"kind"`
	SetType        string      `json:"set_type"`
	CreatedAt      time.Time   `json:"created_at"`
	Sharing        string      `json:"sharing"`
	Title          string      `json:"title"`
	TrackCount     int         `json:"track_count"`
	URI            string      `json:"uri"`
	SecretToken    interface{} `json:"secret_token"`
	Duration       int         `json:"duration"`
	LikesCount     int         `json:"likes_count"`
	ArtworkURL     string      `json:"artwork_url"`
	Public         bool        `json:"public"`
	DisplayDate    time.Time   `json:"display_date"`
	UserID         int         `json:"user_id"`
	IsAlbum        bool        `json:"is_album"`
	ID             int         `json:"id"`
	RepostsCount   int         `json:"reposts_count"`
	Permalink      string      `json:"permalink"`
	PermalinkURL   string      `json:"permalink_url"`
	PublishedAt    time.Time   `json:"published_at"`
	LastModified   time.Time   `json:"last_modified"`
	User           User        `json:"user"`
}

type Collection struct {
	Playlist  Playlist  `json:"playlist"`
	CreatedAt time.Time `json:"created_at"`
	Type      string    `json:"type"`
	User      User      `json:"user"`
	Track     Track     `json:"track"`
	UUID      string    `json:"uuid"`
}

type PlayHistory struct {
	Collection []struct {
		PlayedAt int64 `json:"played_at"`
		TrackID  int   `json:"track_id"`
		Track    Track `json:"track"`
	} `json:"collection"`
	NextHref string      `json:"next_href"`
	QueryUrn interface{} `json:"query_urn"`
}

type CollectionQuery struct {
	Collection []Collection `json:"collection"`
	NextHref   string       `json:"next_href"`
	QueryUrn   interface{}  `json:"query_urn"`
}
type Format struct {
	Protocol string `json:"protocol"`
	MimeType string `json:"mime_type"`
}
type Transcodings struct {
	URL      string `json:"url"`
	Preset   string `json:"preset"`
	Duration int    `json:"duration"`
	Snipped  bool   `json:"snipped"`
	Format   Format `json:"format"`
	Quality  string `json:"quality"`
}
type Media struct {
	Transcodings []Transcodings `json:"transcodings"`
}
type PublisherMetadata struct {
	Urn string `json:"urn"`
	ID  int    `json:"id"`
}
