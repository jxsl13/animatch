package anidb

import (
	"sync"
	"time"

	"github.com/jxsl13/animatch/common"
	"go.felesatra.moe/anidb"
)

const (
	// These values are registered with the AniDB
	ClientName    = "animatch"
	ClientVersion = 100
)

var (
	onceClient = sync.Once{}
	client     = (*Client)(nil)
)

type Client = anidb.Client

func NewClient() *Client {
	onceClient.Do(func() {
		c := &Client{
			Name:    ClientName,
			Version: ClientVersion,
			Limiter: common.NewRateLimiter(time.Second, 1),
		}
		client = c
	})
	return client
}
