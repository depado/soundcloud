package soundcloud

import "fmt"

const (
	UserPath      = "/users/{userID}"
	UserLikesPath = UserPath + "/track_likes"
)

func (c *Client) User(id string) *UserService {
	return &UserService{
		service: service{
			client:     c.c,
			path:       "/users/{userID}",
			pathParams: map[string]string{"userID": id},
		},
		userID: id,
	}
}

type UserService struct {
	service
	userID string
}

func (us *UserService) Get() (*User, error) {
	u := &User{}
	if err := us.service.Get(u); err != nil {
		return u, fmt.Errorf("get user %s: %w", us.userID, err)
	}

	return u, nil
}

func (us *UserService) Limit(limit int) *UserService {
	us.service.Limit(limit)
	return us
}

func (us *UserService) Offset(offset int) *UserService {
	us.service.Offset(offset)
	return us
}

func (us *UserService) Likes() (Tracks, error) {
	us.service.path = UserLikesPath
	collection := &CollectionQuery{}

	if err := us.service.Get(collection); err != nil {
		return nil, fmt.Errorf("get tracks for user %s: %w", us.userID, err)
	}

	tracks := make([]Track, len(collection.Collection))
	for i, t := range collection.Collection {
		tracks[i] = t.Track
	}

	return tracks, nil
}

func (us *UserService) Playlists() ([]Playlist, error) {
	panic("not implemented")
}
