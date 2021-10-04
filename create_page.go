package feedbin

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

// CreatePage can be used to [create a new entry](https://feedbin.com/blog/2019/08/20/save-webpages-to-read-later/) from the URL of an article.
//
// If successful, the response will be the full [entry](https://github.com/feedbin/feedbin-api/blob/master/content/entries.md).
func (r *Feedbin) CreatePage(ctx context.Context, request *CreatePageReq) (*CreatePageResp, error) {
	bs, _ := json.Marshal(request)
	uri := "https://api.feedbin.com/v2/pages.json"

	resp := new(CreatePageResp)
	_, err := r.request(ctx, http.MethodPost, uri, bytes.NewReader(bs), true, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
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
