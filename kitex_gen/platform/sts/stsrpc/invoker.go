// Code generated by Kitex v0.6.1. DO NOT EDIT.

package stsrpc

import (
	server "github.com/cloudwego/kitex/server"
	sts "github.com/xh-polaris/service-idl-gen-go/kitex_gen/platform/sts"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler sts.StsRpc, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
