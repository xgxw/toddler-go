package client

import (
	"github.com/xgxw/toddler-go/pb"
)

// IDemoClient DemoClient interface
type IDemoClient interface {
	IClient
	pb.DemoServiceClient
}

// DemoClient Demo rpc client
type DemoClient struct {
	IClient
	pb.DemoServiceClient
}

// NewDemoClient 创建一个连接 demo 的 gRPC 客户端。
func NewDemoClient(opts Options) (*DemoClient, error) {
	client, err := NewClient(opts)
	if err != nil {
		return &DemoClient{}, err
	}

	return &DemoClient{
		client,
		pb.NewDemoServiceClient(client.conn),
	}, nil
}
