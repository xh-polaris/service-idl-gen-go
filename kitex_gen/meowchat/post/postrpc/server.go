// Code generated by Kitex v0.6.1. DO NOT EDIT.
package postrpc

import (
	server "github.com/cloudwego/kitex/server"
	post "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/post"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler post.PostRpc, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
