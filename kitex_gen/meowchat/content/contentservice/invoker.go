// Code generated by Kitex v0.8.0. DO NOT EDIT.

package contentservice

import (
	server "github.com/cloudwego/kitex/server"
	content "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/content"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler content.ContentService, opts ...server.Option) server.Invoker {
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
