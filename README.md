# go-feedbin

[![codecov](https://codecov.io/gh/chyroc/go-feedbin/branch/master/graph/badge.svg?token=Z73T6YFF80)](https://codecov.io/gh/chyroc/go-feedbin)
[![go report card](https://goreportcard.com/badge/github.com/chyroc/go-feedbin "go report card")](https://goreportcard.com/report/github.com/chyroc/go-feedbin)
[![test status](https://github.com/chyroc/go-feedbin/actions/workflows/test.yml/badge.svg)](https://github.com/chyroc/go-feedbin/actions)
[![Apache-2.0 license](https://img.shields.io/badge/License-Apache%202.0-brightgreen.svg)](https://opensource.org/licenses/Apache-2.0)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/chyroc/go-feedbin)
[![Go project version](https://badge.fury.io/go/github.com%2Fchyroc%2Fgo-feedbin.svg)](https://badge.fury.io/go/github.com%2Fchyroc%2Fgo-feedbin)

![](./header.jpg)

Feedbin API Documentation: https://github.com/feedbin/feedbin-api.

## Install

```shell
go get github.com/chyroc/go-feedbin
```

## Usage

## Create Page

```go
package main

import (
	"context"
	"fmt"

	"github.com/chyroc/go-feedbin"
)

func main() {
	url := ""
	cli := feedbin.New(feedbin.WithCredential("username", "password"))

	resp, err := cli.CreatePage(context.Background(), &feedbin.CreatePageReq{
		URL: url,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("title", resp.Title)
	fmt.Println("content", resp.Content)
}
```

## Get Subscriptions

```go
package main

import (
	"context"
	"fmt"

	"github.com/chyroc/go-feedbin"
)

func main() {
	cli := feedbin.New(feedbin.WithCredential("username", "password"))

	resp, err := cli.GetSubscriptions(context.Background(), &feedbin.GetSubscriptionsReq{})
	if err != nil {
		panic(err)
	}
	fmt.Println("subscriptions length:", len(resp.Subscriptions))
	for _, v := range resp.Subscriptions {
		fmt.Println(v.ID, v.Title, v.FeedURL)
	}
}
```

## Extracting Content

```go
package main

import (
	"context"
	"fmt"

	"github.com/chyroc/go-feedbin"
)

func main() {
	url := ""
	cli := feedbin.New()

	resp, err := cli.ExtractingContent(context.Background(), &feedbin.ExtractingContentReq{
		URL: url,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("title", resp.Title)
	fmt.Println("content", resp.Content)
}
```