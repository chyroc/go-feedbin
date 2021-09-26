package feedbin

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func Test_ExtracingContent(t *testing.T) {
	as := assert.New(t)
	uri := "https://www.economist.com/the-americas/2021/09/25/a-monk-in-14th-century-italy-wrote-about-the-americas"
	req := &ExtractingContentReq{URL: uri}

	t.Run("", func(t *testing.T) {
		as.Regexp(`https://extract.feedbin.com/parser/.*?/a4ce4c9d571df63c40e2a79495011211433c8ee9\?base64_url=aHR0cHM6Ly93d3cuZWNvbm9taXN0LmNvbS90aGUtYW1lcmljYXMvMjAyMS8wOS8yNS9hLW1vbmstaW4tMTR0aC1jZW50dXJ5LWl0YWx5LXdyb3RlLWFib3V0LXRoZS1hbWVyaWNhcw==`,
			feedbinIns.ExtractingContentURL(ctx, req))
	})

	t.Run("", func(t *testing.T) {
		resp, err := feedbinIns.ExtractingContent(ctx, req)
		as.Nil(err)
		as.NotNil(resp)
		spew.Dump(resp)
		as.Equal("A monk in 14th-century Italy wrote about the Americas", resp.Title)
		as.Equal("https://www.economist.com/img/b/1280/720/90/sites/default/files/images/2021/09/articles/main/20210925_amp502.jpg", resp.LeadImageURL)
		as.Equal("That was long before Christopher Columbus set sail | The Americas", resp.Excerpt)
		as.Contains(resp.Content, "In 1960 the remains of Norse buildings were found on Newfoundland.")
	})
}
