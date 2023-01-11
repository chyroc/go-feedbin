package feedbin

import (
	"context"
	"net/http"
	"net/url"
	"time"
)

func (r *Feedbin) GetSubscriptions(ctx context.Context, request *GetSubscriptionsReq) (*GetSubscriptionsResp, error) {
	uri, _ := url.Parse("https://api.feedbin.com/v2/subscriptions.json")
	q := uri.Query()
	if !request.Since.IsZero() {
		q.Set("since", request.Since.Format(time.RFC3339Nano))
	}
	if request.Mode != "" {
		q.Set("mode", request.Mode)
	}
	uri.RawQuery = q.Encode()
	resp := new(GetSubscriptionsResp)

	_, err := r.request(ctx, http.MethodGet, uri.String(), nil, true, &resp.Subscriptions)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type Subscription struct {
	ID        int         `json:"id"`
	CreatedAt time.Time   `json:"created_at"`
	FeedID    int         `json:"feed_id"`
	Title     string      `json:"title"`
	FeedURL   string      `json:"feed_url"`
	SiteURL   string      `json:"site_url"`
	JSONFeed  interface{} `json:"json_feed"`
}

type GetSubscriptionsReq struct {
	Since time.Time // since=2013-03-08T09:44:20.449047Z will get all subscriptions created after the iso 8601 timestamp.
	Mode  string    // the only mode available is extended. This includes more metadata for the feed.
}

type GetSubscriptionsResp struct {
	Subscriptions []*Subscription `json:"subscriptions"`
}
