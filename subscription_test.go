package feedbin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Sub(t *testing.T) {
	as := assert.New(t)

	t.Run("", func(t *testing.T) {
		_, err := feedbinNoAuthIns.GetSubscriptions(ctx, &GetSubscriptionsReq{})
		as.NotNil(err)
		as.Equal("HTTP Basic: Access denied.", err.Error())
	})

	t.Run("", func(t *testing.T) {
		resp, err := feedbinIns.GetSubscriptions(ctx, &GetSubscriptionsReq{
			mode: "extended",
		})
		as.Nil(err)
		as.True(len(resp.Subscriptions) > 0)
	})
}
