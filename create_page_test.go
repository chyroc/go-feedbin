package feedbin

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

var (
	Username = os.Getenv("FEEDBIN_USERNAME")
	Password = os.Getenv("FEEDBIN_PASSWORD")
)

var (
	feedbinIns       = New(WithCredential(Username, Password))
	feedbinNoAuthIns = New(WithCredential("", ""))
	ctx              = context.Background()
)

func Test_Page(t *testing.T) {
	as := assert.New(t)

	t.Run("", func(t *testing.T) {
		_, err := feedbinNoAuthIns.CreatePage(ctx, &CreatePageReq{})
		as.NotNil(err)
		as.Equal("HTTP Basic: Access denied.", err.Error())
	})

	t.Run("", func(t *testing.T) {
		resp, err := feedbinIns.CreatePage(ctx, &CreatePageReq{})
		spew.Dump(resp, err)
		as.NotNil(err)
		as.Equal("Missing required key: url", err.Error())
	})

	t.Run("", func(t *testing.T) {
		resp, err := feedbinIns.CreatePage(ctx, &CreatePageReq{
			URL: "https://36kr.com/p/1412646752032130",
		})
		as.Nil(err)
		as.NotNil(resp)
		as.Contains(resp.Title, replaceSpace("这次，iPhone 终于要换 USB-C 接口了？"))
		as.Contains(resp.Summary, "在 USB Type-C 日渐普及的今天，iPhone 上的 Lightning 闪电接口是个让人没有爱只有恨的存在。")
		testContains(t, resp.Content, []string{
			"<strong>所以我们是否可以准备海淘欧版",
			"<strong>04 影响我买 iPhone 13 吗",
			"<li><p>电子设备与充电器分开销售",
			"https://img.36krcdn.com/20210924/v2_bea4f111a7a449a2a1633160a217e070_img_000",
		})
		as.Equal("https://extract.feedbin.com/parser/feedbin/3d89bf3a15bac8adce264906e5d95e278a18ea26?base64_url=aHR0cHM6Ly8zNmtyLmNvbS9wLzE0MTI2NDY3NTIwMzIxMzA=", resp.ExtractedContentURL)
	})

	t.Run("", func(t *testing.T) {
		resp, err := feedbinIns.CreatePage(ctx, &CreatePageReq{
			URL: "https://www.reddit.com/r/golang/comments/pvilb0/code_optimizations/",
		})
		as.Nil(err)
		as.NotNil(resp)
		as.Contains(resp.Title, "r/golang - Code optimizations")
		as.Contains(resp.Summary, "So I've recently stumbled upon the go build -gcflags=-m option and I find a lot of things end up on the heap that I sometimes prefer wouldn't")
		testContains(t, resp.Content, []string{
			"Thanks! And happy coding :)</p>",
			"I have noticed almost all the make commands end up on the heap for me",
		})
		as.Equal("https://extract.feedbin.com/parser/feedbin/b3a014bb9092a8d5e3ea8d59405dab7b2437b2ec?base64_url=aHR0cHM6Ly93d3cucmVkZGl0LmNvbS9yL2dvbGFuZy9jb21tZW50cy9wdmlsYjAvY29kZV9vcHRpbWl6YXRpb25zLw==", resp.ExtractedContentURL)
	})
}

func testContains(t *testing.T, text string, contains []string) {
	as := assert.New(t)
	for _, v := range contains {
		as.Contains(text, v)
	}
}

func replaceSpace(s string) string {
	return strings.ReplaceAll(s, " ", "\u00a0")
}
