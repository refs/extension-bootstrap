package server

// this package hooks the service implementation (also known as a handler) as a go-micro service

import (
	"github.com/micro/go-micro"
	"github.com/owncloud/ocis-pkg/service/grpc"
	"github.com/refs/extension-bootstrap/pkg/proto"
	svc "github.com/refs/extension-bootstrap/pkg/service"
)

// Service undocummented
func Service() (micro.Service, error) {

	// create our grpc service
	service := grpc.NewService(
		grpc.Namespace("com.hello"),
		grpc.Name("yeller"),
		grpc.Address("localhost:10001"),
	)

	// declare our service handler
	handler := svc.NewService()

	// register it
	proto.RegisterYellerServiceHandler(service.Server(), handler)

	// initialize it
	service.Init()

	return service, nil
}
