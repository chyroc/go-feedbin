package feedbin_test

import (
	"context"
	"fmt"

	"github.com/chyroc/go-feedbin"
)

func ExampleFeedbin_CreatePage() {
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

func ExampleFeedbin_GetSubscriptions() {
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

func ExampleFeedbin_ExtractingContent() {
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
