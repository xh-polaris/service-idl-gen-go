// Code generated by Kitex v0.7.1. DO NOT EDIT.
package systemrpc

import (
	server "github.com/cloudwego/kitex/server"
	system "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/system"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler system.SystemRpc, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
