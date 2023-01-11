package feedbin

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func (r *Feedbin) request(ctx context.Context, method, uri string, body io.Reader, withAuth bool, resp interface{}) (string, error) {
	req, err := http.NewRequest(method, uri, body)
	if err != nil {
		return "", err
	}
	if withAuth && r.password != "" {
		req.SetBasicAuth(r.username, r.password)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("User-Agent", fmt.Sprintf("go-feedbin/%s (https://github.com/chyroc/go-feedbin)", version))

	response, err := r.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	bs, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	if response.StatusCode == http.StatusUnauthorized {
		if s := string(bs); s != "" {
			return "", fmt.Errorf(strings.TrimSpace(s))
		}
		return "", fmt.Errorf(strings.TrimSpace(response.Status))
	}

	res := new(baseResp)
	if _ = json.Unmarshal(bs, res); res.Err() != nil {
		return "", res.Err()
	}

	if resp != nil {
		if err = json.Unmarshal(bs, resp); err != nil {
			return "", err
		}
	}

	return string(bs), nil
}

type baseResp struct {
	Status   int         `json:"status"`
	Message  interface{} `json:"message"`
	Messages string      `json:"messages"`
	Errors   []struct {
		Pages string `json:"pages"`
	} `json:"errors"`
}

func (r *baseResp) Err() error {
	if r.Messages != "" {
		return fmt.Errorf(r.Messages)
	}
	for _, v := range r.Errors {
		if v.Pages != "" {
			return fmt.Errorf(v.Pages)
		}
	}
	return nil
}
