package feedbin

import (
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func (r *Feedbin) ExtractingContentURL(ctx context.Context, request *ExtractingContentReq) string {
	signature := hmacSha1(r.password, request.URL)
	base64URL := base64.URLEncoding.EncodeToString([]byte(request.URL))
	return fmt.Sprintf("https://extract.feedbin.com/parser/%s/%s?base64_url=%s", url.PathEscape(r.username), signature, base64URL)
}

// ExtractingContent uses [Mercury Parser](https://github.com/postlight/mercury-parser) to attempt to extract the full content of a webpage.
//
// The extract service is only available to consumer apps that sync with feedbin.com. It is possible to get a cached article without an account, which is probably why you occasionally get a result.
func (r *Feedbin) ExtractingContent(ctx context.Context, request *ExtractingContentReq) (*ExtractingContentResp, error) {
	signature := hmacSha1(r.password, request.URL)
	base64URL := base64.URLEncoding.EncodeToString([]byte(request.URL))
	uri := fmt.Sprintf("https://extract.feedbin.com/parser/%s/%s?base64_url=%s", url.PathEscape(r.username), signature, base64URL)

	resp := new(ExtractingContentResp)
	if _, err := r.request(ctx, http.MethodGet, uri, nil, false, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type ExtractingContentReq struct {
	URL string `json:"url,omitempty"`
}

type ExtractingContentResp struct {
	Title         string      `json:"title"`
	Author        interface{} `json:"author"`
	DatePublished time.Time   `json:"date_published"`
	Dek           interface{} `json:"dek"`
	LeadImageURL  string      `json:"lead_image_url"`
	Content       string      `json:"content"`
	NextPageURL   interface{} `json:"next_page_url"`
	URL           string      `json:"url"`
	Domain        string      `json:"domain"`
	Excerpt       string      `json:"excerpt"`
	WordCount     int         `json:"word_count"`
	Direction     string      `json:"direction"`
	TotalPages    int         `json:"total_pages"`
	RenderedPages int         `json:"rendered_pages"`
}

func hmacSha1(secret, text string) string {
	key := []byte(secret)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(text))
	return fmt.Sprintf("%x", mac.Sum(nil))
}
