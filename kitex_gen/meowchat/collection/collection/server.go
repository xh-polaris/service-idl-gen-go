// Code generated by Kitex v0.6.2. DO NOT EDIT.
package collection

import (
	server "github.com/cloudwego/kitex/server"
	collection "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/collection"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler collection.Collection, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
