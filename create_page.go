package feedbin

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// CreatePage can be used to [create a new entry](https://feedbin.com/blog/2019/08/20/save-webpages-to-read-later/) from the URL of an article.
//
// If successful, the response will be the full [entry](https://github.com/feedbin/feedbin-api/blob/master/content/entries.md).
func (r *Feedbin) CreatePage(ctx context.Context, request *CreatePageReq) (*CreatePageResp, error) {
	bs, _ := json.Marshal(request)
	req, err := http.NewRequest(http.MethodPost, "https://api.feedbin.com/v2/pages.json", bytes.NewBuffer(bs))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(r.username, r.password)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	response, err := r.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	bs, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode == http.StatusUnauthorized {
		if s := string(bs); s != "" {
			return nil, fmt.Errorf(strings.TrimSpace(s))
		}
		return nil, fmt.Errorf(strings.TrimSpace(response.Status))
	}

	resp := new(createPageResp)
	if err = json.Unmarshal(bs, resp); err != nil {
		return nil, err
	} else if resp.Err() != nil {
		return nil, resp.Err()
	}

	return &resp.CreatePageResp, nil
}

type CreatePageReq struct {
	URL   string `json:"url,omitempty"`
	Title string `json:"title,omitempty"` // The title is optional and will only be used if Feedbin cannot find the title of the content.
}

type CreatePageResp struct {
	ID                  int64       `json:"id"`
	FeedID              int         `json:"feed_id"`
	Title               string      `json:"title"`
	Author              interface{} `json:"author"`
	Summary             string      `json:"summary"`
	Content             string      `json:"content"`
	URL                 string      `json:"url"`
	ExtractedContentURL string      `json:"extracted_content_url"`
	Published           time.Time   `json:"published"`
	CreatedAt           time.Time   `json:"created_at"`
}

type baseResp struct {
	Status  int         `json:"status"`
	Message interface{} `json:"message"`
	Errors  []struct {
		Pages string `json:"pages"`
	} `json:"errors"`
}

type createPageResp struct {
	baseResp
	CreatePageResp
}

func (r *baseResp) Err() error {
	for _, v := range r.Errors {
		if v.Pages != "" {
			return fmt.Errorf(v.Pages)
		}
	}
	return nil
}
