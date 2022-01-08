package feedbin

import (
	"context"
	"net/http"
)

func (r *Feedbin) GetTaggings(ctx context.Context) (*GetTaggingsResp, error) {
	uri := "https://api.feedbin.com/v2/taggings.json"
	resp := new(GetTaggingsResp)

	_, err := r.request(ctx, http.MethodGet, uri, nil, true, &resp.Taggings)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type Tagging struct {
	ID     int    `json:"id"`
	FeedID int    `json:"feed_id"`
	Name   string `json:"name"`
}

type GetTaggingsResp struct {
	Taggings []*Tagging `json:"taggings"`
}
