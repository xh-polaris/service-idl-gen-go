// Code generated by Kitex v0.6.2. DO NOT EDIT.

package momentrpc

import (
	server "github.com/cloudwego/kitex/server"
	moment "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/moment"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler moment.MomentRpc, opts ...server.Option) server.Invoker {
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
