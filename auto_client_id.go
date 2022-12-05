package soundcloud

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"

	"golang.org/x/net/html"
)

type multiError []error

func (e multiError) Error() string {
	var str string

	for _, err := range e {
		if len(str) == 0 {
			str = err.Error()
		}
		str = fmt.Sprintf("%s ; %s", str, err.Error())
	}

	return str
}

var clientIDRE = regexp.MustCompile(`,client_id:"(.*?)"`)

func find() ([]string, error) {
	resp, err := http.Get("https://soundcloud.com")
	if err != nil {
		return nil, fmt.Errorf("find: %w", err)
	}
	defer resp.Body.Close()

	return extractScripts(resp.Body), nil
}

func extractScripts(r io.Reader) []string {
	var scripts []string

	z := html.NewTokenizer(r)

	for t := z.Next(); t != html.ErrorToken; t = z.Next() {
		if t == html.StartTagToken {
			tok := z.Token()
			for _, a := range tok.Attr {
				if a.Key == "src" {
					scripts = append(scripts, a.Val)
				}
			}
		}
	}

	return scripts
}

// NewClientIDFromPublicHTML allow to retrieve a token from the public page
// of SoundCloud.
func NewClientIDFromPublicHTML() (string, error) {
	var token string
	scripts, err := find()
	if err != nil {
		return "", fmt.Errorf("new token from public HTML: %w", err)
	}

	ec := make(chan error)
	tc := make(chan string)
	s := len(scripts)

	for _, script := range scripts {
		go func(url string, c chan<- string, e chan<- error) {
			resp, err := http.Get(url)
			if err != nil {
				e <- err
				return
			}
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				e <- err
				return
			}

			f := clientIDRE.FindSubmatch(body)
			if f != nil {
				c <- string(f[1])
			} else {
				e <- nil
			}
		}(script, tc, ec)
	}

	var errs = multiError{}
	for i := 0; i < s; i++ {
		select {
		case e := <-ec:
			if e != nil {
				errs = append(errs, e)
			}
		case t := <-tc:
			token = t
		}
	}

	if token == "" {
		if len(errs) > 0 {
			return "", errs
		} else {
			return "", errors.New("no token was found")
		}
	}

	return token, nil
}
