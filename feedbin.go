package feedbin

import (
	"net/http"
	"time"
)

type Feedbin struct {
	username   string
	password   string
	timeout    time.Duration
	httpClient *http.Client
}

type ClientOptionFunc func(*Feedbin)

func WithCredential(username, password string) ClientOptionFunc {
	return func(r *Feedbin) {
		r.username = username
		r.password = password
	}
}

func New(options ...ClientOptionFunc) *Feedbin {
	return newClient(options)
}

func newClient(options []ClientOptionFunc) *Feedbin {
	r := &Feedbin{
		timeout: time.Second * 3,
	}
	for _, v := range options {
		if v != nil {
			v(r)
		}
	}

	r.httpClient = &http.Client{
		Timeout: r.timeout,
	}

	return r
}
