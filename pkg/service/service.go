package service

import (
	"context"

	"github.com/refs/extension-bootstrap/pkg/proto"
)

// here we implement the compiled go-micro service

// Service implements the YellerService interface
type Service struct{}

// Yell implements the YellerService interface
func (s Service) Yell(c context.Context, req *proto.YellerRequest, res *proto.YellerResponse) error {
	// do business related stuff here
	return nil
}
