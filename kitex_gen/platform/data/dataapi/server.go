// Code generated by Kitex v0.10.3. DO NOT EDIT.
package dataapi

import (
	server "github.com/cloudwego/kitex/server"
	data "github.com/xh-polaris/service-idl-gen-go/kitex_gen/platform/data"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler data.DataApi, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

func RegisterService(svr server.Server, handler data.DataApi, opts ...server.RegisterOption) error {
	return svr.RegisterService(serviceInfo(), handler, opts...)
}
