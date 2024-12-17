// Code generated by Kitex v0.12.0. DO NOT EDIT.
package data

import (
	server "github.com/cloudwego/kitex/server"
	core_api "github.com/xh-polaris/service-idl-gen-go/kitex_gen/platform/core_api"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler core_api.Data, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

func RegisterService(svr server.Server, handler core_api.Data, opts ...server.RegisterOption) error {
	return svr.RegisterService(serviceInfo(), handler, opts...)
}
