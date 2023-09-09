// Code generated by Kitex v0.7.1. DO NOT EDIT.
package contentservice

import (
	server "github.com/cloudwego/kitex/server"
	content "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/content"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler content.ContentService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
