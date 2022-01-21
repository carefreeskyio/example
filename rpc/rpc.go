package rpc

import (
	"carefreesky/global"
	"context"
	"github.com/carefreeskyio/example/proto"
	"github.com/carefreeskyio/rpcxclient"
	"strings"
	"sync"
	"time"
)


type Client struct {
	XClient *rpcxclient.Client
}

var (
	c    *Client
	once sync.Once
)

func NewClient() (*Client, error) {
	var (
		err        error
		rpcXClient *rpcxclient.Client
	)

	if c == nil {
		once.Do(func() {
			rpcXClient, err = rpcxclient.NewClient(initOptions())
			if err != nil {
				return
			}
			c = &Client{
				XClient: rpcXClient,
			}
		})
	}

	return c, err
}

// 获取初始化rpcXClient客户端属性，可根据实际需求修改
func initOptions() (options rpcxclient.Options) {
	options = rpcxclient.DefaultOptions
	options.BasePath = "/carefreesky"
	options.ServerName = "CarefreeSky"
	options.Addr = strings.Split(global.Config.Registry.Addr, " ")
	options.Group = global.Config.Registry.Group
	options.Timeout = time.Duration(global.Config.Rpc.WithTimeout) * time.Second

	return options
}

func (c *Client) Hello(ctx context.Context, request *proto.DemoHelloRequest, response *proto.DemoHelloResponse) (err error) {
	return c.XClient.Call(ctx, "Hello", request, response)
}


