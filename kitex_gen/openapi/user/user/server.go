// Code generated by Kitex v0.11.3. DO NOT EDIT.
package user

import (
	server "github.com/cloudwego/kitex/server"
	user "github.com/xh-polaris/service-idl-gen-go/kitex_gen/openapi/user"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler user.User, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

func RegisterService(svr server.Server, handler user.User, opts ...server.RegisterOption) error {
	return svr.RegisterService(serviceInfo(), handler, opts...)
}