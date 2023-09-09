// Code generated by Kitex v0.7.1. DO NOT EDIT.

package sts

import (
	server "github.com/cloudwego/kitex/server"
	core_api "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/core_api"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler core_api.Sts, opts ...server.Option) server.Invoker {
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
